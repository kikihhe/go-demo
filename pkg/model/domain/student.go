package domain

type Student struct {
	ID       int64
	Name     string
	Password string
}

// 指定表名
func (student *Student) TableName() string {
	return "student"
}

type StudentDao interface {
	Create(student *Student) (int, error)
	GetById(id int64) (*Student, error)
}
