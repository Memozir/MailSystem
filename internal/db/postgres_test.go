package db

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetUserById(t *testing.T) {
	mainContext := context.Background()
	withTM, closeWithTM := context.WithTimeout(mainContext, 3*time.Second)
	defer closeWithTM()
	db := NewDb(withTM)
	testId := "1"
	user := db.GetUserById(testId)
	fmt.Println(user)

	if user == nil {
		t.Error("User is empty")
	}
}
