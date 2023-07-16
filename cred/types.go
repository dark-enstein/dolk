package cred

import (
	"encoding/json"
	"log"
)

var (
	ErrFailedToMarshalJson = "failed to marshal json"
)

// Config is the object containing credentials information as pertaining the
// request. This would make up the response json. It ha
type Config struct {
	// Config contains the details of the credential information to be
	//returned to the client
	Config []byte
}

// convert login config data into json that will be sent to the client
func (cfg *Config) String() string {
	jsonResp, err := json.Marshal(cfg.Config)

	if err != nil {
		log.Println(ErrFailedToMarshalJson)
		return ""
	}

	return string(jsonResp)
}
