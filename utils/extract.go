package utils

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ExtractArchive(destPath string, archivePath string) error {
	if strings.Contains(archivePath, ".zip") || strings.Contains(archivePath, ".jar") {
		return extractZIP(destPath, archivePath)
	} else {
		return extractTarGz(destPath, archivePath)
	}
}

func extractZIP(destPath string, zipPath string) error {
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

func extractTarGz(destPath string, archivePath string) error {
	file, err := os.Open(archivePath)
	if err != nil {
		return err
	}
	defer file.Close()

	gzr, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	for {
		header, _ := tr.Next()
		if err == io.EOF {
			fmt.Printf("    Extracted %s to %s\n", archivePath, destPath)

			break
		}

		targetPath := filepath.Join(destPath, header.Name)


		switch header.Typeflag {
		case tar.TypeDir:
			os.MkdirAll(targetPath, os.FileMode(header.Mode))
		case tar.TypeReg:
			destFile, _ := os.Create(targetPath)

			io.Copy(destFile, tr)
			destFile.Close()
		}
	}

	return nil
}