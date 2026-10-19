package domain

import "context"

type User struct {
	ID           int64
	Email        string
	Username     string
	PasswordHash string
	Bio          *string
	Image        *string
}

type UserRepo interface {
	Create(ctx context.Context, u *User) error
	ByEmail(ctx context.Context, email string) (*User, error)
	ByUsername(ctx context.Context, username string) (*User, error)
	ByID(ctx context.Context, id int64) (*User, error)
	Update(ctx context.Context, u *User) error
}
