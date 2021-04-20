package handler

type User struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Company   string   `json:"company"`
	Interests []string `json:"interests"`
}

// Instructor type represent a instructor for a course
type Instructor struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Company   string   `json:"company"`
	Expertise []string `json:"expertise"`
}

// Course is course being taught
type Course struct {
	ID           int      `json:"id"`
	InstructorID int      `json:"instructor_id"`
	Name         string   `json:"name"`
	Topics       []string `json:"topics"`
	Attendees    []int    `json:"attendees"`
}

type Message struct {
	Data string `json:"data"`
}
