package typeentity

func checkEmptyContent(content string) error {
	if content == "" {
		return ErrorTypeIsBlank
	}

	return nil
}
