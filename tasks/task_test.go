package tasks

import (
	"context"
	"testing"
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
			tr := &task{}
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
			tr := &task{}
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
