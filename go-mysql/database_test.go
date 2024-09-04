package go_mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:dummyPassword_123@tcp(127.0.0.1:3306)/go_database")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
