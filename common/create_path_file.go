package common

import (
	"os"
)

func CreatePathFileForUpload(warrantyNo string) string {
	path := "./uploads/" + warrantyNo + "/"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}
	return path
}
