package catshowcomplete

import (
	"github.com/scuba13/AmacoonServices/internal/catshowregistration"
	"github.com/scuba13/AmacoonServices/internal/catshowresult"	
)

type CatShowComplete struct {
	Registration catshowregistration.Registration `json:"registration"`
	CatShowResult catshowresult.CatShowResult `json:"result"`
}