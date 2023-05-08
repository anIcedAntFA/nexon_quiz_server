package questionbusiness

import (
	"context"
	"nexon_quiz/common"
	questionentity "nexon_quiz/modules/question/entity"

	"github.com/google/uuid"
)

type FindPlayingQuestionListRepository interface {
	FindPlayingQuestionList(
		ctx context.Context,
		quantity int,
		typeSettingIds []uuid.UUID,
		difficultySettingId uuid.UUID,
		categorySettingIds []uuid.UUID,
		moreKeys ...string,
	) ([]questionentity.Question, error)
}

type playingQuestionListBusiness struct {
	storage FindPlayingQuestionListRepository
}

func NewPlayingQuestionListBusiness(
	storage FindPlayingQuestionListRepository,
) *playingQuestionListBusiness {
	return &playingQuestionListBusiness{
		storage: storage,
	}
}

func (biz *playingQuestionListBusiness) GetPlayingQuestionList(
	ctx context.Context,
	quantity int,
	typeSettingIds []uuid.UUID,
	difficultySettingId uuid.UUID,
	categorySettingIds []uuid.UUID,
	moreKeys ...string,
) ([]questionentity.Question, error) {

	questions, err := biz.storage.FindPlayingQuestionList(
		ctx,
		quantity,
		typeSettingIds,
		difficultySettingId,
		categorySettingIds,
	)

	if err != nil {
		return nil, common.ErrorCannotListEntity(questionentity.EntityName, err)
	}

	return questions, nil
}
