package judge

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"errors"
)

type JudgeRepository struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

func NewJudgeRepository(db *gorm.DB, logger *logrus.Logger) *JudgeRepository {
	return &JudgeRepository{
		DB:     db,
		Logger: logger,
	}
}

func (r *JudgeRepository) GetAllJudges() ([]Judge, error) {
	r.Logger.Infof("Repository GetAllJudges")

	var judges []Judge
	if err := r.DB.Find(&judges).Error; err != nil {
		r.Logger.WithError(err).Errorf("error finding all judges")
		return nil, err
	}

	r.Logger.Infof("Repository GetAllJudges OK")
	return judges, nil
}

func (r *JudgeRepository) GetJudgeById(id uint) (*Judge, error) {
	r.Logger.Infof("Repository GetJudgeById")

	var judge Judge
	if err := r.DB.Where("id = ?", id).First(&judge).Error; err != nil {
		r.Logger.WithError(err).Errorf("error finding judge with id %v", id)
		return nil, err
	}

	r.Logger.Infof("Repository GetJudgeById OK")
	return &judge, nil
}

func (r *JudgeRepository) UpdateJudge(id uint, updatedJudge *Judge) (*Judge, error) {
	r.Logger.Infof("Repository UpdateJudge")

	var judge Judge
	if err := r.DB.First(&judge, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.Logger.WithError(err).Errorf("error finding judge with id %v", id)
			return nil, err
		}
		return nil, err
	}

	if err := r.DB.Model(&judge).Updates(updatedJudge).Error; err != nil {
		r.Logger.WithError(err).Errorf("error updating judge with id %v", id)
		return nil, err
	}

	r.Logger.Infof("Repository UpdateJudge OK")
	return &judge, nil
}


func (r *JudgeRepository) DeleteJudge(id uint) error {
	r.Logger.Infof("Repository DeleteJudge")

	var judge Judge
	if err := r.DB.First(&judge, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.Logger.WithError(err).Errorf("error finding judge with id %v", id)
			return err
		}
		return err
	}

	if err := r.DB.Delete(&judge).Error; err != nil {
		r.Logger.WithError(err).Errorf("error delete judge with id %v", id)
		return err
	}

	r.Logger.Infof("Repository DeleteJudge OK")
	return nil
}

