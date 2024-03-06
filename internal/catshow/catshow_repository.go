package catshow

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CatShowRepository struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

func NewCatShowRepository(db *gorm.DB, logger *logrus.Logger) *CatShowRepository {
	return &CatShowRepository{
		DB:     db,
		Logger: logger,
	}
}

func (r *CatShowRepository) CreateCatShow(catShow *CatShow) (*CatShow, error) {
	r.Logger.Infof("Repository CreateCatShow")
	if err := r.DB.Create(catShow).Error; err != nil {
		return nil, err
	}
	r.Logger.Infof("Repository CreateCatShow OK")
	return catShow, nil
}

func (r *CatShowRepository) GetCatShowByID(id uint) (*CatShow, error) {
	r.Logger.Infof("Repository GetCatShowByID")
	var catShow CatShow
	if err := r.DB.Preload("Federation").
				   Preload("Federation.Country").
				   Preload("Club").
				   Preload("Country").
				   Preload("CatShowSubs").
				   Preload("CatShowJudges.Judge").
				   Preload("CatShowJudges.Judge.Country").
				   First(&catShow, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.Logger.Errorf("Repository GetCatShowByID not found")
			return nil, nil
		}
		return nil, err
	}
	r.Logger.Infof("Repository GetCatShowByID OK")
	return &catShow, nil
}

func (r *CatShowRepository) UpdateCatShow(id uint, updatedCatShow *CatShow) error {
    r.Logger.Infof("Repository UpdateCatShow")

    // Localiza o registro do CatShow pelo ID
    catShow := CatShow{}
    if err := r.DB.First(&catShow, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            r.Logger.Errorf("No CatShow found with id: %d", id)
            return err
        }
        r.Logger.Errorf("Failed to retrieve CatShow with id %d: %v", id, err)
        return err
    }
    r.Logger.Infof("By ID: %d", catShow.ID)

    // Inicia uma nova transação no banco de dados
    tx := r.DB.Begin()

    // Atualiza os campos de CatShow
    if err := tx.Model(&catShow).Updates(updatedCatShow).Error; err != nil {
        tx.Rollback()
        r.Logger.Errorf("Update CatShow failed: %v", err)
        return err
    }

    r.Logger.Infof("CatShow fields updated successfully")

    // Atualiza CatShowSubs relacionados, se não for nulo
    if updatedCatShow.CatShowSubs != nil {
        for _, sub := range updatedCatShow.CatShowSubs {
            // Popula o CatShowID corretamente
            sub.CatShowID = catShow.ID
            // CatShowSub existente, atualiza
            if err := tx.Model(&CatShowSub{}).Where("id = ?", sub.ID).Updates(&sub).Error; err != nil {
                tx.Rollback()
                r.Logger.Errorf("Update CatShowSub failed: %v", err)
                return err
            }
            r.Logger.Infof("CatShowSub fields updated successfully")
        }
    }

    // Atualiza CatShowJudges relacionados, se não for nulo
    if updatedCatShow.CatShowJudges != nil {
        for _, judge := range updatedCatShow.CatShowJudges {
            // Popula o CatShowID corretamente
            judge.CatShowID = catShow.ID
            // CatShowJudge existente, atualiza
            if err := tx.Model(&CatShowJudge{}).Where("id = ?", judge.ID).Updates(&judge).Error; err != nil {
                tx.Rollback()
                r.Logger.Errorf("Update CatShowJudge failed: %v", err)
                return err
            }
            r.Logger.Infof("CatShowJudge fields updated successfully")
        }
    }

    // Commit da transação
    if err := tx.Commit().Error; err != nil {
        r.Logger.Errorf("Transaction commit failed: %v", err)
        return err
    }

    r.Logger.Infof("Repository UpdateCatShow OK")

    return nil
}












