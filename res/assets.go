package res

import (
	"embed"
)

//go:embed assets/*
var assets embed.FS

// Public function to access the icon data
func GetIcon() ([]byte, error) {
	return assets.ReadFile("assets/icon.png")
}

// Public function to access the icon data
func GetIconIco() ([]byte, error) {
	return assets.ReadFile("assets/icon.ico")
}
