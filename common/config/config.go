package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/common/database"
)

//ERC20
// 0x83F3c5020Ef0f44C8Ef4993124740D3fe8D1470C

// SBT
// 0x9b11f74888dF35573B934567088578eEe485B663

// topic
// 0x9Af4f4b7C831b0c79574CCDE7C04e33F99BF6438

// ListenerConfig is the configuration of Listener.
type ListenerConfig struct {
	Confirmation      uint64 `json:"confirmation"`
	Endpoint          string `json:"endpoint"`
	StartHeight       uint64 `json:"startHeight"` // Can only be configured to contract deployment height
	BlockTime         int64  `json:"blockTime"`
	FetchLimit        uint64 `json:"fetchLimit"`
	TopicContractAddr string `json:"TopicContractAddr"`
	SBTContractAddr   string `json:"SBTContractAddr"`
}

// RpcConfig is the configuration of rpc.
type RpcConfig struct {
	Port string `json:"port"`
}

type Config struct {
	Rpc      *RpcConfig      `json:"rpc"`
	Listener *ListenerConfig `json:"listener"`
	DB       database.Config `json:"db"`
}

func NewConfig(file string) (*Config, error) {
	buf, err := os.ReadFile(filepath.Clean(file))
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = json.Unmarshal(buf, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
