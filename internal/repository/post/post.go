package post

import (
	"api/internal/types"
	"database/sql"
	"fmt"
)

type PostDB struct {
	db *sql.DB
}

func NewPostDB(db *sql.DB) *PostDB {
	return &PostDB{db: db}
}

func (p *PostDB) UpdateProfileDB(profile *types.Profile) error {
	_, err := p.db.Exec("UPDATE users SET mobile=$1, birthdate=$2, username=$3 WHERE id=$4", profile.Mobile, profile.BirthDate, profile.AuthorName, profile.AuthorId)
	if err != nil {
		fmt.Println("Update notify:", err)
		return err
	}

	return nil
}
