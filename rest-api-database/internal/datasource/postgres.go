package datasource

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v4/pgxpool"
)

type postgres struct {
	pool *pgxpool.Pool
}

func NewPostgres(pool *pgxpool.Pool) *postgres {
	p := postgres{pool}
	return &p
}

func (p *postgres) GetAllCourses() ([]Course, error) {
	var courses []Course
	rows, err := p.pool.Query(context.Background(), `select c.id, c.name, c."instructorId" from "Courses" as c`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var course Course
		err = rows.Scan(&course.ID, &course.Name, &course.InstructorID)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

func (p *postgres) GetAllInstructors() ([]Instructor, error) {
	var instructors []Instructor
	rows, err := p.pool.Query(context.Background(), `select i.id, i.name, i.email, i.company from "Instructors" as i`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var instructor Instructor
		err = rows.Scan(&instructor.ID, &instructor.Name, &instructor.Email, &instructor.Company)
		if err != nil {
			return nil, err
		}
		instructors = append(instructors, instructor)
	}
	return instructors, nil
}

func (p *postgres) GetAllUsers() ([]User, error) {
	var users []User
	rows, err := p.pool.Query(context.Background(), `select u.id, u.name, u.email, u.company from "Users" as u`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Company)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (p *postgres) GetCoursesByID(id int) (*Course, error) {
	var course Course
	err := p.pool.QueryRow(context.Background(),
		`select c.id, c."instructorId", c.name, agg.attendees, agg.topics
		from "Courses" as c
			 left join (select caagg."courseId", attendees, topics
						from (select ca."courseId", array_agg(ca."attendeeId") as attendees
							  from "CourseAttendee" as ca
							  group by ca."courseId") as caagg
								 left join (select ct."courseId", array_agg(ct.topic) as topics
											from "CourseTopic" as ct
											group by ct."courseId") as ctagg
										   on caagg."courseId" = ctagg."courseId")
		as agg on c.id = agg."courseId"
		where c.id = $1;`, id).Scan(&course.ID, &course.InstructorID, &course.Name, &course.Attendees, &course.Topics)
	if err != nil {
		return nil, err
	}
	return &course, nil
}

func (p *postgres) GetInstructorByID(id int) (*Instructor, error) {
	var instructor Instructor
	err := p.pool.QueryRow(context.Background(),
		`select i.id, i.name, i.email, i.company, agg.expertise
		from "Instructors" as i
		left join (select ie."instructorId", array_agg(ie.topic) as expertise
			   from "InstructorExpertise" as ie
			   group by ie."instructorId") as agg on i.id = agg."instructorId"
		where i.id = 1;`, id).Scan(&instructor.ID, &instructor.Name, &instructor.Email, &instructor.Company, &instructor.Expertise)
	if err != nil {
		return nil, err
	}
	return &instructor, nil
}

func (p *postgres) GetUserByID(id int) (*User, error) {
	var user User
	err := p.pool.QueryRow(context.Background(),
		`select u.id, u.name, u.email, u.company, agg.interests
		from "Users" as u
		left join (select ui."userId", array_agg(ui.topic) as interests
			   from "UserInterest" as ui
			   group by ui."userId") as agg on u.id = agg."userId"
		where u.id = $1;`, id).Scan(&user.ID, &user.Name, &user.Email, &user.Company, &user.Interests)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (p *postgres) GetCoursesForInstructor(instructorID int) ([]Course, error) {
	var courses []Course
	rows, err := p.pool.Query(context.Background(),
		`select c.id, c.name, c."instructorId", agg.topics
		from "Courses" as c
				 left join (select ct."courseId", array_agg(ct.topic) as topics
							from "CourseTopic" as ct
							group by ct."courseId") as agg on c.id = agg."courseId"
		where c."instructorId" = $1;`, instructorID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var course Course
		err = rows.Scan(&course.ID, &course.Name, &course.InstructorID, &course.Topics)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

func (p *postgres) GetCoursesForUser(userID int) ([]Course, error) {
	var courses []Course

	rows, err := p.pool.Query(context.Background(),
		`select c.id, c.name, c."instructorId", agg.topics
		from (select caagg."courseId", array_agg(ct.topic) as topics
		  	from (select ca."attendeeId", ca."courseId" from "CourseAttendee" as ca where ca."attendeeId" = $1) as caagg
				left join "CourseTopic" as ct on caagg."courseId" = ct."courseId"
		  	group by caagg."courseId") as agg
			left join "Courses" as c on c.id = agg."courseId";`, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var course Course

		err = rows.Scan(&course.ID, &course.Name, &course.InstructorID, &course.Topics)
		if err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}

	return courses, nil
}

func (p *postgres) CreateNewUser(user *User) (int, error) {
	tx, err := p.pool.Begin(context.Background())
	if err != nil {
		return -1, err
	}
	defer tx.Rollback(context.Background())
	var id int
	err = tx.QueryRow(context.Background(), `INSERT into "Users" (name, email, company) VALUES ($1, $2, $3) returning id`, user.Name, user.Email, user.Company).Scan(&id)

	if err != nil {
		return -1, err
	}

	if user.Interests != nil {
		var ds *goqu.InsertDataset
		rows := make([]interface{}, len(user.Interests))
		for i, interest := range user.Interests {
			rows[i] = goqu.Record{"userId": id, "topic": interest}
		}

		ds = goqu.Insert("UserInterest").Rows(rows)

		sql, _, err := ds.ToSQL()
		if err != nil {
			return -1, err
		}
		_, err = tx.Exec(context.Background(), sql)
		if err != nil {
			return -1, err
		}
	}

	if err := tx.Commit(context.Background()); err != nil {
		return -1, err
	}
	return id, nil
}

func (p *postgres) AddUserInterest(id int, interests []string) (int, error) {
	tx, err := p.pool.Begin(context.Background())
	if err != nil {
		return -1, err
	}

	defer tx.Rollback(context.Background())

	var ds *goqu.InsertDataset
	rows := make([]interface{}, len(interests))
	for i, interest := range interests {
		rows[i] = goqu.Record{"userId": id, "topic": interest}
	}
	ds = goqu.Insert("UserInterest").Rows(rows)

	sql, _, err := ds.ToSQL()
	if err != nil {
		return -1, err
	}
	_, err = tx.Exec(context.Background(), sql)
	if err != nil {
		return -1, err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return -1, err
	}
	return len(interests), nil
}
