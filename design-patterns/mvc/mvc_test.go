package mvc

import "testing"

func TestMVC(t *testing.T) {
	model := NewStudent("Dan")
	view := NewStudentView()
	controller := NewStudentController(model, view)

	if controller.View() != "Student name is Dan" {
		t.Errorf("Initial view state is incorrect")
	}

	controller.SetStudentName("Jesse")

	if controller.View() != "Student name is Jesse" {
		t.Errorf("Updated view state is incorrect")
	}
}
