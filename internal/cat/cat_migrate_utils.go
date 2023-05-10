package cat

import (
	"errors"
	"log"
	"time"

	"regexp"
	"strings"
	"unicode"

	"github.com/scuba13/AmacoonServices/internal/title"
	"gorm.io/gorm"
)

func uintPtr(n uint) *uint {
	return &n
}

func findCountryIdByCode(db *gorm.DB, countryCode string) (*uint, error) {

	if countryCode == "" || countryCode == "0" {
		return nil, nil
	}

	var countries struct {
		ID   uint `gorm:"primaryKey"`
		Code string
	}
	err := db.Table("countries").Where("code = ?", countryCode).First(&countries).Error
	if err != nil {
		return nil, err
	}

	return &countries.ID, nil
}

func getFederationID(db *gorm.DB, federationName string) (*uint, error) {

	if federationName == "" || federationName == "0" {
		return nil, nil
	}

	var federations struct {
		ID   uint `gorm:"primaryKey"`
		Name string
	}
	err := db.Table("federations").Where("name = ?", federationName).First(&federations).Error
	if err != nil {
		return nil, err
	}

	return &federations.ID, nil
}


func getBreedID(db *gorm.DB, breedName string) (uint, error) {

	if breedName == "" || breedName == "0" {
		return 0, nil
	}

	var breeds struct {
		ID        uint `gorm:"primaryKey"`
		BreedName string
	}
	err := db.Table("breeds").Where("breed_name = ?", breedName).First(&breeds).Error
	if err != nil {
		return 0, err
	}

	return breeds.ID, nil
}

func getColorID(db *gorm.DB, emsCode, breedCode string) (uint, error) {

	if emsCode == "" || emsCode == "0" || breedCode == "" || breedCode == "0" {
		return 0, nil
	}

	var color struct {
		ID        uint `gorm:"primaryKey"`
		EmsCode   string
		BreedCode string
	}

	err := db.Table("colors").Where("ems_code = ? AND breed_code = ?", emsCode, breedCode).First(&color).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, errors.New("color not found")
		}
		return 0, err
	}

	return color.ID, nil
}

func getCatteryID(db *gorm.DB, breederName string) (*uint, error) {

	if breederName == "" || breederName == "0" {
		return nil, nil
	}

	var cattery struct {
		ID   uint `gorm:"primaryKey"`
		Name string
	}

	err := db.Table("catteries").Where("name = ?", breederName).First(&cattery).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("cattery not found")
		}
		return nil, err
	}

	return &cattery.ID, nil
}

func getOwnerID(db *gorm.DB, ownerName string) (uint, error) {

	if ownerName == "" || ownerName == "0" {
		return 0, nil
	}

	var owner struct {
		ID   uint `gorm:"primaryKey"`
		Name string
	}

	err := db.Table("owners").Where("name = ?", ownerName).First(&owner).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, errors.New("owner not found")
		}
		return owner.ID, err
	}

	return owner.ID, nil
}




func cleanParentName(name string) string {
	// Remover a vírgula e o asterisco
	name = strings.ReplaceAll(name, ",", "")
	name = strings.ReplaceAll(name, "*", "")

	// Converter a string para minúsculas
	name = strings.ToLower(name)

	// Lista das palavras a serem removidas
	wordsToRemove := []string{"ch", "pr", "ic", "ip", "gic", "gip", "sc", "sp", "nw", "aw", "bw", "cew", "mw", "nsw", "sw", "ww", "jw", "dsw", "dm", "dsm", "dvm"}

	// Cria uma expressão regular a partir das palavras a serem removidas
	regexPattern := `(\b(?:` + strings.Join(wordsToRemove, "|") + `)\b\s*)`
	regex := regexp.MustCompile(regexPattern)

	// Remove as palavras do nome
	cleanedName := regex.ReplaceAllString(name, "")

	// Se os dois primeiros caracteres são letras e o terceiro é um espaço, remova-os
	if len(cleanedName) >= 3 && unicode.IsLetter(rune(cleanedName[0])) && unicode.IsLetter(rune(cleanedName[1])) && unicode.IsSpace(rune(cleanedName[2])) {
		cleanedName = cleanedName[3:]
	}

	// Se as duas primeiras letras forem "br" ou "gb", remova-as
	if len(cleanedName) >= 2 && (cleanedName[:2] == "br" || cleanedName[:2] == "gb") {
		cleanedName = cleanedName[2:]
	}

	// Remove um ponto no início da string, se houver
	if len(cleanedName) > 0 && cleanedName[0] == '.' {
		cleanedName = cleanedName[1:]
	}

	// Se o primeiro caractere for um apóstrofo, remova-o e todos os números seguintes até encontrar uma letra
	if len(cleanedName) > 0 && cleanedName[0] == '\'' {
		cleanedName = cleanedName[1:]
		if len(cleanedName) > 0 {
			cleanedName = strings.TrimLeftFunc(cleanedName, unicode.IsDigit)
		}
	}
	// Se os dois primeiros caracteres são números, remova-os
	if len(cleanedName) >= 2 && unicode.IsDigit(rune(cleanedName[0])) && unicode.IsDigit(rune(cleanedName[1])) {
		cleanedName = cleanedName[2:]
	}

	// Remover espaços no início e no fim
	cleanedName = strings.TrimSpace(cleanedName)

	// Remover espaços extras entre as palavras
	space := regexp.MustCompile(`\s+`)
	cleanedName = space.ReplaceAllString(cleanedName, " ")

	return cleanedName
}

func insertTitles(db *gorm.DB, catID uint, titleCodes []string) error {
	log.Println("Entering insertTitles")
	if len(titleCodes) == 0 {
		return nil
	}
	for _, titleCode := range titleCodes {
		var title title.Title
		result := db.Table("titles").Where("code = ?", titleCode).First(&title)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("No document found for title code: %s\n", titleCode)
			continue
		} else if result.Error != nil {
			return result.Error
		}

		titlesCat := TitlesCat{
			CatID:        catID,
			TitleID:      title.ID,
			Date:         time.Now(),
			FederationID: 0,
		}

		result = db.Create(&titlesCat)
		if result.Error != nil {
			return result.Error
		}
	}
	log.Println("Leaving insertTitles")
	return nil
}

