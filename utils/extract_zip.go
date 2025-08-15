package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ExtractZIP(destPath string, zipPath string) error {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		filePath := filepath.Join(destPath, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		destFile, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, rc)		
		if err != nil {
			return err
		}
	}
	fmt.Printf("    Extracted %s to %s\n", zipPath, destPath)
	return nil
}