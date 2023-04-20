package questionentity

type Filter struct {
	Category   string             `json:"category" form:"category"`
	Type       []QuestionType     `json:"type" form:"type"`
	Difficulty QuestionDifficulty `json:"difficulty" form:"difficulty"`
}
