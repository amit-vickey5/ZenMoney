package main

import (
	"log"

	"amit.vickey/ZenMoney/common"
	"amit.vickey/ZenMoney/dashboard"
	"amit.vickey/ZenMoney/health"
	"amit.vickey/ZenMoney/middleware"
	"amit.vickey/ZenMoney/repo"
	"amit.vickey/ZenMoney/setu/consent"
	"amit.vickey/ZenMoney/setu/redirect"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger(), middleware.AuthorizationHook())

	router.GET("/health", health.HealthCheck)
	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200,  gin.H{
			"consent_handle": "abcd",
		})
	})
	router.GET("/dashboard", dashboard.DashboardData)
	router.POST("/Consent/Create", consent.Create)
	router.GET("/Consent/Redirect", redirect.Redirect)
	router.POST("/Consent/Notification", consent.SetuWebhook)

	if err := repo.InitRepoClient(); err != nil {
		log.Fatal(err)
	}

	common.InitLogger()

	router.Run(":8080")
}
