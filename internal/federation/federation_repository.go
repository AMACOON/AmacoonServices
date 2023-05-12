package federation

import (
    "strconv"

    "github.com/sirupsen/logrus"
    "gorm.io/gorm"
)

type FederationRepository struct {
    DB     *gorm.DB
    Logger *logrus.Logger
}

func NewFederationRepository(db *gorm.DB, logger *logrus.Logger) *FederationRepository {
    return &FederationRepository{
        DB:     db,
        Logger: logger,
    }
}

func (r *FederationRepository) GetFederationByID(id string) (*Federation, error) {
    r.Logger.Infof("Repository GetFederationByID")
    uintID, err := strconv.ParseUint(id, 10, 64)
    if err != nil {
        r.Logger.WithError(err).Errorf("invalid id: %s", id)
        return nil, err
    }
    var federation Federation
    if err := r.DB.First(&federation, uintID).Error; err != nil {
        r.Logger.WithError(err).Errorf("failed to get federation with ID %s", id)
        return nil, err
    }
   
    r.Logger.Infof("Repository GetFederationByID OK")
    return &federation, nil
}

func (r *FederationRepository) GetAllFederations() ([]Federation, error) {
    r.Logger.Infof("Repository GetAllFederations")
    var federations []Federation
    if err := r.DB.Find(&federations).Error; err != nil {
        r.Logger.WithError(err).Errorf("failed to get all federations")
        return nil, err
    }
   
    r.Logger.Infof("Repository GetAllFederations OK")
    return federations, nil
}


