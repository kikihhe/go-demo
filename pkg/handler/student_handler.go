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
	r := result.NewResult(context)
	student, err := s.studentService.GetStudentById(context)
	if err != nil {
		r.FailMessage(err.Error())
	} else {
		r.SuccessData(student)
	}
}

func (s *studentHandler) Ping(context *gin.Context) {
	r := result.NewResult(context)
	r.Success("ping成功", nil)
}
