package typeentity

import (
	"errors"
)

var (
	ErrorTypeIsBlank             = errors.New("type cannot be blank")
	ErrorTypeInvalid             = errors.New("type is not valid")
	ErrorTypeAlreadyExisted      = errors.New("type is already existed")
	ErrorTypeDeleted             = errors.New("type has been deleted")
	ErrorTypeNotFound            = errors.New("type not found")
	ErrorCannotCreateType        = errors.New("cannot create type")
	ErrorCannotCreateTypeList    = errors.New("cannot create type list")
	ErrorCannotGetType           = errors.New("cannot get type")
	ErrorCannotGetListType       = errors.New("cannot get type list")
	ErrorCannotUpdateType        = errors.New("cannot update type")
	ErrorCannotDeleteType        = errors.New("cannot delete type")
	ErrorRequesterIsNotRootAdmin = errors.New("no permission, only root admin can do this")
)
