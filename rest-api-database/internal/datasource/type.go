package datasource

type User struct {
	ID        int      `json:"id" db:"id"`
	Name      string   `json:"name" db:"name"`
	Email     string   `json:"email" db:"email"`
	Company   string   `json:"company" db:"company"`
	Interests []string `json:"interests,omitempty" db:"interests"`
}

// Instructor type represent a instructor for a course
type Instructor struct {
	ID        int      `json:"id" db:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Company   string   `json:"company"`
	Expertise []string `json:"expertise,omitempty"`
}

// Course is course being taught
type Course struct {
	ID           int      `json:"id" db:"id"`
	InstructorID int      `json:"instructor_id"`
	Name         string   `json:"name"`
	Topics       []string `json:"topics,omitempty"`
	Attendees    []int    `json:"attendees,omitempty"`
}
