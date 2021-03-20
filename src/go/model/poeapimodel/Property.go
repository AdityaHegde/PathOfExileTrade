package poeapimodel

// Property is exported
type Property struct {
	Name        string      `json:"name"`
	Values      interface{} `json:"values"`
	DisplayMode int         `json:"displayMode"`
	Type        int         `json:"type"`
	Progress    float32     `json:"progress"`
}
