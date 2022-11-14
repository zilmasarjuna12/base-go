package domain

import "context"

// User
type User struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	FirstName string `gorm:"column:firstname" json:"firstname"`
	LastName  string `gorm:"column:lastname" json:"lastname"`
}

func (User) TableName() string {
	return "user"
}

// UserUsecase represent the usecase
type UserUsecase interface {
	Fetch(ctx context.Context) ([]User, error)
	Create(ctx context.Context, user User) (User, error)
	Update(ctx context.Context, user User) (User, error)
	GetByID(ctx context.Context, id uint) (User, error)
	Delete(ctx context.Context, id uint) error
}

// UserRepository represent the repository
type UserRepository interface {
	Fetch(ctx context.Context) ([]User, error)
	Create(ctx context.Context, user User) (User, error)
	Update(ctx context.Context, user User) (User, error)
	GetByID(ctx context.Context, id uint) (User, error)
	Delete(ctx context.Context, id uint) error
}
