package typestorage

import (
	"context"
	"nexon_quiz/common"
	typeentity "nexon_quiz/modules/type/entity"
	"time"

	"github.com/google/uuid"
)

func (ts *typeMySQLStorage) DeleteTypeById(
	ctx context.Context,
	id uuid.UUID,
) error {
	if err := ts.db.Table(typeentity.Type{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"deleted_at": time.Now(),
		}).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
