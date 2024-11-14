package config

import (
	"os"
)

type FontPair struct {
	Import   string `json:"import"`
	Headline string `json:"headline"`
	Body     string `json:"body"`
}

type ElementCategory struct {
	Name               string   `json:"name"`
	Elements           []string `json:"elements"`
	InsertableElements []string `json:"insertable-elements"`
	CSSProperties      []string `json:"css-properties"`
}

type Config struct {
	// BaseUrl                                  string
}

func IsDevelopment() bool {
	return os.Getenv("GOWIRE_ENV") == "development"
}
