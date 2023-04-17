package middleware

import (
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"

	"github.com/gin-gonic/gin"
)

func Recover(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.Header("Content-Type", "application/json")

				if appErr, found := err.(*common.AppError); found {
					ctx.AbortWithStatusJSON(appErr.StatusCode, appErr)
					// panic(err)
					return
				}

				appErr := common.ErrorInternal(err.(error))
				ctx.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
				// return
			}
		}()

		ctx.Next()
	}
}
