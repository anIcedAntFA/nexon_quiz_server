package middleware

import (
	"errors"
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"

	"github.com/gin-gonic/gin"
)

func RequiredRole(appCtx appctx.AppContext, allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		hasFound := false

		for _, role := range allowedRoles {
			if requester.GetRole() == role {
				hasFound = true
				break
			}
		}

		if !hasFound {
			panic(common.ErrorNoPermission(errors.New("invalid role")))
		}

		ctx.Next()
	}
}
