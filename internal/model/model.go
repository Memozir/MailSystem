package model

import (
	"fmt"
)

type User struct {
	Id        uint64 `db:"id"`
	Phone     string `db:"phone"`
	Pass      string `db:"pass"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	BirthDate string `db:"birth_date"`
}

func (user *User) String() string {
	return fmt.Sprintf(
		"id=%d phone=%s pass=%s firstName=%s secondName=%s birthDate=%s",
		user.Id,
		user.Phone,
		user.Pass,
		user.FirstName,
		user.LastName,
		user.BirthDate)
}

type UserAuth struct {
	ClientId uint64 `db:"client_id"`
	RoleCode int8   `db:"role_code"`
}
