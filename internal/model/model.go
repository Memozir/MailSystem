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

type Employee struct {
	EmployeeId   uint64 `db:"id"`
	UserId       uint64 `db:"user"`
	RoleCode     int8   `db:"role"`
	DepartmentId uint64 `db:"department"`
}

/*
	'id': 2,
	'status': 3,
	'weight': 2100,
	'sender': '8921312',
	'receiver': '895975',
	'courier': 'qwe323rty',
	'date_of_creation': '2002-11-01',
	'date_of_receipt': '2002-11-04',
	'num_department': 3,
	'address': 'Baturina1'
*/

type Package struct {
	Id             uint64 `db:"id" json:"id"`
	Status         int    `db:"status" json:"status"`
	Weight         int    `db:"weight" json:"weight"`
	Sender         string `db:"sender" json:"sender"`
	Receiver       string `db:"receiver" json:"receiver"`
	Courier        uint64 `db:"courier" json:"courier"`
	DateOfCreation string `db:"create_date" json:"date_of_creation"`
	DateOfReceipt  string `db:"deliver_date" json:"date_of_receipt"`
	NumDepartment  uint64 `db:"department_receiver" json:"num_department"`
	Address        string `db:"sender_address" json:"address"`
	Type           int    `db:"type" json:"type"`
}
