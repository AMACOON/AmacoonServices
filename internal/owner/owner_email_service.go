package owner

import (
	"fmt"

	"github.com/matcornic/hermes/v2"
	"github.com/scuba13/AmacoonServices/internal/utils"
	"github.com/sirupsen/logrus"
	"strings"
	"github.com/spf13/viper"
)

type OwnerEmailService struct {
	SmtpService *utils.SmtpService
	Logger      *logrus.Logger
}

func NewOwnerEmailService(smtpService *utils.SmtpService, logger *logrus.Logger) *OwnerEmailService {
	return &OwnerEmailService{
		SmtpService: smtpService,
		Logger:      logger,
	}
}

func (s *OwnerEmailService) SendOwnerValidationEmail(owner *Owner) error {
	s.Logger.Infof("Sending Owner Validation Email %s", owner.Email) // trocar para env from

	from := viper.GetString("smtp.systemEmail")
	to := []string{viper.GetString("smtp.adminEmail")}
	subject := fmt.Sprintf("Owner Validation Request: %s", owner.Name)
	url := viper.GetString("server.host")
	
	// Consulta para obter os nomes dos clubes
	clubNames:= getClubNamesSeparatedByComma(*owner)

	// Configura o Hermes
	h := hermes.Hermes{
		Theme: new(hermes.Default),
	}

	email := hermes.Email{
		Body: hermes.Body{
			Name: "CatClubSystem Admin",
			Title: "New Owner Validation Request",
			Intros: []string{
				fmt.Sprintf("Name: %s", owner.Name),
				fmt.Sprintf("Email: %s", owner.CPF),
				fmt.Sprintf("CPF: %s", owner.CPF),
				fmt.Sprintf("Address: %s", owner.Address),
				fmt.Sprintf("City: %s", owner.City),
				fmt.Sprintf("State: %s", owner.State),
				fmt.Sprintf("ZipCode: %s", owner.ZipCode),
				fmt.Sprintf("Country: %s", owner.Country.Name),
				fmt.Sprintf("Phone: %s", owner.Phone),
				fmt.Sprintf("Clubs: %s", clubNames),
				
			},
			Actions: []hermes.Action{
				{
					Instructions: fmt.Sprintf("Click here to validate the owner: %s", owner.Name),
					Button: hermes.Button{
						Color: "#000000", // Cor preta
						Text:  "Validate New Owner",
						Link:  fmt.Sprintf("%s/api/owners/%d/%s/valid", url, owner.ID, owner.ValidId),
					},
				},
			},
			Outros: []string{
				"For more information, please access the CatClubSystem administration page.",
				"http://br.catclubsystem.com/index.php/adm",
				fmt.Sprintf("%s/api/owners/%d/%s/valid", url, owner.ID, owner.ValidId),
			},
		},
	}

	// Gera o email em HTML
	body, err := h.GenerateHTML(email)
	if err != nil {
		s.Logger.Errorf("Failed to render email template: %v", err)
		return err
	}

	err = s.SmtpService.SendEmail(from, to, subject, body)
	if err != nil {
		s.Logger.WithError(err).Error("Failed to send validation email")
		return err
	}

	return nil
}

func (s *OwnerEmailService) SendOwnerValidationConfirmationEmail(owner *Owner) error {
	s.Logger.Infof("Sending Owner Validation Confirmation Email %s", owner.Email) // trocar para env from

	from := "sistema@catclubsystem.com"
	to := []string{owner.Email}
	subject := fmt.Sprintf("Owner Confirmation Catclubsystem: %s", owner.Name)

	// Consulta para obter os nomes dos clubes
	clubNames:= getClubNamesSeparatedByComma(*owner)


	// Configura o Hermes
	h := hermes.Hermes{
		Theme: new(hermes.Default),
	}

	email := hermes.Email{
		Body: hermes.Body{
			Name: owner.Name,
			Title: "Your registration has been successfully validated",
			Intros: []string{
				fmt.Sprintf("Name: %s", owner.Name),
				fmt.Sprintf("Email: %s", owner.CPF),
				fmt.Sprintf("CPF: %s", owner.CPF),
				fmt.Sprintf("Address: %s", owner.Address),
				fmt.Sprintf("City: %s", owner.City),
				fmt.Sprintf("State: %s", owner.State),
				fmt.Sprintf("ZipCode: %s", owner.ZipCode),
				fmt.Sprintf("Country: %s", owner.Country.Name),
				fmt.Sprintf("Phone: %s", owner.Phone),
				fmt.Sprintf("Clubs: %s", clubNames),
				
			},
			Actions: []hermes.Action{
				{
					Instructions: fmt.Sprintf("Click here to go to the website.: %s", owner.Name),
					Button: hermes.Button{
						Color: "#000000", // Cor preta
						Text:  "Go to site",
						Link:  fmt.Sprintf("http://br.catclubsystem.com/"),
					},
				},
			},
			Outros: []string{
				"Thank you for using CatSystemClub.",
			},
		},
	}

	// Gera o email em HTML
	body, err := h.GenerateHTML(email)
	if err != nil {
		s.Logger.Errorf("Failed to render email template: %v", err)
		return err
	}

	err = s.SmtpService.SendEmail(from, to, subject, body)
	if err != nil {
		s.Logger.WithError(err).Error("Failed to send validation email")
		return err
	}

	return nil
}



func getClubNamesSeparatedByComma(owner Owner) string {
	var clubNames []string

	for _, ownerClub := range owner.Clubs {
		if ownerClub.Club != nil {
			clubNames = append(clubNames, ownerClub.Club.Name)
		}
	}

	return strings.Join(clubNames, ", ")
}

