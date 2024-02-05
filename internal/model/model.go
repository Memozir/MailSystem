package model

import (
	"fmt"
)

type User struct {
	Id         uint64 `db:"id"`
	Phone      string `db:"phone"`
	Pass       string `db:"pass"`
	FirstName  string `db:"first_name"`
	LastName   string `db:"last_name"`
	MiddleName string `db:"middle_name"`
	BirthDate  string `db:"birth_date"`
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

type EmployeeInfo struct {
	EmployeeId uint64 `db:"id" json:"employee_id"`
	RoleName   string `db:"role_name" json:"role_name"`
	Login      string `db:"login" json:"login"`
	Pass       string `db:"pass" json:"pass"`
	FirstName  string `db:"first_name" json:"first_name"`
	LastName   string `db:"last_name" json:"last_name"`
	MiddleName string `db:"middle_name" json:"middle_name"`
}

func (e *EmployeeInfo) String() string {
	return fmt.Sprintf(
		"id=%d role=%s pass=%s firstName=%s lastName=%s middleName=%s",
		e.EmployeeId,
		e.RoleName,
		e.Pass,
		e.FirstName,
		e.LastName,
		e.MiddleName,
	)
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
	Courier        string `db:"courier" json:"courier"`
	DateOfCreation string `db:"create_date" json:"date_of_creation"`
	DateOfReceipt  string `db:"deliver_date" json:"date_of_receipt"`
	NumDepartment  uint64 `db:"department_receiver" json:"num_department"`
	Address        string `db:"sender_address" json:"address"`
	Type           int    `db:"type" json:"type"`
}

type Department struct {
	Index   uint64 `db:"index" json:"index"`
	Address string `db:"name" json:"address"`
}
