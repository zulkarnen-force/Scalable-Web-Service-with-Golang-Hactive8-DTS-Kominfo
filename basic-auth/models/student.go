package models

var Students []*Student

type Student struct {
	Id    string
	Name  string
	Grade int32
}

func GetStudents() []*Student {

	return Students
}

func SelectStudent(id string) *Student {
	for _, each := range Students {
		if each.Id == id {
			return each
		}
	}

	return nil
}