package repository

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	go_mysql "github.com/mhaatha/go-learn/go-mysql"
	"github.com/mhaatha/go-learn/go-mysql/entity"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(go_mysql.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "Test Repository",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(go_mysql.GetConnection())

	ctx := context.Background()

	comment, err := commentRepository.FindById(ctx, 41)
	if err != nil {
		panic(err)
	}

	fmt.Print(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(go_mysql.GetConnection())

	ctx := context.Background()

	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Print(comment)
	}
}
