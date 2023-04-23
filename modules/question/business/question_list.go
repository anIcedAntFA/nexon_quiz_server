package questionbusiness

import (
	"context"
	"log"
	"math"
	"nexon_quiz/common"
	questionentity "nexon_quiz/modules/question/entity"
	userentity "nexon_quiz/modules/user/entity"
)

type QuestionListStorage interface {
	QuestionList(
		ctx context.Context,
		filter *questionentity.Filter,
		queryParams *common.QueryParams,
		moreKeys ...string,
	) ([]questionentity.Question, error)
}

type questionListBusiness struct {
	storage   QuestionListStorage
	requester common.Requester
}

func NewQuestionListBusiness(
	storage QuestionListStorage,
	requester common.Requester,
) *questionListBusiness {
	return &questionListBusiness{
		storage:   storage,
		requester: requester,
	}
}

func (biz *questionListBusiness) QuestionList(
	ctx context.Context,
	filter *questionentity.Filter,
	queryParams *common.QueryParams,
) ([]questionentity.Question, *questionentity.QuestionPagingResult, error) {
	ctxStore := context.WithValue(ctx, common.CurrentUser, biz.requester)

	questions, err := biz.storage.QuestionList(ctxStore, filter, queryParams)

	if biz.requester.GetRole() == userentity.RoleUser.String() {
		for _, question := range questions {
			for _, answer := range *question.Answers {
				answer.Correct = 0
				log.Println("answer", answer)
			}
		}
	}

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
