package handler

type User struct {
	ID        int      `json:"id" example:"1"`
	Name      string   `json:"name" example:"John Doe"`
	Email     string   `json:"email" example:"johndoe@gmail.com"`
	Company   string   `json:"company" example:"Acme Inc."`
	Interests []string `json:"interests" example:"golang,python"`
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
	Data string `json:"data" example:"John Doe"`
}

type HTTPError struct {
	Code     int         `json:"-"`
	Message  interface{} `json:"message"`
	Internal error       `json:"-"` // Stores the error returned by an external dependency
}
