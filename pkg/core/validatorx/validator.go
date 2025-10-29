package validatorx

import (
	"errors"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type ValidationErrors validator.ValidationErrors

var zhTrans ut.Translator
var enTrans ut.Translator

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		uni := ut.New(zh.New(), en.New())
		zhTrans, _ = uni.GetTranslator("zh")
		enTrans, _ = uni.GetTranslator("en")

		// 获取struct tag里自定义的json或form作为字段名
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("json")
			if name != "" {
				return name
			}
			name = fld.Tag.Get("form")
			if name != "" {
				return name
			}
			return fld.Name
		})

		_ = zh_translations.RegisterDefaultTranslations(v, zhTrans)
		_ = en_translations.RegisterDefaultTranslations(v, enTrans)

	}
}

func ZhTranslate(err error) string {
	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		var errMessages []string
		for _, err := range errs {
			errMessages = append(errMessages, err.Translate(zhTrans))

		}
		return strings.Join(errMessages, ";")
	}

	return err.Error()
}

func EnTranslate(err error) string {
	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		var errMessages []string
		for _, err := range errs {
			errMessages = append(errMessages, err.Translate(enTrans))

		}
		return strings.Join(errMessages, ";")
	}

	return err.Error()
}
