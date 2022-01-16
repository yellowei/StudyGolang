package aboutGoTest

var TestSenior = Senior{

}

var TestStudent = Student{
	name: "jack",
	Age:  12,
	a: &TestSenior,
}


type A interface {
	GetA() string
}

type Student struct {
	a A
	name string
	Age  int
}

func (s *Student)GetA() string  {

	return "Student"
}

type Senior struct {
	Student
}

func (s *Senior)GetA() string  {

	return "Senior"
}

