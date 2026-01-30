package membership

import (
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Service struct {
	Repo   *Repository
	Logger *logrus.Logger
}

func NewService(repo *Repository, logger *logrus.Logger) *Service {
	return &Service{
		Repo:   repo,
		Logger: logger,
	}
}

func (s *Service) Create(req *MembershipRequest) error {
	s.Logger.Info("Service Membership Create")

	req.ProtocolCode = uuid.New().String()
	req.Status = "PENDING"
	req.SubmittedAt = time.Now()

	return s.Repo.Create(req)
}
