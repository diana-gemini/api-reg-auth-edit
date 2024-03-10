package types

import (
	"time"
)

type User struct {
	Id           int
	Username     string
	Email        string
	PasswordHash string
	Mobile       string
	BirthDate    string
}

type Profile struct {
	AuthorId   int
	AuthorName string
	Email      string
	Mobile     string
	BirthDate  string
}

type CreateUserData struct {
	Id        int
	Email     string
	Username  string
	Password  string
	Mobile    string
	BirthDate string
	Token     *string
	Expired   *time.Time
}

type GetUserData struct {
	Email    string
	Password string
}

type EditProfile struct {
	AuthorId  int
	Email     string
	Username  string
	Mobile    string
	BirthDate string
}

type Err struct {
	StatusCode int
	StatusText string
}

type ErrText struct {
	Username string
	Email    string
	Pass1    string
	Pass2    string
}

type UserService interface {
	CreateUser(user *CreateUserData) error
	CheckUserExists(user *CreateUserData) (bool, ErrText)
	CheckLogin(user *GetUserData) (int, error)
	AddToken(userid int, cookie string) error
	RemoveToken(token string) error
	GetUserByToken(token string) (user *User, err error)
}
type UserRepo interface {
	CreateUserDB(user *User)
	GetUserEmailDB(user string) error
	CheckLoginDB(user *GetUserData) (int, error)
	AddTokenDB(userid int, cookieToken string) error
	RemoveTokenDB(token string) error
	GetUserByToken(token string) (user *User, err error)
}

type PostService interface {
	UpdateProfile(*EditProfile) error
}

type PostRepo interface {
	UpdateProfileDB(*Profile) error
}
