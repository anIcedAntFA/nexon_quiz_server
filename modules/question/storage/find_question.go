package questionstorage

import (
	"context"
	"nexon_quiz/common"
	questionentity "nexon_quiz/modules/question/entity"

	"gorm.io/gorm"
)

func (s *questionMySQLStorage) FindQuestion(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*questionentity.Question, error) {
	db := s.db

	for _, v := range moreKeys {
		db = db.Preload(v)
	}

	var result questionentity.Question

	if err := db.Where(condition).First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrorRecordNotFound
		}

		return nil, common.ErrorDB(err)
	}

	return &result, nil
}
