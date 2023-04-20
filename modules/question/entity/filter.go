package questionentity

type Filter struct {
	Category   string `json:"category" form:"category"`
	Type       string `json:"type" form:"type"`
	Difficulty string `json:"difficulty" form:"difficulty"`
}
