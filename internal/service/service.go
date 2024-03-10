package service

import (
	"api/internal/repository"
	"api/internal/service/post"
	"api/internal/service/user"
	"api/internal/types"
)

type Service struct {
	UserService types.UserService
	PostService types.PostService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserService: user.NewUserService(repo.UserRepo),
		PostService: post.NewPostService(repo.PostRepo),
	}
}
