package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DashboardController struct{}

func (ctrl *DashboardController) ShowDashboard(c *gin.Context) {
	// ユーザーの認証情報を取得し、ダッシュボード用のデータを取得する処理など

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"title": "Dashboard",
		// ダッシュボードに表示するデータを渡す
	})
}

func NewDashboardController() *DashboardController {
	return &DashboardController{}
}
