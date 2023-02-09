package repository

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	db_conn "belajar-golang-database/db_conn"
	"belajar-golang-database/entity"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(db_conn.InitDB())
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
	commentRepository := NewCommentRepository(db_conn.InitDB())

	comment, err := commentRepository.FindById(context.Background(), 64)

	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(db_conn.InitDB())

	comments, err := commentRepository.FindAll(context.Background())

	if err != nil {
		panic(err)
	}

	for _,comment := range comments {
		fmt.Println(comment)
	}
}
