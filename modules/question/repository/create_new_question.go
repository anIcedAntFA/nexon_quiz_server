package questionrepository

import (
	"context"
	"nexon_quiz/common"
	categoryentity "nexon_quiz/modules/category/entity"
	difficultyentity "nexon_quiz/modules/difficulty/entity"
	questionentity "nexon_quiz/modules/question/entity"
	typeentity "nexon_quiz/modules/type/entity"

	"gorm.io/gorm"
)

type CreateQuestionStorage interface {
	FindQuestionByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*questionentity.Question, error)

	InsertNewQuestion(
		ctx context.Context,
		newQuestion *questionentity.QuestionCreate,
	) error
}

type FindTypeStorage interface {
	FindTypeByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*typeentity.Type, error)
}

type FindDifficultyStorage interface {
	FindDifficultyByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*difficultyentity.Difficulty, error)
}

type FindCategoryStorage interface {
	FindCategoryByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*categoryentity.Category, error)
}

type createQuestionRepository struct {
	questionStorage   CreateQuestionStorage
	typeStorage       FindTypeStorage
	difficultyStorage FindDifficultyStorage
	categoryStorage   FindCategoryStorage
}

func NewCreateQuestionRepository(
	questionStorage CreateQuestionStorage,
	typeStorage FindTypeStorage,
	difficultyStorage FindDifficultyStorage,
	categoryStorage FindCategoryStorage,
) *createQuestionRepository {
	return &createQuestionRepository{
		questionStorage:   questionStorage,
		typeStorage:       typeStorage,
		difficultyStorage: difficultyStorage,
		categoryStorage:   categoryStorage,
	}
}

func (repo *createQuestionRepository) CreateNewQuestion(
	ctx context.Context,
	newQuestion *questionentity.QuestionCreate,
) error {
	oldQuestion, err := repo.questionStorage.FindQuestionByCondition(
		ctx,
		map[string]interface{}{"content": newQuestion.Content},
	)

	if err == nil && newQuestion.Content == oldQuestion.Content {
		return common.NewCustomError(
			err,
			questionentity.ErrorQuestionAlreadyExisted.Error(),
			"ErrorQuestionAlreadyExisted",
		)
	}

	_, err = repo.typeStorage.FindTypeByCondition(
		ctx,
		map[string]interface{}{"id": newQuestion.TypeId},
	)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return typeentity.ErrorTypeNotFound
		}

		return common.NewCustomError(
			err,
			typeentity.ErrorTypeInvalid.Error(),
			"ErrorQuestionTypeInvalid",
		)
	}

	_, err = repo.difficultyStorage.FindDifficultyByCondition(
		ctx,
		map[string]interface{}{"id": newQuestion.DifficultyId},
	)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return difficultyentity.ErrorDifficultyNotFound
		}

		return common.NewCustomError(
			err,
			difficultyentity.ErrorDifficultyInvalid.Error(),
			"ErrorQuestionDifficultyInvalid",
		)
	}

	_, err = repo.categoryStorage.FindCategoryByCondition(
		ctx,
		map[string]interface{}{"id": newQuestion.CategoryId},
	)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return categoryentity.ErrorCategoryNotFound
		}

		return common.NewCustomError(
			err,
			categoryentity.ErrorCategoryInvalid.Error(),
			"ErrorQuestionCategoryInvalid",
		)
	}

	if err := repo.questionStorage.InsertNewQuestion(ctx, newQuestion); err != nil {
		return common.NewCustomError(
			err,
			questionentity.ErrorCannotCreateQuestion.Error(),
			"ErrorCannotCreateQuestion",
		)
	}

	return nil
}
