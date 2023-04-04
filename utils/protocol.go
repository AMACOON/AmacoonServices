package util

import (
	
	"amacoonservices/models/utils"
	"fmt"

	"gorm.io/gorm"
)

type ProtocolService struct {
	DB *gorm.DB
}

func NewProtocolService(db *gorm.DB) *ProtocolService {
	return &ProtocolService{DB: db}
}

func (ps *ProtocolService) GenerateProtocolNumber() (string, error) {
	var protocol models.Protocol
	var protocolNumber string

	if err := ps.DB.Create(&protocol).Error; err != nil {
		return "", err
	}

	// Gere o número do protocolo amigável
	protocolNumber = fmt.Sprintf("P%08d", protocol.ID)

	// Atualize o número do protocolo na tabela Protocol
	if err := ps.DB.Model(&protocol).Update("protocol_number", protocolNumber).Error; err != nil {
		return "", err
	}

	return protocolNumber, nil
}
