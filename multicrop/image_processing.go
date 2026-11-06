package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/davidbyttow/govips/v2/vips"
)

/*
Converts input image to a webp proxy for displaying in
the frontend.
*/
func GenerateWebpProxy(absPath string) (string, error) {

	img, err := vips.NewImageFromFile(absPath)
	if err != nil {
		return "", err
	}

	buf, _, err := img.ExportWebp(nil)

	if err != nil {
		return "", err
	}

	// Create temp file
	filePath := fmt.Sprintf("%d.%s", time.Now().Unix(), "webp")
	filePath = filepath.Join(os.TempDir(), filePath)

	// 0600: file can be RW by owner
	err = os.WriteFile(filePath, buf, 0600)
	if err != nil {
		return "", err
	}
	return filePath, nil
}
