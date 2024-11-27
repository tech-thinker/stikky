package utils

import "encoding/base64"

func Base64Encode(plain string) (string, error) {
	return base64.StdEncoding.EncodeToString([]byte(plain)), nil
}

func Base64Decode(encodededText string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(encodededText)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}
