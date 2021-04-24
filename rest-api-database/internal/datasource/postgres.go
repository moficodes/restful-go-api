package datasource

import (
	"context"
	"log"

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
	return nil, nil
}

func (p *postgres) GetAllInstructors() ([]Instructor, error) {
	return nil, nil
}

func (p *postgres) GetAllUsers() ([]User, error) {
	return nil, nil
}

func (p *postgres) GetCoursesByID(id int) (*Course, error) {
	return nil, nil
}

func (p *postgres) GetInstructorByID(id int) (*Instructor, error) {
	return nil, nil
}

func (p *postgres) GetUserByID(id int) (*User, error) {
	var user User
	err := p.pool.QueryRow(context.Background(), `select u.id, u.name, u.email, u.company, array_agg(ui.topic) as interests from "Users" as u left join "UserInterest" as ui on u.id = ui."userId" where u.id = $1 group by u.id;`, id).Scan(&user.ID, &user.Name, &user.Email, &user.Company, &user.Interests)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, nil
}
