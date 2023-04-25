package categoryentity

import (
	"errors"
)

var (
	ErrorCategoryIsBlank          = errors.New("category cannot be blank")
	ErrorCategoryInvalid          = errors.New("category is not valid")
	ErrorCategoryAlreadyExisted   = errors.New("category is already existed")
	ErrorCategoryDeleted          = errors.New("category has been deleted")
	ErrorCategoryNotFound         = errors.New("category not found")
	ErrorCannotCreateCategory     = errors.New("cannot create category")
	ErrorCannotCreateCategoryList = errors.New("cannot create category list")
	ErrorCannotGetCategory        = errors.New("cannot get category")
	ErrorCannotGetListCategory    = errors.New("cannot get category list")
	ErrorCannotUpdateCategory     = errors.New("cannot update category")
	ErrorCannotDeleteCategory     = errors.New("cannot delete category")
	ErrorRequesterIsNotAdmin      = errors.New("no permission, only owner can do this")
)
