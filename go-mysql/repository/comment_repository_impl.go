package repository

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/mhaatha/go-learn/go-mysql/entity"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	sqlCommand := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := repository.DB.ExecContext(ctx, sqlCommand, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}

	comment.Id = int32(id)
	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	comment := entity.Comment{}

	sqlCommand := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, sqlCommand, id)
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		// id exists
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		// id not exists
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " Not found")
	}
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	sqlCommand := "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, sqlCommand)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []entity.Comment

	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}

	return comments, nil
}
