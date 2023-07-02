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
	config   []byte
	metadata string
}

func (cfg *Config) String() string {
	var config string
	err := json.Unmarshal(cfg.config, config)

	if err != nil {
		log.Println(ErrFailedToUnmarshalJson)
		return ErrFailedToUnmarshalJson
	}

	return fmt.Sprintf("config: %v\nmetadata: %v", config,
		cfg.metadata)
}
