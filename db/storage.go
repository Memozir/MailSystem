package db

import (
	"context"
)

type Storage interface {
	CreateTables(context.Context) error
}
