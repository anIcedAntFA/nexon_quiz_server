package difficultystorage

import (
	"context"
	"nexon_quiz/common"
	difficultyentity "nexon_quiz/modules/difficulty/entity"

	"gorm.io/gorm"
)

func (ds *difficultyMySQLStorage) FindDifficultyByCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*difficultyentity.Difficulty, error) {
	var data difficultyentity.Difficulty

	if err := ds.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, difficultyentity.ErrorDifficultyNotFound
		}

		return nil, common.ErrorDB(err)
	}

	return &data, nil
}
