package questionbusiness

import (
	"context"
	"math"
	"nexon_quiz/common"
	questionentity "nexon_quiz/modules/question/entity"
)

type FindQuestionListStorage interface {
	FindQuestionList(
		ctx context.Context,
		filter *questionentity.QuestionFilter,
		queryParams *common.QueryParams,
		moreKeys ...string,
	) ([]questionentity.Question, error)
}

type questionListBusiness struct {
	storage   FindQuestionListStorage
	requester common.Requester
}

func NewQuestionListBusiness(
	storage FindQuestionListStorage,
	requester common.Requester,
) *questionListBusiness {
	return &questionListBusiness{
		storage:   storage,
		requester: requester,
	}
}

func (biz *questionListBusiness) GetQuestionList(
	ctx context.Context,
	filter *questionentity.QuestionFilter,
	queryParams *common.QueryParams,
) ([]questionentity.Question, *questionentity.QuestionPagingResult, error) {
	ctxStore := context.WithValue(ctx, common.CurrentUser, biz.requester)

	questions, err := biz.storage.FindQuestionList(ctxStore, filter, queryParams)

	var pagingResult questionentity.QuestionPagingResult

	if len(questions) > 0 {
		pagingResult = questionentity.QuestionPagingResult{
			PreviousPage: queryParams.CurrentPage - 1,
			CurrentPage:  queryParams.CurrentPage,
			NextPage:     queryParams.CurrentPage + 1,
			PageSize:     queryParams.PageSize,
			TotalItems:   int(queryParams.TotalItems),
			TotalPages:   int(math.Ceil(float64(queryParams.TotalItems) / float64(queryParams.PageSize))),
		}
	}

	if err != nil {
		return nil, nil, common.ErrorCannotListEntity(questionentity.EntityName, err)
	}

	return questions, &pagingResult, nil
}
