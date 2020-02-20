package user

import "database/sql"

type User struct {
	ID       int    `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func getUser(db *sql.DB, username string, password string) (User, error) {
	var user User
	sqlStatement := `SELECT id, username, password from todo_user WHERE username = $1 AND password = $2`
	err := db.QueryRow(
		sqlStatement, username, password,
	).Scan(
		&user.ID, &user.Username, &user.Password,
	)
	if err != nil {
		return User{
			ID:       0,
			Username: "",
			Password: "",
		}, err
	}
	return user, nil
}
