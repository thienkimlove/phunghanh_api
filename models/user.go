package models

type User struct {
	Id        int64  `db:"ID" json:"id"`
	Password  string `db:"Password" json:"password"`
	Name string `db:"Name" json:"name"`
	Email string `db:"Email" json:"email"`
}
