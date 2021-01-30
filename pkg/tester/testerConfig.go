package tester

import (
	"encoding/json"
	"io/ioutil"
)

type TesterConfig struct {

	Method string `json:"method"`
	Url string `json:"url"`
	JwtHeader string `json:"jwtHeader"`
	JwtHeaderValue string `json:"jwtHeaderValue"`
	JsonBodyPath string `json:"jsonBodyPath"`

}

func LoadConfig(path string) (*TesterConfig, error) {
	config := &TesterConfig{}
	rawConfig, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(rawConfig, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
