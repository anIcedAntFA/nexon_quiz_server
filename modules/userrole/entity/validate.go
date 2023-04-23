package userroleentity

func checkValidContent(content UserRoleContent) error {
	for _, userRole := range allUserRoles {
		if content == userRole {
			return nil
		}
	}

	return ErrorUserRoleInvalid
}
