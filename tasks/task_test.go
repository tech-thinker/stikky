package tasks

import (
	"context"
	"testing"

	"github.com/tech-thinker/stikky/config"
	"github.com/tech-thinker/stikky/utils"
)

func Test_task_Base64Encode(t *testing.T) {
	type args struct {
		ctx   context.Context
		plain string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test_base64_encode",
			args: args{
				ctx:   context.Background(),
				plain: "hello",
			},
			want:    "aGVsbG8=",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := NewTask(config.NewAppConfig())
			got, err := tr.Base64Encode(tt.args.ctx, tt.args.plain)
			if (err != nil) != tt.wantErr {
				t.Errorf("task.Base64Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("task.Base64Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_task_Base64Decode(t *testing.T) {
	type args struct {
		ctx           context.Context
		encodededText string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test_base64_decode",
			args: args{
				ctx:           context.Background(),
				encodededText: "aGVsbG8=",
			},
			want:    "hello",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := NewTask(config.NewAppConfig())
			got, err := tr.Base64Decode(tt.args.ctx, tt.args.encodededText)
			if (err != nil) != tt.wantErr {
				t.Errorf("task.Base64Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("task.Base64Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_task_EncryptDecrypt(t *testing.T) {
	privateKey, publicKey, _ := utils.GenerateKeyPair(4096)
	type args struct {
		ctx        context.Context
		plainText  string
		privateKey string
		publicKey  string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "encrypt and decrypt should success if valid key is provided",
			args: args{
				ctx:        context.Background(),
				plainText:  "hello",
				privateKey: privateKey,
				publicKey:  publicKey,
			},
			want:    "hello",
			wantErr: false,
		},
		{
			name: "encrypt and decrypt should fail if invalid key is provided",
			args: args{
				ctx:        context.Background(),
				plainText:  "hello",
				privateKey: "invalid_key",
				publicKey:  "invalid_key",
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.NewAppConfig()
			cfg.SetPrivateKey(tt.args.privateKey)
			cfg.SetPublicKey(tt.args.publicKey)

			tr := NewTask(cfg)

			got, err := tr.Encrypt(tt.args.ctx, tt.args.plainText)

			if (err != nil) != tt.wantErr {
				t.Errorf("task.Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, err = tr.Decrypt(tt.args.ctx, got)

			if got != tt.want {
				t.Errorf("task.Encrypt() = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_task_UUIDGenerate(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "test_uuid_generate",
			args: args{
				ctx: context.Background(),
			},
			want:    36, // length of string
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := NewTask(config.NewAppConfig())
			got, err := tr.UUIDGenerate(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("task.UUIDGenerate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("task.UUIDGenerate() = %v, want %v", got, tt.want)
			}
		})
	}
}
