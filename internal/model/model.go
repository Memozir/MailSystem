package model

import (
	"fmt"
)

type User struct {
	// id SERIAL PRIMARY KEY,
	// phone VARCHAR(11) NOT NULL,
	// pass VARCHAR(20) NOT NULL,
	// first_name VARCHAR(50) NOT NULL,
	// second_name VARCHAR(50) NOT NULL,
	// birth_date DATE

	Id         uint64 `db:"id"`
	Phone      string `db:"phone"`
	Pass       string `db:"pass"`
	FirstName  string `db:"first_name"`
	SecondName string `db:"second_name"`
	BirthDate  []byte `db:"birth_date"`
}

func (user *User) String() string {
	return fmt.Sprintf(
		"id=%d phone=%s pass=%s firstName=%s secondName=%s birthDate=%s",
		user.Id,
		user.Phone,
		user.Pass,
		user.FirstName,
		user.SecondName,
		user.BirthDate)
}
