package questionentity

type QuestionFilter struct {
	Type       []string `json:"type_id" form:"type_id"`
	Difficulty string   `json:"difficulty_id" form:"difficulty_id"`
	Category   []string `json:"category_id" form:"category_id"`
}
