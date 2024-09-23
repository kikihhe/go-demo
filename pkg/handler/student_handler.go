package handler

import (
	"github.com/gin-gonic/gin"
	result "go-demo/pkg/model/response"
	"go-demo/pkg/service"
)

type StudentHandler interface {
	GetById(context *gin.Context)
	Ping(context *gin.Context)
}

type studentHandler struct {
	studentService *service.StudentService
}

func NewStudentHandler() StudentHandler {
	return &studentHandler{
		studentService: service.NewStudentService(),
	}
}

func (s *studentHandler) GetById(context *gin.Context) {
	var result result.Result
	student, err := s.studentService.GetStudentById(context)
	if err != nil {
		result.FailMessage(err.Error())
	} else {
		result.SuccessData(student)
	}
}

func (s *studentHandler) Ping(context *gin.Context) {
	var result result.Result
	result.Success("ping成功", nil)
}
