package utils

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"fmt"
)

func  SaveFiles(protocolNumber string, files []*multipart.FileHeader) ([]Files, error) {
	savedFiles := []Files{}

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return nil, fmt.Errorf("error opening file: %v", err)
		}
		defer src.Close()

		fileName := protocolNumber + "_" + file.Filename
		filePath := filepath.Join("uploads", protocolNumber, fileName)

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return nil, fmt.Errorf("error creating directories: %v", err)
		}

		dst, err := os.Create(filePath)
		if err != nil {
			return nil, fmt.Errorf("error creating file: %v", err)
		}
		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			return nil, fmt.Errorf("error copying file: %v", err)
		}

		newFile := Files{
			Name:     file.Filename,
			Type:     file.Header.Get("Content-Type"),
			Path: filePath,
		}
		savedFiles = append(savedFiles, newFile)
	}

	return savedFiles, nil
}
