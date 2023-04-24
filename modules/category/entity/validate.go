package categoryentity

func checkEmptyContent(content string) error {
	if content == "" {
		return ErrorCategoryIsBlank
	}

	return nil
}
