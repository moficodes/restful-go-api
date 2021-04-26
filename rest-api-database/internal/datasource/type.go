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
	Name      string   `json:"name" db:"name"`
	Email     string   `json:"email" db:"email"`
	Company   string   `json:"company" db:"company"`
	Expertise []string `json:"expertise,omitempty" db:"expertise"`
}

// Course is course being taught
type Course struct {
	ID           int      `json:"id" db:"id"`
	InstructorID int      `json:"instructor_id" db:"instructorId"`
	Name         string   `json:"name" db:"name"`
	Topics       []string `json:"topics,omitempty" db:"topics"`
	Attendees    []int    `json:"attendees,omitempty" db:"attendees"`
}
