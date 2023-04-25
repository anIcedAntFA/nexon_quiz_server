package difficultyentity

import (
	"errors"
)

var (
	ErrorDifficultyIsBlank          = errors.New("difficulty cannot be blank")
	ErrorDifficultyInvalid          = errors.New("difficulty is not valid")
	ErrorDifficultyAlreadyExisted   = errors.New("difficulty is already existed")
	ErrorDifficultyDeleted          = errors.New("difficulty has been deleted")
	ErrorDifficultyNotFound         = errors.New("difficulty not found")
	ErrorCannotCreateDifficulty     = errors.New("cannot create difficulty")
	ErrorCannotCreateDifficultyList = errors.New("cannot create difficulty list")
	ErrorCannotGetDifficulty        = errors.New("cannot get difficulty")
	ErrorCannotGetListDifficulty    = errors.New("cannot get difficulty list")
	ErrorCannotUpdateDifficulty     = errors.New("cannot update difficulty")
	ErrorCannotDeleteDifficulty     = errors.New("cannot delete difficulty")
	ErrorRequesterIsNotRootAdmin    = errors.New("no permission, only root admin can do this")
)
