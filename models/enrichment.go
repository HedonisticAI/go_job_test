package models

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

type Age_Response struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

type Gender_Response struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
}

type country struct {
	CountryID   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}
type Nationality_Response struct {
	Count   int       `json:"count"`
	Name    string    `json:"name"`
	Country []country `json:"country"`
}

const age_api = "https://api.agify.io/?name="
const gender_api = "https://api.genderize.io/?name="
const nationality_api = "https://api.nationalize.io/?name="

func get_gender(Name string) string {
	res, err := http.Get(gender_api + Name)
	if err != nil {
		slog.Debug("Error getting Data from %s with name %s "+err.Error(), gender_api, Name)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		slog.Debug("Error reading respone body from %s with name %s "+err.Error(), gender_api, Name)
		return ""
	}
	var result Gender_Response
	if err := json.Unmarshal(body, &result); err != nil {
		slog.Debug("Error during Unmarshalling" + err.Error())
		return ""
	}
	return result.Gender

}

func get_nation(Name string) string {
	res, err := http.Get(nationality_api + Name)
	if err != nil {
		slog.Debug("Error getting Data from %s with name %s "+err.Error(), nationality_api, Name)
		return ""
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		slog.Debug("Error reading respone body from %s with name %s "+err.Error(), nationality_api, Name)
		return ""
	}
	var national_data Nationality_Response
	if err := json.Unmarshal(body, &national_data); err != nil {
		slog.Debug("Error during Unmarshalling" + err.Error())
		return ""
	}
	result := national_data.Country[0]
	for _, iter := range national_data.Country {
		if iter.Probability > result.Probability {
			result = iter
		}
	}
	return result.CountryID
}

func get_age(Name string) int {
	res, err := http.Get(age_api + Name)
	if err != nil {
		slog.Debug("Error getting Data from %s with name %s "+err.Error(), age_api, Name)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		slog.Debug("Error reading respone body from %s with name %s "+err.Error(), age_api, Name)
		return -1
	}
	var result Age_Response
	if err := json.Unmarshal(body, &result); err != nil {
		slog.Debug("Error during Unmarshalling" + err.Error())
		return -1
	}
	return result.Age
}

func (h *Human) Enrich() {
	h.Gender = get_gender(h.Name)

	h.Nationality = get_nation(h.Name)

	h.Age = get_age(h.Name)

}
