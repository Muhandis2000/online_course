package utils

import (
	"log"
	"os"
	"path/filepath"
)

type Logger struct {
	Info  *log.Logger
	Error *log.Logger
}

func NewLogger(directory, filename string) (*Logger, error) {
	if err := os.MkdirAll(directory, 0755); err != nil {
		return nil, err
	}
	file, err := os.OpenFile(filepath.Join(directory, filename), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &Logger{
		Info:  log.New(file, "INFO: ", log.Ldate|log.Ltime),
		Error: log.New(file, "ERROR: ", log.Ldate|log.Ltime),
	}, nil
}
