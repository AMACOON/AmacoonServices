package utils

import (
	"bytes"
	"io"
	"mime/multipart"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/sirupsen/logrus"
)

type FilesService struct {
	S3Client *s3.S3
	Logger   *logrus.Logger
}

func NewFilesService(s3Client *s3.S3, logger *logrus.Logger) *FilesService {
	return &FilesService{
		S3Client: s3Client,
		Logger:   logger,
	}
}

func (s *FilesService) SaveFiles(protocolNumber string, files []*multipart.FileHeader) ([]Files, error) {
	s.Logger.Infof("Saving %d files for protocol %s", len(files), protocolNumber)
	savedFiles := []Files{}

	bucket := "amacoondocs" // Atualize com o nome do seu bucket

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			s.Logger.Errorf("error opening file: %v", err)
			return nil, err
		}
		defer src.Close()

		fileName := file.Filename
		filePath := filepath.Join("services", protocolNumber, fileName)

		buf := new(bytes.Buffer)
		if _, err := io.Copy(buf, src); err != nil {
			s.Logger.Errorf("error copying file: %v", err)
			return nil, err
		}

		_, err = s.S3Client.PutObject(&s3.PutObjectInput{
			Bucket:      aws.String(bucket),
			Key:         aws.String(filePath),
			Body:        bytes.NewReader(buf.Bytes()),
			ContentType: aws.String(file.Header.Get("Content-Type")),
		})
		if err != nil {
			s.Logger.Errorf("error uploading file to S3: %v", err)
			return nil, err
		}

		newFile := Files{
			Name: file.Filename,
			Type: file.Header.Get("Content-Type"),
			Path: filePath,
		}
		savedFiles = append(savedFiles, newFile)
	}
	s.Logger.Infof("Saved %d files for protocol %s", len(files), protocolNumber)
	return savedFiles, nil
}
