package sonyflakex

import (
	"fmt"
	"github.com/0xjasoncao/gin-scaffold/configs/config"
	"github.com/sony/sonyflake/v2"
)

var sf *sonyflake.Sonyflake

func InitSonyFlake(config *config.Config) {
	// In Settings, the default MachineID uses the lower 16 bits of the IPv4 address.
	// You should configure it according to your actual environment to ensure uniqueness
	// (e.g., when running in Docker, avoid using IP if it may be duplicated).
	settings := sonyflake.Settings{}

	var err error
	// Create a new SonyFlake instance with the specified settings
	sf, err = sonyflake.New(settings)

	if err != nil {
		// Panic if Sonyflake initialization fails
		panic(fmt.Sprintf("Failed to initialize Sonyflake: %v \n", err))
	}

	// Print SonyFlake instance details for debugging
	fmt.Printf("Sonyflake initialized successfully\n%#v \n", sf)
}

// NewSonyFlakeId generates a new unique SonyFlake ID.
func NewSonyFlakeId() int64 {
	id, _ := sf.NextID()
	return id
}
