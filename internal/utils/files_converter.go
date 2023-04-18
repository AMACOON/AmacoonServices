package utils

import(
"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertFilesReqToFiles(filesReqList []FilesReq) []Files {
	filesList := make([]Files, len(filesReqList))

	for i, filesReq := range filesReqList {
		filesList[i] = Files{
			ID:       primitive.NewObjectID(), // Gere um novo ObjectID
			Name:     filesReq.Name,
			Type:     filesReq.Type,
			Base64:   filesReq.Base64,
		}
	}

	return filesList
}