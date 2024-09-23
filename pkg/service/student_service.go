package service

import (
	"github.com/gin-gonic/gin"
	"go-demo/pkg/dao"
	"go-demo/pkg/model/domain"
)

type StudentService struct {
	studentDao domain.StudentDao
}

func NewStudentService() *StudentService {
	return &StudentService{
		studentDao: dao.NewStudentDao(),
	}
}

func (s *StudentService) GetStudentById(c *gin.Context) (*domain.Student, error) {
	id := c.GetInt64("id")
	student, err := dao.NewStudentDao().GetById(id)
	if err != nil {
		return nil, err
	}
	return student, nil
}
