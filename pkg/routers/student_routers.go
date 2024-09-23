package routers

import (
	"github.com/gin-gonic/gin"
	"go-demo/pkg/handler"
)

func studentRouter(group *gin.RouterGroup, h handler.StudentHandler) {
	group.GET("/getById", h.GetById)
	group.GET("/ping", h.Ping)
}
