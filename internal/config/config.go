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
	UserRole    int8  = 0
	CourierRole uint8 = 1
)

// Package statuses
/*
const (
	PACKAGE_STATUS_CREATED          = 0
	PACKAGE_STATUS_DELIVERY         = 1
	PACKAGE_STATUS_READY_TO_RECEIVE = 2
	PACKAGE_STATUS_RECEIVED         = 3
)

const (
	PACKAGE_TYPE_LETTER = 0
	PACKAGE_TYPE_BOX    = 1
	PACKAGE_TYPE_CARD   = 2
)
*/
