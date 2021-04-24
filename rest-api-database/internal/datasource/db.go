package datasource

type DB interface {
	GetAllCourses() ([]Course, error)
	GetAllInstructors() ([]Instructor, error)
	GetAllUsers() ([]User, error)
	GetCoursesByID(int) (*Course, error)
	GetInstructorByID(int) (*Instructor, error)
	GetUserByID(int) (*User, error)
}
