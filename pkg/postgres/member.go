package postgres

import "database/sql"

type Member struct {
	Id        int            `json:"id"`
	FirstName sql.NullString `json:"first_name"`
	LastName  sql.NullString `json:"last_name"`
	Email     sql.NullString `json:"email"`
}

type MemberRepository interface {
	FindMembers() ([]Member, error)
}

var _ MemberRepository = &Postgres{}

func (p Postgres) FindMembers() (u []Member, err error) {

	rows, err := p.db.Query("SELECT id, first_name, last_name, email FROM sys_ms_member")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var m Member
		err := rows.Scan(&m.Id, &m.FirstName, &m.LastName, &m.Email)
		if err != nil {
			return nil, err
		}
		u = append(u, m)
	}
	return
}
