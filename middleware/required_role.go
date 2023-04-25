package middleware

import (
	"errors"
	"log"
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"

	"github.com/gin-gonic/gin"
)

func RequiredRole(appCtx appctx.AppContext, allowedRoles ...int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("requester1")

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)
		log.Println("requester", requester)

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
