package questionstorage

import (
	"context"
	"nexon_quiz/common"
	questionentity "nexon_quiz/modules/question/entity"

	"github.com/google/uuid"
)

func (s *questionMySQLStorage) FindPlayingQuestionList(
	ctx context.Context,
	quantity int,
	typeSettingIds []uuid.UUID,
	difficultySettingId uuid.UUID,
	categorySettingIds []uuid.UUID,
	moreKeys ...string,
) ([]questionentity.Question, error) {
	db := s.db.
		Table(questionentity.Question{}.TableName()).
		Where("deleted_at IS NULL")

	for _, v := range moreKeys {
		db = db.Preload(v, "deleted_at IS NULL")
	}

	db = db.Where("type in (?)", typeSettingIds)

	db = db.Where("difficulty = ?", difficultySettingId)

	db = db.Where("category in (?)", categorySettingIds)

	var data []questionentity.Question

	if err := db.
		Limit(quantity).
		Find(&data).
		Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return data, nil
}
