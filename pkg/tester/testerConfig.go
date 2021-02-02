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
	byteArrayBody []byte

}

func CreateTesterConfigFromPath(path string) (*TesterConfig, error) {
	config := &TesterConfig{}
	rawConfig, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(rawConfig, config)
	if err != nil {
		return nil, err
	}
	config.byteArrayBody, err = ioutil.ReadFile(config.JsonBodyPath)
	if err != nil {
		panic("Error al cargar el json, verifique que JsonBodyPath sea una ubicación existente")
	}

	return config, nil
}

func (config *TesterConfig) getBodyAsByteArray() []byte {
	return config.byteArrayBody
}
