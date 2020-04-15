package db

import "context"

type User struct {
	Name string `db:"name" json:"full_name"`
	Age  int    `db:"age" json:"age"`
}

func (s *pgStore) ListUsers(ctx context.Context) (users []User, err error) {
	s.db.Select(&users, "SELECT * FROM users ORDER BY name ASC")

	return
}
