package tasks

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/tech-thinker/stikky/config"
	"github.com/tech-thinker/stikky/utils"
)

type Task interface {
	Base64Encode(ctx context.Context, plain string) (string, error)
	Base64Decode(ctx context.Context, encodededText string) (string, error)
	Encrypt(ctx context.Context, plainText string) (string, error)
	Decrypt(ctx context.Context, cipherText string) (string, error)
	UUIDGenerate(ctx context.Context) (string, error)
}

type task struct {
	cfg config.AppConfig
}

func (*task) Base64Encode(ctx context.Context, plain string) (string, error) {
	return utils.Base64Encode(plain)
}

func (*task) Base64Decode(ctx context.Context, encodededText string) (string, error) {
	return utils.Base64Decode(encodededText)
}

func (t *task) Encrypt(ctx context.Context, plainText string) (string, error) {
	fmt.Println("Public Key: ", t.cfg.GetPublicKey())
	pubKey, err := utils.LoadPublicKeyFromString(t.cfg.GetPublicKey())
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return utils.EncryptWithPublicKey(plainText, pubKey), nil
}

func (t *task) Decrypt(ctx context.Context, cipherText string) (string, error) {
	fmt.Println("Private Key: ", t.cfg.GetPrivateKey())
	pkey, err := utils.LoadPrivateKeyFromString(t.cfg.GetPrivateKey())
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return utils.DecryptWithPrivateKey(cipherText, pkey), nil
}

func (*task) UUIDGenerate(ctx context.Context) (string, error) {
	return uuid.New().String(), nil
}

func NewTask(cfg config.AppConfig) Task {
	return &task{
		cfg: cfg,
	}
}
