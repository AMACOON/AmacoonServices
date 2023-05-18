package utils

import (
	"github.com/google/uuid"
)




func GenerateRandomString() string {
	uuidWithHyphen := uuid.New().String()
	uuidWithoutHyphen := uuidWithHyphen[0:8] + uuidWithHyphen[9:13] + uuidWithHyphen[14:18] + uuidWithHyphen[19:23] + uuidWithHyphen[24:36]
	randomString := uuidWithoutHyphen[0:10]
	return randomString
}