package typebusiness

import (
	"context"
	"math"
	"nexon_quiz/common"
	typeentity "nexon_quiz/modules/type/entity"
)

type FindTypeListStorage interface {
	FindTypeList(
		ctx context.Context,
		filter *typeentity.TypeFilter,
		queryParams *common.QueryParams,
		moreKeys ...string,
	) ([]typeentity.Type, error)
}

type findTypeListBusiness struct {
	storage FindTypeListStorage
}

func NewFindTypeListBusiness(
	storage FindTypeListStorage,
) *findTypeListBusiness {
	return &findTypeListBusiness{
		storage: storage,
	}
}

func (biz *findTypeListBusiness) GetTypeList(
	ctx context.Context,
	filter *typeentity.TypeFilter,
	queryParams *common.QueryParams,
) ([]typeentity.Type, *typeentity.TypePagingResult, error) {

	data, err := biz.storage.FindTypeList(ctx, filter, queryParams)

	var pagingResult typeentity.TypePagingResult

	if len(data) > 0 {
		pagingResult = typeentity.TypePagingResult{
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
			"ErrorCannotGetTypeList",
		)
	}

	return data, &pagingResult, nil
}
