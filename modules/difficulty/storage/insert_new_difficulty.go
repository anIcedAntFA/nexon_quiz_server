package difficultystorage

import (
	"context"
	"nexon_quiz/common"
	difficultyentity "nexon_quiz/modules/difficulty/entity"
)

func (ds *difficultyMySQLStorage) InsertNewDifficulty(
	ctx context.Context,
	newDifficulty *difficultyentity.DifficultyCreate,
) error {
	if err := ds.db.Create(newDifficulty).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
