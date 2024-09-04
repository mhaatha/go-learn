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

func TestInsertData(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlCommand := "INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES('01', 'Hafidz', 'hafidz123@gmail.com', 100000, 5.0, '1999-09-19', true)"
	_, err := db.ExecContext(ctx, sqlCommand)
	if err != nil {
		panic(err)
	}

	sqlCommand = "INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES('02', 'Athaya', 'athaya123@gmail.com', 200000, 5.0, '1999-09-19', true)"
	_, err = db.ExecContext(ctx, sqlCommand)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new data to customer table")
}
