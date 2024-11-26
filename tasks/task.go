package tasks

import (
	"context"
	"encoding/base64"

	"github.com/google/uuid"
)

type Task interface {
	Base64Encode(ctx context.Context, plain string) (string, error)
	Base64Decode(ctx context.Context, encodededText string) (string, error)
	UUIDGenerate(ctx context.Context) (string, error)
}

type task struct {
}

func (*task) Base64Encode(ctx context.Context, plain string) (string, error) {
	return base64.StdEncoding.EncodeToString([]byte(plain)), nil
}

func (*task) Base64Decode(ctx context.Context, encodededText string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(encodededText)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

func (*task) UUIDGenerate(ctx context.Context) (string, error) {

	return uuid.New().String(), nil
}

func NewTask() Task {
	return &task{}
}
