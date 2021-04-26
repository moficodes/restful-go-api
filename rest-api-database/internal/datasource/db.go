package datasource

type DB interface {
	GetAllCourses() ([]Course, error)
	GetAllInstructors() ([]Instructor, error)
	GetAllUsers() ([]User, error)
	GetCoursesByID(int) (*Course, error)
	GetInstructorByID(int) (*Instructor, error)
	GetUserByID(int) (*User, error)
	GetCoursesForInstructor(int) ([]Course, error)
	GetCoursesForUser(int) ([]Course, error)

	CreateNewUser(*User) (int, error)
	AddUserInterest(int, []string) (int, error)
}
