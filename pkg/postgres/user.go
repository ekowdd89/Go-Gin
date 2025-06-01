package postgres

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserRepository interface {
	FindUsers() ([]User, error)
}

var _ UserRepository = &Postgres{}

func (p Postgres) FindUsers() (u []User, err error) {
	rows, err := p.db.Query("SELECT id, name, password, email FROM sys_ms_users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email); err != nil {
			return nil, err
		}
		u = append(u, user)
	}

	println(u)
	return
}
