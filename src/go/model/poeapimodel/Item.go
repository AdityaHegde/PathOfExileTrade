package poeapimodel

// Item is exported
type Item struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Category  interface{} `json:"category"`
	TypeLine  string      `json:"typeLine"`
	DescrText string      `json:"descrText"`

	CraftedMods          []string   `json:"craftedMods"`
	EnchantMods          []string   `json:"enchantMods"`
	ExplicitMods         []string   `json:"explicitMods"`
	ImplicitMods         []string   `json:"implicitMods"`
	Properties           []Property `json:"properties"`
	AdditionalProperties []Property `json:"additionalProperties"`

	X      uint   `json:"x"`
	Y      uint   `json:"y"`
	Icon   string `json:"icon"`
	League string `json:"league"`
}
