package utils

import (
    "fmt"
    "math/rand"
    "time"
    "github.com/sirupsen/logrus"
)


type ProtocolService struct {
    ProtocolRepo      *ProtocolRepository
    Logger *logrus.Logger
}

func NewProtocolService(protocolRepo *ProtocolRepository, logger *logrus.Logger) *ProtocolService {
	return &ProtocolService{
        ProtocolRepo:      protocolRepo,
        Logger: logger,
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
        // Gere um número de protocolo aleatório
        protocolNumber := u.generateProtocolNumber(letter)

        // Verifique se o número do protocolo já existe no banco de dados
        exists, err := u.ProtocolRepo.ProtocolNumberExists(protocolNumber)
        if err != nil {
            return "", err
        }

        // Se o número do protocolo não existir, salve-o no banco de dados e retorne
        if !exists {
            err := u.ProtocolRepo.SaveProtocolNumber(protocolNumber)
            if err != nil {
                return "", err
            }
            return protocolNumber, nil
        }
    }
    u.Logger.Infof("Service GenerateUniqueProtocolNumber OK")
    return "", fmt.Errorf("não foi possível gerar um número de protocolo único após várias tentativas")
}
