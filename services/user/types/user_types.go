package user_types

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID `db:"id" json:"id"`
	Username      string    `db:"username" json:"username"`
	Email         string    `db:"email" json:"email"`
	Mobile_phone  string    `db:"mobile_phone" json:"mobile_phone"`
	Password_hash string    `db:"password_hash" json:"password_hash"`
	Host_Status   string    `db:"host_status" json:"host_status"`
	Interests     []string  `db:"interests" json:"interests"`
	States        []string  `db:"states" json:"states"`
	Created_at    time.Time `db:"created_at" json:"created_at"`
	Updated_at    time.Time `db:"updated_at" json:"updated_at"`
}

type UserResponsePayload struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Host_Status  string    `db:"host_status" json:"host_status"`
	Interests    []string  `db:"interests" json:"interests"`
	Mobile_phone string    `json:"mobile_phone"`
	Created_at   string    `json:"created_at"`
	Updated_at   string    `json:"updated_at"`
}

type UserLoginPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ResponseWithToken struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Token   string `json:"token"`
}

type UserApi interface {
	GetUserById(id int) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByUsername(username string) (*User, error)
	CreateUser(user *User) (*User, error)
	UpdateUser(user *User) (*User, error)
	// make sure this deletion removes everything related to the user
	DeleteUser(id int) error
}
