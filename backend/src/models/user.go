package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Create() error {
	cmd := `INSERT INTO users (name, email, password) VALUES (?, ?, ?)`
	stmt, err := Db.Prepare(cmd)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(u.Name, u.Email, u.Password)
	if err != nil {
		return err
	}
	return err
}

func Read(id int) (user User, err error) {
	cmd := `SELECT * FROM users WHERE user_id = ?`
	err = Db.QueryRow(cmd, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}
	return user, err
}
