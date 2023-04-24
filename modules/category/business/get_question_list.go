package categorybusiness

import (
	"context"
	"math"
	"nexon_quiz/common"
	categoryentity "nexon_quiz/modules/category/entity"
)

type FindCategoryListStorage interface {
	FindCategoryList(
		ctx context.Context,
		filter *categoryentity.CategoryFilter,
		queryParams *common.QueryParams,
		moreKeys ...string,
	) ([]categoryentity.Category, error)
}

type findCategoryListBusiness struct {
	storage FindCategoryListStorage
}

func NewFindCategoryListBusiness(
	storage FindCategoryListStorage,
) *findCategoryListBusiness {
	return &findCategoryListBusiness{
		storage: storage,
	}
}

func (biz *findCategoryListBusiness) GetQuestionList(
	ctx context.Context,
	filter *categoryentity.CategoryFilter,
	queryParams *common.QueryParams,
) ([]categoryentity.Category, *categoryentity.CategoryPagingResult, error) {

	data, err := biz.storage.FindCategoryList(ctx, filter, queryParams)

	var pagingResult categoryentity.CategoryPagingResult

	if len(data) > 0 {
		pagingResult = categoryentity.CategoryPagingResult{
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
			"ErrorCannotGetCategoryList",
		)
	}

	return data, &pagingResult, nil
}
