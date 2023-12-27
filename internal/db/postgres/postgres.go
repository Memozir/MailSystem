package db

import (
	"context"
	"fmt"
	"log"
	"mail_system/internal/model"
	"os"

	"github.com/jackc/pgx/v5"

	pgxPool "github.com/jackc/pgx/v5/pgxpool"

	"mail_system/internal/config"
)

type ResultDB struct {
	Val any
	Err error
}

type Storage interface {
	Reset()
	CreateUser(ctx context.Context,
		firstName string,
		secondName string,
		middleName string,
		login string,
		pass string,
		birth string) ResultDB
	CreateEmployee(
		ctx context.Context, userId uint8, departmentId uint64, roleCode uint8) ResultDB
	DeleteEmployee(ctx context.Context, employeeId uint64) error
	CreateRole(ctx context.Context, code uint8, name string) ResultDB
	GetAddressByName(ctx context.Context, addressName string, apartment string) (uint8, error)
	CreateClient(ctx context.Context, userId uint8, addressName string, apartment string) error
	GetRoleByName(ctx context.Context, roleName string) ResultDB
	AuthUser(ctx context.Context, login string, pass string) (ResultDB, error)
	AddPackageToClient(ctx context.Context, clientId uint64, packageId uint64) error
	GetDepartmentByReceiver(ctx context.Context, receiverId uint64) (uint64, error)
	GetEmployeeByLogin(ctx context.Context, login string) (ResultDB, error)
	GetEmployeeDepartment(ctx context.Context, login string) ResultDB
	ProducePaymentInfo(ctx context.Context, packageId uint64, packageType int, weight int) error
	AddEmployeeToPackageResponsibleList(ctx context.Context, employeeId uint64, packageId uint64) error
	GetUserIdByLogin(ctx context.Context, login string) (ResultDB, error)
	GetSenderReceiverIdByLogin(ctx context.Context, senderLogin string, receiverLogin string) (ResultDB, error)
	CreatePackage(
		ctx context.Context,
		weight int,
		packageType int,
		senderId uint64,
		receiverId uint64,
		departmentReceiver uint64,
		createDate string,
		deliverDate string) (uint64, error)
	AddPackageToStorehouse(
		ctx context.Context,
		departmentId uint64,
		packageId uint64,
		isImport bool) error
	GetEmployeePackages(
		ctx context.Context,
		employeeId uint64,
		employeeRole uint8) ([]model.Package, error)
	GetCourierDeliverPackages(ctx context.Context, departmentId uint64) ([]model.Package, error)
	GetEmployeeDepartmentByRole(ctx context.Context, departmentId uint64, role int) (uint64, error)
	BeginTran(ctx context.Context) (pgx.Tx, error)
	GetAllDepartments(ctx context.Context) ([]model.Department, error)
	GetClientDepartments(ctx context.Context, clientId uint64) ([]model.Department, error)
	GetEmployeeDepartments(ctx context.Context, employeeId uint64) ([]model.Department, error)
	GetClientByLogin(ctx context.Context, login string) (uint64, error)
	CreateAddress(ctx context.Context, departmentId uint64, addressName string) error
	DeleteAddress(ctx context.Context, addressName string) error
	CheckAdminAddress(ctx context.Context, adminId uint64, addressName string) (bool, error)
	ChangePackageStatus(ctx context.Context, packageID uint64, status uint8) error
	GetClientPackages(ctx context.Context, clientId uint64) ([]model.Package, error)
	GetDepartmentEmployees(ctx context.Context, departmentId uint64) ([]model.EmployeeInfo, error)
}

type PostgresDB struct {
	connPool *pgxPool.Pool
	cfg      *config.ConfigDb
}

func (db *PostgresDB) Reset() {
	db.connPool.Reset()
}

func NewDb(ctx context.Context) (db *PostgresDB) {
	db = new(PostgresDB)
	var err error

	db.cfg = &config.ConfigDb{
		Host:         os.Getenv("DB_HOST"),
		Port:         os.Getenv("DB_PORT"),
		Pass:         os.Getenv("DB_PASSWORD"),
		User:         os.Getenv("DB_USER"),
		DbName:       os.Getenv("DB_NAME"),
		SSLMode:      os.Getenv("DB_SSL_MODE"),
		MaxPoolConns: os.Getenv("DB_MAX_CONN_POOLS")}

	url := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s pool_max_conns=%s",
		db.cfg.User,
		db.cfg.Pass,
		db.cfg.Host,
		db.cfg.Port,
		db.cfg.DbName,
		db.cfg.SSLMode,
		db.cfg.MaxPoolConns,
	)

	db.connPool, err = pgxPool.New(ctx, url)

	if err != nil {
		// log.Fatalf("Failed connection to Postgres: %s", err)
		log.Panicf("Failed connection to Postgres: %s", err)
	}
	if err = db.connPool.Ping(ctx); err != nil {
		log.Panicf("Failed connection to Postgres: %s", err)
	}

	log.Printf("Connection to database on %s was Success", db.cfg.Host)
	return db
}

func (db *PostgresDB) BeginTran(ctx context.Context) (pgx.Tx, error) {
	res, err := db.connPool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return res, err
}
