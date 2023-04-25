package difficultyentity

func checkEmptyContent(content string) error {
	if content == "" {
		return ErrorDifficultyIsBlank
	}

	return nil
}
