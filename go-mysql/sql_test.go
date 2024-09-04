package go_mysql

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlCommand := "INSERT INTO customer(id, name) VALUES('01', 'Hafidz Athaya')"
	_, err := db.ExecContext(ctx, sqlCommand)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new data to customer table")
}
