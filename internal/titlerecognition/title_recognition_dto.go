package titlerecognition

import (
	"time"

	"github.com/scuba13/AmacoonServices/internal/service"
)

type TitleRecognitionRequest struct {
	CatData        service.CatServiceRequest   `json:"catData"`
	OwnerData      service.OwnerServiceRequest `json:"ownerData"`
	Titles         []TitleRequest              `json:"titles"`
	Status         string                      `json:"status"`
	ProtocolNumber string                      `json:"protocolNumber"`
	RequesterID    string                      `json:"requesterID"`
}

type TitleRequest struct {
	TitleID     string    `json:"titleID"`
	TitleCode   string    `json:"titleCode"`
	TitleName   string    `json:"titleName"`
	Certificate string    `json:"certificate"`
	Date        time.Time `json:"date"`
	Judge       string    `json:"judge"`
}
