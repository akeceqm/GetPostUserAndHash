package database

import "database/sql"

type User struct {
	Id       int            `db:"id"`
	Login    string         `db:"login"`
	Password string         `db:"password"`
	IpAddres sql.NullString `db:"ipaddres"`
}
