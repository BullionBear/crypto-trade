package config

import (
	"os"

	"github.com/yosuke-furukawa/json5/encoding/json5"
)

// Config represents the application configuration.
// Config struct matches the structure of the JSON configuration.
type AlexConfig struct {
	Symbol  string `json:"symbol"`
	Account struct {
		ApiKey    string `json:"api_key"`
		ApiSecret string `json:"api_secret"`
	} `json:"account"`
}

// LoadConfig loads configuration from the specified file path.
func LoadAlexConfig(path string) (*AlexConfig, error) {
	var config AlexConfig
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err := json5.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

type NikoConfig struct {
	Symbol     string `json:"symbol"`
	GrpcClient struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"grpc_client"`
	MongoUri string `json:"mongo_uri"`
	Balance  []struct {
		Coin   string  `json:"coin"`
		Amount float64 `json:"amount"`
	} `json:"balance"`
}

func LoadNikoConfig(path string) (*NikoConfig, error) {
	var config NikoConfig
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err := json5.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
