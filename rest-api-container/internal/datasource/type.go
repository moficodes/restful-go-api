package datasource

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type intarr []int
type stringarr []string

// custom Scan function for scanning intarr values from database
func (i *intarr) Scan(value interface{}) error {
	source, ok := value.(string) // input example 👉🏻 {"(david,38,url,1)","(david2,2,\"url 2\",2)"}
	if !ok {
		return errors.New("incompatible type")
	}
	log.Println(source)
	source = strings.Trim(source, "{}\\\"")
	var res intarr
	ints := strings.Split(source, ",")
	for _, i := range ints {
		num, err := strconv.Atoi(i)
		if err != nil {
			return fmt.Errorf("can not convert %v to int : %v", i, err)
		}
		res = append(res, num)
	}
	*i = res
	return nil
}

func (i *stringarr) Scan(value interface{}) error {
	source, ok := value.(string) // input example 👉🏻 {"(david,38,url,1)","(david2,2,\"url 2\",2)"}
	if !ok {
		return errors.New("incompatible type")
	}
	source = strings.Trim(source, "{}\\")
	log.Println(source)
	var res stringarr
	strings := strings.Split(source, ",")
	for _, val := range strings {
		res = append(res, val)
	}
	*i = res
	return nil
}

type User struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Company   string    `json:"company" db:"company"`
	Interests stringarr `json:"interests,omitempty" db:"interests"`
}

// Instructor type represent a instructor for a course
type Instructor struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Company   string    `json:"company" db:"company"`
	Expertise stringarr `json:"expertise,omitempty" db:"expertise"`
}

// Course is course being taught
type Course struct {
	ID           int       `json:"id" db:"id"`
	InstructorID int       `json:"instructor_id" db:"instructorId"`
	Name         string    `json:"name" db:"name"`
	Topics       stringarr `json:"topics,omitempty" db:"topics"`
	Attendees    intarr    `json:"attendees,omitempty" db:"attendees"`
}
