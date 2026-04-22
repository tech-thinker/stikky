package storage

import (
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
)

func GetDatabasePath(a fyne.App) string {
	p := a.Storage().RootURI().Path()
	os.MkdirAll(p, 0755)
	return filepath.Join(p, "stikky.db")
}
