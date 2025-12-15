package pokeapi

type Pokemon struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	BaseExperience     int    `json:"base_experience"`
	Height             int    `json:"height"`
	Weight             int    `json:"weight"`
	IsDefault          bool   `json:"is_default"`
	Order              int    `json:"order"`
	LocationEncounters string `json:"location_area_encounters"`

	Abilities   []AbilitySlot `json:"abilities"`
	Forms       []NamedAPI    `json:"forms"`
	GameIndices []GameIndex   `json:"game_indices"`
	HeldItems   []HeldItem    `json:"held_items"`
	Moves       []MoveSlot    `json:"moves"`
	Stats       []StatSlot    `json:"stats"`
	Types       []TypeSlot    `json:"types"`
	PastTypes   []PastType    `json:"past_types"`

	Species NamedAPI `json:"species"`
	Cries   Cries    `json:"cries"`
}

type NamedAPI struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type AbilitySlot struct {
	IsHidden bool      `json:"is_hidden"`
	Slot     int       `json:"slot"`
	Ability  *NamedAPI `json:"ability"`
}

type GameIndex struct {
	GameIndex int      `json:"game_index"`
	Version   NamedAPI `json:"version"`
}

type HeldItem struct {
	Item NamedAPI `json:"item"`
}

type MoveSlot struct {
	Move NamedAPI `json:"move"`
}

type StatSlot struct {
	BaseStat int      `json:"base_stat"`
	Effort   int      `json:"effort"`
	Stat     NamedAPI `json:"stat"`
}

type TypeSlot struct {
	Slot int      `json:"slot"`
	Type NamedAPI `json:"type"`
}

type PastType struct {
	Generation NamedAPI `json:"generation"`
}

type Cries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}
