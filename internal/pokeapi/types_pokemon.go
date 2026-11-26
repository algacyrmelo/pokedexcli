package pokeapi

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		Stat     NamedAPIResource `json:"stat"`
		Effort   int              `json:"effort"`
		BaseStat int              `json:"base_stat"`
	} `json:"stats"`
	Types []struct {
		Type NamedAPIResource `json:"type"`
		Slot int              `json:"slot"`
	} `json:"types"`
}
