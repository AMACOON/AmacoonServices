package membership

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

func NewRepository(db *gorm.DB, logger *logrus.Logger) *Repository {
	return &Repository{DB: db, Logger: logger}
}

func (r *Repository) Create(req *MembershipRequest) error {
	r.Logger.Info("Repository Membership Create")
	return r.DB.Create(req).Error
}

func (r *Repository) GetByProtocol(protocol string) (*MembershipRequest, error) {
	r.Logger.Info("Repository Membership GetByProtocol")

	var req MembershipRequest
	err := r.DB.
		Preload("Cats").
		Preload("Files").
		Where("protocol_code = ?", protocol).
		First(&req).Error

	if err != nil {
		return nil, err
	}
	return &req, nil
}
