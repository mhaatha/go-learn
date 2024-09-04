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

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlCommand := "SELECT * FROM customer"
	rows, err := db.QueryContext(ctx, sqlCommand)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)
	}
	defer rows.Close()
}
