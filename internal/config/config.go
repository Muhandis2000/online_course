package config

import (
	"encoding/json"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		Port string `json:"port"`
	} `json:"server"`
	Database struct {
		Host   string `json:"host"`
		Port   int    `json:"port"`
		User   string `json:"user"`
		DBName string `json:"dbname"`
	} `json:"database"`
	Log struct {
		Directory string `json:"directory"`
		Filename  string `json:"filename"`
	} `json:"log"`
}

func LoadConfig() (Config, error) {
	var cfg Config
	if err := godotenv.Load(); err != nil {
		return cfg, err
	}
	file, err := os.Open("config/config.json")
	if err != nil {
		return cfg, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	return cfg, decoder.Decode(&cfg)
}
