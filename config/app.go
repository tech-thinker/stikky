package config

import (
	"github.com/tech-thinker/stikky/models"
)

type AppConfig interface {
	GetPrivateKey() string
	GetPublicKey() string
	GetPublicKeys() []models.PublicKey
}

type appConfig struct {
	PrivateKey string
	PublicKey  string
	PublicKeys []models.PublicKey
}

func (c *appConfig) GetPrivateKey() string {
	return c.PrivateKey
}

func (c *appConfig) GetPublicKey() string {
	return c.PublicKey
}

func (c *appConfig) GetPublicKeys() []models.PublicKey {
	return c.PublicKeys
}

func NewAppConfig() AppConfig {
	cfg := appConfig{}

	return &cfg
}
