package title

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func InsertTitles(db *gorm.DB, logger *logrus.Logger) {
	titles := []Title{
		{Name: "Champion", Code: "CH", Type: "Championship/Premiorship Titles", Certificate: "CAC", Amount: 3, Observation: "3 juízes diferentes"},
		{Name: "Premior", Code: "PR", Type: "Championship/Premiorship Titles", Certificate: "CAP", Amount: 3, Observation: "3 juízes diferentes"},
		{Name: "International Champion", Code: "IC", Type: "Championship/Premiorship Titles", Certificate: "CACIB", Amount: 3, Observation: "3 juízes diferentes"},
		{Name: "International Premior", Code: "IP", Type: "Championship/Premiorship Titles", Certificate: "CAPIB", Amount: 3, Observation: "3 juízes diferentes"},
		{Name: "Grand International Champion", Code: "GIC", Type: "Championship/Premiorship Titles", Certificate: "CAGCIB", Amount: 6, Observation: "4 juízes diferentes"},
		{Name: "Grand International Premior", Code: "GIP", Type: "Championship/Premiorship Titles", Certificate: "CAGPIB", Amount: 6, Observation: "4 juízes diferentes"},
		{Name: "Supreme Champion", Code: "SC", Type: "Championship/Premiorship Titles", Certificate: "CACS", Amount: 9, Observation: "5 juízes diferentes"},
		{Name: "Supreme Premior", Code: "SP", Type: "Championship/Premiorship Titles", Certificate: "CAPS", Amount: 9, Observation: "5 juízes diferentes"},
		{Name: "Junior Winner", Code: "JW", Type: "Winner Titles", Certificate: "", Amount: 0, Observation: ""},
		{Name: "Distinguised Senior Winner", Code: "DSW", Type: "Winner Titles", Certificate: "", Amount: 0, Observation: ""},
		{Name: "National Winner", Code: "NW", Type: "Winner Titles", Certificate: "", Amount: 0, Observation: "Fornecido pela Federação."},
		{Name: "American Winner", Code: "AW", Type: "Winner Titles", Certificate: "BIS", Amount: 1, Observation: "Anexar o Certificado que comprove Best in Show no campeonato."},
		{Name: "Baltic Winner", Code: "BW", Type: "Winner Titles", Certificate: "BIS", Amount: 1, Observation: "Anexar o Certificado que comprove Best in Show no campeonato."},
		{Name: "Central European Winner", Code: "CEW", Type: "Winner Titles", Certificate: "BIS", Amount: 1, Observation: "Anexar o Certificado que comprove Best in Show no campeonato."},
		{Name: "Mediterranean Winner", Code: "MW", Type: "Winner Titles", Certificate: "BIS", Amount: 1, Observation: "Anexar o Certificado que comprove Best in Show no campeonato."},
		{Name: "North Sea Winner", Code: "NSW", Type: "Winner Titles", Certificate: "BIS", Amount: 1, Observation: "Anexar o Certificado que comprove Best in Show no campeonato."},
		{Name: "Scandinavian Winner", Code: "SW", Type: "Winner Titles", Certificate: "BIS", Amount: 1, Observation: "Anexar o Certificado que comprove Best in Show no campeonato."},
		{Name: "World Winner", Code: "WW", Type: "Winner Titles", Certificate: "BIS", Amount: 1, Observation: "Anexar o Certificado que comprove Best in Show no campeonato."},
		{Name: "Distinguished Merit", Code: "DM", Type: "Merit Titles", Certificate: "IC/IP para DM", Amount: 5, Observation: "5 certificados de filhos com o título IC/IP ou superior."},
		{Name: "Distinguished Show Merit", Code: "DSM", Type: "Merit Titles", Certificate: "BIS para DSM", Amount: 10, Observation: "Em no mínimo 2 anos"},
		{Name: "Distinguished Variety Merit", Code: "DVM", Type: "Merit Titles", Certificate: "BIV para DVM", Amount: 10, Observation: "Em no mínimo 2 anos"},
	}

	for _, title := range titles {
		result := db.Create(&title)
		if result.Error != nil {
			logger.Infof("Error inserting title '%s': %v\n", title.Name, result.Error)
		} else {
			logger.Infof("Inserted title '%s' with ID %d\n", title.Name, title.ID)
		}
	}
}
