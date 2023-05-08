package questionrepository

import (
	"context"
	"log"
	"nexon_quiz/common"
	gamesettingentity "nexon_quiz/modules/gamesetting/entity"
	questionentity "nexon_quiz/modules/question/entity"

	"github.com/google/uuid"
)

type FindPlayingQuestionListStorage interface {
	FindPlayingQuestionList(
		ctx context.Context,
		quantity int,
		typeSettingIds []uuid.UUID,
		difficultySettingId uuid.UUID,
		categorySettingIds []uuid.UUID,
		moreKeys ...string,
	) ([]questionentity.Question, error)
}

type FindGameSettingStorage interface {
	FindGameSettingByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*gamesettingentity.GameSetting, error)
}

type playingQuestionListRepository struct {
	playingQuestionStorage FindPlayingQuestionListStorage
	gameSettingStorage     FindGameSettingStorage
	requester              common.Requester
}

func NewPlayingQuestionListRepository(
	playingQuestionStorage FindPlayingQuestionListStorage,
	gameSettingStorage FindGameSettingStorage,
	requester common.Requester,
) *playingQuestionListRepository {
	return &playingQuestionListRepository{
		playingQuestionStorage: playingQuestionStorage,
		gameSettingStorage:     gameSettingStorage,
		requester:              requester,
	}
}

func (repo *playingQuestionListRepository) FindPlayingQuestionList(
	ctx context.Context,
	quantity int,
	typeSettingIds []uuid.UUID,
	difficultySettingId uuid.UUID,
	categorySettingIds uuid.UUID,
	moreKeys ...string,
) ([]questionentity.Question, error) {
	gameSetting, err := repo.gameSettingStorage.FindGameSettingByCondition(
		ctx,
		map[string]interface{}{"user_id": repo.requester.GetUserId()},
	)

	log.Println("gameSetting", gameSetting)

	questions, err := repo.playingQuestionStorage.FindPlayingQuestionList(
		ctx,
		16,
		[]uuid.UUID{uuid.MustParse("52621674-47cb-498d-9df6-b1a69d3dcc4a")},
		uuid.MustParse("6ae8d984-96c0-4e25-a0a5-3ae2ae4c1aeb"),
		[]uuid.UUID{uuid.MustParse("a98d022e-073a-45c2-9a8e-f4d0a8ea52db")},
	)

	if err != nil {
		return nil, common.ErrorCannotListEntity(questionentity.EntityName, err)
	}

	return questions, nil
}
