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

	file_i, err := os.Open(absPath)

	img, err := vips.NewImageFromReader(file_i)
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	// Create temp file
	filePath := fmt.Sprintf("%d.%s", time.Now().Unix(), "webp")
	filePath = filepath.Join(os.TempDir(), filePath)

	// 0600: file can be RW by owner
	file_o, err := os.Create(filePath)
	if err != nil {
		return "", err
	}

	err = img.SaveToWriterWebp(file_o, nil)
	if err != nil {
		return "", err
	}

	file_i.Close()
	file_o.Close()

	return filePath, nil
}
