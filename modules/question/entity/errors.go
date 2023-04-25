package questionentity

import (
	"errors"
)

var (
	ErrorQuestionContentIsBlank    = errors.New("question content cannot be blank")
	ErrorQuestionTypeIsBlank       = errors.New("question type cannot be blank")
	ErrorQuestionDifficultyIsBlank = errors.New("question difficulty cannot be blank")
	ErrorQuestionCategoryIsBlank   = errors.New("question category cannot be blank")
	ErrorQuestionInvalid           = errors.New("question is not valid")
	ErrorQuestionAlreadyExisted    = errors.New("question is already existed")
	ErrorQuestionDeleted           = errors.New("question has been deleted")
	ErrorQuestionNotFound          = errors.New("question not found")
	ErrorCannotCreateQuestion      = errors.New("cannot create question")
	ErrorCannotCreateQuestionList  = errors.New("cannot create question list")
	ErrorCannotGetQuestion         = errors.New("cannot get question")
	ErrorCannotGetListQuestion     = errors.New("cannot get question list")
	ErrorCannotUpdateQuestion      = errors.New("cannot update question")
	ErrorCannotDeleteQuestion      = errors.New("cannot delete question")
	ErrorRequesterIsNotAdmin       = errors.New("no permission, only owner can do this")
)
