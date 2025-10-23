package initialize

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gin-scaffold/internal/config"

	"gin-scaffold/pkg/logging"
	"gin-scaffold/pkg/sonyflakex"
)

type Options struct {
	ConfigFileDir string
	Version       string
}

func RunServer(ctx context.Context, options *Options) error {
	state := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	cleanFunc, err := runServer(ctx, options)
	if err != nil {
		return err
	}

EXIT:
	for {
		sig := <-sc
		logging.WithContext(ctx).Sugar().Infof("Catched signal[%s].", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			state = 0
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	cleanFunc()
	logging.WithContext(ctx).Sugar().Infof("Graceful shutdown completed successfully. Service is now stopped.")
	time.Sleep(time.Second)
	os.Exit(state)
	return nil
}

func runServer(ctx context.Context, options *Options) (func(), error) {
	// Load Config
	config.MustLoad(options.ConfigFileDir)
	//Init Logger
	err := logging.InitLogger(ctx, config.C.Logger.Outputs...)
	if err != nil {
		return nil, err
	}
	config.C.PrintWithJSON()
	//Init SonyFlake
	sonyflakex.InitSonyFlake()

	// Build Injector
	injector, injectorCleanFunc, err := BuildInjector(ctx, config.C)
	if err != nil {
		return nil, err
	}

	//Init HTTP Server
	engine := injector.Engine

	httpCleanFunc, err := initHttpServer(ctx, engine)
	if err != nil {
		return nil, err
	}

	return func() {
		httpCleanFunc()
		injectorCleanFunc()
	}, nil
}

func initHttpServer(ctx context.Context, handler http.Handler) (func(), error) {
	serverConf := config.C.Http
	addr := fmt.Sprintf("%s:%d", serverConf.Host, serverConf.Port)
	srv := &http.Server{
		Addr:           addr,
		Handler:        handler,
		ReadTimeout:    time.Duration(serverConf.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(serverConf.WriteTimeout) * time.Second,
		IdleTimeout:    time.Duration(serverConf.IdleTimeout) * time.Second,
		MaxHeaderBytes: serverConf.MaxHeaderBytes << 20,
	}

	go func() {
		logging.WithContext(ctx).Sugar().Infof("HTTP server is running at %s", addr)
		err := srv.ListenAndServe()
		if err != nil && !errors.As(http.ErrServerClosed, &err) {
			panic(err)
		}
	}()

	return func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(serverConf.ShutdownTimeout))
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			logging.WithContext(ctx).Error(err.Error())
		}

	}, nil
}
