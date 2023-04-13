package utils

import (
    "fmt"
    "math/rand"
    "time"
)


type ProtocolService struct {
}

func NewProtocolService() *ProtocolService {
	return &ProtocolService{}
}

func (u *ProtocolService) GenerateProtocolNumber(letter string) string {
    // generate random 9-digit string
    rand.Seed(time.Now().UnixNano())
    protocolNumber := fmt.Sprintf("%s%09d", letter, rand.Intn(1000000000))
    return protocolNumber
}

