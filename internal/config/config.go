package config

type ConfigDb struct {
	Host         string
	Port         string
	Pass         string
	User         string
	DbName       string
	SSLMode      string
	MaxPoolConns string
}

const (
	UserRole int8 = 0
)
