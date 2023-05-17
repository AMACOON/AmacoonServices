package judge

import (
	"github.com/sirupsen/logrus"
)

type JudgeService struct {
	JudgeRepo *JudgeRepository
	Logger    *logrus.Logger
}

func NewJudgeService(judgeRepo *JudgeRepository, logger *logrus.Logger) *JudgeService {
	return &JudgeService{
		JudgeRepo: judgeRepo,
		Logger:    logger,
	}
}

func (s *JudgeService) GetAllJudges() ([]Judge, error) {
	s.Logger.Infof("Service GetAllJudges")
	judges, err := s.JudgeRepo.GetAllJudges()
	if err != nil {
		s.Logger.WithError(err).Error("failed to get all judges")
		return nil, err
	}
	s.Logger.Infof("Service GetAllJudges OK")
	return judges, nil
}

func (s *JudgeService) GetJudgeByID(id uint) (*Judge, error) {
	s.Logger.Infof("Service GetJudgeByID")
	judge, err := s.JudgeRepo.GetJudgeById(id)
	if err != nil {
		s.Logger.Errorf("error fetching judge from repository: %v", err)
		return nil, err
	}
	s.Logger.Infof("Service GetJudgeByID OK")
	return judge, nil
}

func (s *JudgeService) UpdateJudge(id uint, judge *Judge) (*Judge, error) {
	s.Logger.Infof("Service UpdateJudge")

	updatedJudge, err := s.JudgeRepo.UpdateJudge(id, judge)
	if err != nil {
		s.Logger.Errorf("error updating judge in repository: %v", err)
		return nil, err
	}

	s.Logger.Infof("Service UpdateJudge OK")
	return updatedJudge, nil
}


func (s *JudgeService) DeleteJudge(id uint) error {
	s.Logger.Infof("Service DeleteJudge")
	err := s.JudgeRepo.DeleteJudge(id)
	if err != nil {
		s.Logger.Errorf("error deleting judge in repository: %v", err)
		return err
	}
	s.Logger.Infof("Service DeleteJudge OK")
	return nil
}
