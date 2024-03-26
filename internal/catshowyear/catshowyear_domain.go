package catshowyear

import (
	"time"

	"github.com/scuba13/AmacoonServices/internal/catshowregistration"
	"github.com/scuba13/AmacoonServices/internal/catshowresult"
)

type CatShowComplete struct {
	Registration  catshowregistration.Registration `json:"registration"`
	CatShowResult catshowresult.CatShowResult      `json:"result"`
}

// CatShowYearGroup agrupa os CatShowDetails por ano.
type CatShowYearGroup struct {
    Year     int             `json:"-"`
    CatShows []CatShowDetail `json:"catShows"`
}


// CatShowDetail representa detalhes de um CatShow, incluindo suas subdivisões e resultados relacionados.
type CatShowDetail struct {
	CatShowYear        int
	CatShowDescription string
	CatShowCity        string
	CatShowSubs        []CatShowSubDetail
}

// CatShowSubDetail combina informações de uma subexposição, resultados e detalhes do gato.
type CatShowSubDetail struct {
	CatShowSubDescription string
	CatShowSubCatShowDate time.Time
	ResultsDescription    string
	ResultScore           int
	ClassDescription      string
	ClassName             string
	ClassCode             string
	JudgesName            string
	CatShowCatNameFull    string
	CatShowCatBreedCode   string
	CatShowCatEmsCode     string
	CatShowCatGender      string
	CatShowCatBirthdate   time.Time
	OwnerName             string
}
