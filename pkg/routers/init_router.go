package routers

import (
	"github.com/gin-gonic/gin"
	"go-demo/configs"
	"go-demo/pkg/handler"
)

func InitRouter() {
	r := gin.Default()

	group := r.Group("/student")
	studentRouter(group, handler.NewStudentHandler())

	addr := configs.GetAddress()
	r.Run(addr)
}
