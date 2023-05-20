package utils

import (
	"mime/multipart"
)

type Files struct {
	Name string             `json:"name"`
	Type string             `json:"type"`
	Path string             `json:"path"`
	Description string      `json:"description"`
}

type FileWithDescription struct {
	File        *multipart.FileHeader
	Description string
}
