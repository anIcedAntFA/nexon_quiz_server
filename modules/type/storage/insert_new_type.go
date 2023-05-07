package typestorage

import (
	"context"
	typeentity "nexon_quiz/modules/type/entity"
)

func (ts *typeMySQLStorage) InsertNewType(
	ctx context.Context,
	newType *typeentity.TypeCreate,
) (int64, error) {
	result := ts.db.Where(typeentity.TypeCreate{
		Content: newType.Content,
	}).FirstOrCreate(newType)

	// if err != nil {
	// 	return common.ErrorDB(err)
	// }

	return result.RowsAffected, nil
}
