package config

import (
	"github.com/tech-thinker/stikky/models"
	"github.com/tech-thinker/stikky/utils"
)

type AppConfig interface {
	GetPrivateKey() string
    SetPrivateKey(string)
	GetPublicKey() string
    SetPublicKey(string)
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

func (c *appConfig) SetPrivateKey(pk string) {
    c.PrivateKey = pk
}

func (c *appConfig) GetPublicKey() string {
	return c.PublicKey
}

func (c *appConfig) SetPublicKey(puk string) {
    c.PublicKey = puk
}

func (c *appConfig) GetPublicKeys() []models.PublicKey {
	return c.PublicKeys
}

func NewAppConfig() AppConfig {
	pk, puk, _ := utils.GenerateKeyPair(4096)
	cfg := appConfig{}
	cfg.PrivateKey = pk
	cfg.PublicKey = puk
	return &cfg
}
