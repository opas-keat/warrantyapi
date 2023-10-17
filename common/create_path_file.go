package common

import (
	"os"
)

func CreatePathFileForUpload(nowDate string) string {
	path := "./uploads/" + nowDate + "/"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}
	return path
}
