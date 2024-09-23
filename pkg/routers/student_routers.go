package routers

import (
	"github.com/gin-gonic/gin"
	"go-demo/pkg/handler"
)

func studentRouter(group *gin.RouterGroup, h handler.StudentHandler) {
	group.GET("/user/getById", h.GetById)
}
