package difficultybusiness

import (
	"context"
	"math"
	"nexon_quiz/common"
	difficultyentity "nexon_quiz/modules/difficulty/entity"
)

type FindDifficultyListStorage interface {
	FindDifficultyList(
		ctx context.Context,
		filter *difficultyentity.DifficultyFilter,
		queryParams *common.QueryParams,
		moreKeys ...string,
	) ([]difficultyentity.Difficulty, error)
}

type findDifficultyListBusiness struct {
	storage FindDifficultyListStorage
}

func NewFindDifficultyListBusiness(
	storage FindDifficultyListStorage,
) *findDifficultyListBusiness {
	return &findDifficultyListBusiness{
		storage: storage,
	}
}

func (biz *findDifficultyListBusiness) GetDifficultyList(
	ctx context.Context,
	filter *difficultyentity.DifficultyFilter,
	queryParams *common.QueryParams,
) ([]difficultyentity.Difficulty, *difficultyentity.DifficultyPagingResult, error) {

	data, err := biz.storage.FindDifficultyList(ctx, filter, queryParams)

	var pagingResult difficultyentity.DifficultyPagingResult

	if len(data) > 0 {
		pagingResult = difficultyentity.DifficultyPagingResult{
			PreviousPage: queryParams.CurrentPage - 1,
			CurrentPage:  queryParams.CurrentPage,
			NextPage:     queryParams.CurrentPage + 1,
			PageSize:     queryParams.PageSize,
			TotalItems:   int(queryParams.TotalItems),
			TotalPages:   int(math.Ceil(float64(queryParams.TotalItems) / float64(queryParams.PageSize))),
		}
	}

	if err != nil {
		return nil, nil, common.NewCustomError(
			err,
			err.Error(),
			"ErrorCannotGetDifficultyList",
		)
	}

	return data, &pagingResult, nil
}
