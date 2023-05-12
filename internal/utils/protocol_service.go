package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

type ProtocolService struct {
	ProtocolRepo *ProtocolRepository
	Logger       *logrus.Logger
}

func NewProtocolService(protocolRepo *ProtocolRepository, logger *logrus.Logger) *ProtocolService {
	return &ProtocolService{
		ProtocolRepo: protocolRepo,
		Logger:       logger,
	}
}

func (u *ProtocolService) generateProtocolNumber(letter string) string {
	// generate random 9-digit string
	rand.Seed(time.Now().UnixNano())
	protocolNumber := fmt.Sprintf("%s%09d", letter, rand.Intn(1000000000))
	return protocolNumber
}

func (u *ProtocolService) GenerateUniqueProtocolNumber(letter string) (string, error) {
	u.Logger.Infof("Service GenerateUniqueProtocolNumber")
	for i := 0; i < 100; i++ {
		// Generate a random protocol number
		protocolNumber := u.generateProtocolNumber(letter)

		// Check if the protocol number already exists in the database
		exists, err := u.ProtocolRepo.ProtocolNumberExists(protocolNumber)
		if err != nil {
			return "", err
		}

		// If the protocol number does not exist, save it to the database and return
		if !exists {
			err := u.ProtocolRepo.SaveProtocolNumber(protocolNumber)
			if err != nil {
				return "", err
			}
			return protocolNumber, nil
		}
	}
	u.Logger.Infof("Service GenerateUniqueProtocolNumber OK")
	return "", fmt.Errorf("failed to generate a unique protocol number after multiple attempts")
}
