package post

import (
	"api/internal/types"
)

type PostService struct {
	repo types.PostRepo
}

func NewPostService(repo types.PostRepo) *PostService {
	return &PostService{repo: repo}
}

func (p *PostService) UpdateProfile(postData *types.EditProfile) error {
	post := &types.Profile{
		AuthorId:   postData.AuthorId,
		AuthorName: postData.Username,
		Mobile:     postData.Mobile,
		BirthDate:  postData.BirthDate,
	}

	err := p.repo.UpdateProfileDB(post)
	if err != nil {
		return err
	}
	return nil
}
