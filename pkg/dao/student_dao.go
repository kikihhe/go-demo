package dao

import (
	"go-demo/pkg/model"
	"go-demo/pkg/model/domain"
	"gorm.io/gorm"
)

type studentDao struct {
	database *gorm.DB
}

func NewStudentDao() domain.StudentDao {
	return &studentDao{
		database: model.GetDB(),
	}
}

func (studentMapper *studentDao) Create(student *domain.Student) (int, error) {
	result := model.GetDB().Create(student)
	return int(result.RowsAffected), result.Error
}

func (studentMapper *studentDao) GetById(id int64) (*domain.Student, error) {
	var student domain.Student
	db := model.GetDB()
	result := db.Model(&domain.Student{}).First(&student, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &student, nil
}
