package handler

import (
	"net/http"
	"new-api/common"
	"new-api/middleware"
	"new-api/model"
	"new-api/router"
	"github.com/gin-gonic/gin"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	common.SetupLogger()
	common.InitDatabase()
	model.InitLogDB()
	
	engine := gin.Default()
	router.SetRouter(engine)
	middleware.SetRouter(engine)
	
	engine.ServeHTTP(w, r)
}
