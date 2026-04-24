package handler

import (
	"net/http"
	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/middleware"
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/router"
	"github.com/gin-gonic/gin"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// 1. 初始化配置和数据库
	common.SetupLogger()
	common.InitDatabase()
	model.InitLogDB()
	
	// 2. 设置路由
	engine := gin.Default()
	router.SetRouter(engine)
	middleware.SetRouter(engine)
	
	// 3. 处理请求
	engine.ServeHTTP(w, r)
}
