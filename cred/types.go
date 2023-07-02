package cred

import (
	"encoding/json"
	"fmt"
	"log"
)

var (
	ErrFailedToUnmarshalJson = "failed to unmarshal json"
)

type Config struct {
	Config   []byte
	Metadata string
}

func (cfg *Config) String() string {
	var config string
	err := json.Unmarshal(cfg.Config, config)

	if err != nil {
		log.Println(ErrFailedToUnmarshalJson)
		return ErrFailedToUnmarshalJson
	}

	return fmt.Sprintf("config: %v\nmetadata: %v", config,
		cfg.Metadata)
}
