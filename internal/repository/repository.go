package repository

import (
	"api/internal/repository/user"
	"api/internal/repository/post"
	"api/internal/types"
	"database/sql"
)

type Repository struct {
	UserRepo types.UserRepo
	PostRepo types.PostRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserRepo: user.NewUserDB(db),
		PostRepo: post.NewPostDB(db),
	}
}
