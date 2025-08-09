package conf

import (
	"encoding/json"
	"os"
)

type Config struct {
	Mode      string     `json:"mode"`
	LocalPort int        `json:"localport"`
	Policy    string     `json:"policy"`
	Upstream  []Upstream `json:"upstream"`
	Workers   string     `json:"workers"`
	LogLevel  string     `json:"log_level"`
	LogFile   string     `json:"log_file"`
}

type Upstream struct {
	UpstreamIP   string `json:"upstream_ip"`
	UpstreamPort int    `json:"upstream_port"`
	Weight       int    `json:"weight"`
}

func ParseConfFile(filePath string) (Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
