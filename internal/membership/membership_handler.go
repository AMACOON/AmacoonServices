package membership

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{Service: service}
}

// POST /api/membership-requests
// Recebe form-data.
// Campos esperados (por enquanto):
// - association_type
// - owner[name], owner[last_name], owner[cpf], owner[email], owner[phone]
// - owner[address], owner[city], owner[state], owner[zipcode]
// - has_fife_cattery (0|1)
// - observation
func (h *Handler) Create(c echo.Context) error {

	hasFife := false
	if v := c.FormValue("has_fife_cattery"); v != "" {
		// Aceita 1/0 e true/false
		if v == "1" || v == "true" || v == "True" {
			hasFife = true
		}
	}

	req := &MembershipRequest{
		AssociationType: c.FormValue("association_type"),

		OwnerName:     c.FormValue("owner[name]"),
		OwnerLastName: c.FormValue("owner[last_name]"),
		CPF:           c.FormValue("owner[cpf]"),
		Email:         c.FormValue("owner[email]"),
		Phone:         c.FormValue("owner[phone]"),
		Address:       c.FormValue("owner[address]"),
		City:          c.FormValue("owner[city]"),
		State:         c.FormValue("owner[state]"),
		ZipCode:       c.FormValue("owner[zipcode]"),

		HasFifeCattery: hasFife,
		Observation:    c.FormValue("observation"),
	}

	// (Opcional) leitura rápida de gatos enviados como cat_count + cat[i][...]
	// Isso evita travar o front se ele já mandar gatos.
	if countStr := c.FormValue("cat_count"); countStr != "" {
		if n, err := strconv.Atoi(countStr); err == nil && n > 0 && n < 50 {
			for i := 0; i < n; i++ {
				name := c.FormValue("cat[" + strconv.Itoa(i) + "][name]")
				birth := c.FormValue("cat[" + strconv.Itoa(i) + "][birth_date]")
				sex := c.FormValue("cat[" + strconv.Itoa(i) + "][sex]")
				breedIDStr := c.FormValue("cat[" + strconv.Itoa(i) + "][breed_id]")

				var breedID uint
				if breedIDStr != "" {
					if bi, err := strconv.Atoi(breedIDStr); err == nil && bi > 0 {
						breedID = uint(bi)
					}
				}

				if name != "" || birth != "" || sex != "" || breedID != 0 {
					req.Cats = append(req.Cats, MembershipCat{
						Name:      name,
						BirthDate: birth,
						Sex:       sex,
						BreedID:   breedID,
					})
				}
			}
		}
	}

	if err := h.Service.Create(req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"protocol": req.ProtocolCode,
		"status":   req.Status,
	})
}
