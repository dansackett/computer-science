package mvc

import "fmt"

type Student struct {
	name string
}

func NewStudent(name string) *Student {
	return &Student{name: name}
}

func (s *Student) GetName() string {
	return s.name
}

func (s *Student) SetName(name string) {
	s.name = name
}

type StudentView struct{}

func NewStudentView() *StudentView {
	return &StudentView{}
}

func (v *StudentView) Output(student *Student) string {
	return fmt.Sprintf("Student name is %s", student.name)
}

type StudentController struct {
	model *Student
	view  *StudentView
}

func NewStudentController(model *Student, view *StudentView) *StudentController {
	return &StudentController{model: model, view: view}
}

func (c *StudentController) SetStudentName(name string) {
	c.model.SetName(name)
}

func (c *StudentController) GetStudent() *Student {
	return c.model
}

func (c *StudentController) View() string {
	return c.view.Output(c.GetStudent())
}
