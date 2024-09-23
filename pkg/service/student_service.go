package service

import (
	"github.com/gin-gonic/gin"
	"go-demo/pkg/dao"
	"go-demo/pkg/model/domain"
	"strconv"
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

	idString := c.Query("id")
	id, _ := strconv.Atoi(idString)
	student, err := dao.NewStudentDao().GetById(int64(id))
	if err != nil {
		return nil, err
	}
	return student, nil
}
