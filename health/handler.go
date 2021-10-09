package health

import (
	"amit.vickey/ZenMoney/common"
	"amit.vickey/ZenMoney/repo"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func HealthCheck(ctx *gin.Context) {
	err := repo.CheckHealth(ctx)
	if err != nil {
		common.ErrorResponse(ctx, errors.Wrap(err, "error-checking-database-status"))
	} else {
		ctx.JSON(200, gin.H{
			"message": "OK!!!",
		})
	}
}
