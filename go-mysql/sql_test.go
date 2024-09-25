package go_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
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

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlCommand := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, sqlCommand)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name, email string
		var balance int32
		var rating float64
		var birthDate, createdAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("==============")
		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)
		fmt.Println("Email: ", email)
		fmt.Println("Balance: ", balance)
		fmt.Println("Rating: ", rating)
		fmt.Println("BirthDate: ", birthDate)
		fmt.Println("Married: ", married)
		fmt.Println("CreatedAt: ", createdAt)
	}
	defer rows.Close()
}

func TestQuerySqlNullable(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlCommand := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, sqlCommand)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("==============")
		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)
		if email.Valid {
			fmt.Println("Email: ", email.String)
		}
		fmt.Println("Balance: ", balance)
		fmt.Println("Rating: ", rating)
		if birthDate.Valid {
			fmt.Println("BirthDate: ", birthDate.Time)
		}
		fmt.Println("Married: ", married)
		fmt.Println("CreatedAt: ", createdAt)
	}
	defer rows.Close()
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	sqlCommand := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "'  LIMIT 1"
	rows, err := db.QueryContext(ctx, sqlCommand)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string

		err = rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses login dengan username:", username)
	} else {
		fmt.Print("Gagal login")
	}

	defer rows.Close()
}

func TestSqlInjectionSafet(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	sqlCommand := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, sqlCommand, username, password)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string

		err = rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses login dengan username:", username)
	} else {
		fmt.Print("Gagal login")
	}

	defer rows.Close()
}

func TestExecSqlSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin123"
	password := "admin123"

	sqlCommand := "INSERT INTO user(username, password) VALUES(?, ?)"
	_, err := db.ExecContext(ctx, sqlCommand, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new data to customer table")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "athaya123@gmail.com"
	comment := "Test komen"

	sqlCommand := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, sqlCommand, email, comment)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment with id", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	sqlCommand := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	statement, err := db.PrepareContext(ctx, sqlCommand)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "hafidz" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment id:", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	sqlCommand := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	// Do transaction
	for i := 0; i < 10; i++ {
		email := "hafidz" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, sqlCommand, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment id:", id)
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}
