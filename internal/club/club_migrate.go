package club

import (
	"gorm.io/gorm"
	"github.com/sirupsen/logrus"
)

type ClubS struct {
	gorm.Model
	ID         int    `gorm:"column:id_clubes;primaryKey"`
	Name       string `gorm:"column:nome"`
	Nickname   string `gorm:"column:apelido"`
	Email      string `gorm:"column:email"`
	Login      string `gorm:"column:login"`
	Password   string `gorm:"column:senha"`
	Permission string `gorm:"column:permissao"`
}

func (ClubS) TableName() string {
	return "clubes"
}


func MigrateClubs(dbOld *gorm.DB, dbNew *gorm.DB, logger *logrus.Logger) error {
    logger.Infof("Migrating clubs...")

    var clubsSource []ClubS
    if err := dbOld.Unscoped().Find(&clubsSource).Error; err != nil {
        logger.WithError(err).Error("Failed to get clubs from old database")
        return err
    }

    var clubsDestination []Club
    for _, clubSource := range clubsSource {
        clubDestination := Club{
            Name:       clubSource.Name,
            Nickname:   clubSource.Nickname,
            Email:      clubSource.Email,
            Login:      clubSource.Login,
            Password:   clubSource.Password,
            Permission: convertPermission(clubSource.Permission), // Convert permission value, if necessary
        }
        clubsDestination = append(clubsDestination, clubDestination)
    }

    if err := dbNew.Create(&clubsDestination).Error; err != nil {
        logger.WithError(err).Error("Failed to migrate clubs to new database")
        return err
    }

    logger.Infof("Clubs migration completed successfully")
    return nil
}

func convertPermission(permission string) int {
    switch permission {
    case "1":
        return 1
    case "2":
        return 2
    case "3":
        return 3
    default:
        return 0
    }
}
