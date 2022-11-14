package usecase

import (
	"base-go/domain"
	"context"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{userRepo}
}

func (usecase userUsecase) Fetch(ctx context.Context) (result []domain.User, err error) {
	result, err = usecase.userRepo.Fetch(ctx)

	if err != nil {
		return
	}

	return
}

func (usecase userUsecase) Create(ctx context.Context, user domain.User) (result domain.User, err error) {
	result, err = usecase.userRepo.Create(ctx, user)
	if err != nil {
		return
	}

	return
}

func (usecase userUsecase) Update(ctx context.Context, user domain.User) (result domain.User, err error) {
	userData, err := usecase.userRepo.GetByID(ctx, user.ID)
	if err != nil {
		return
	}

	userData.FirstName = user.FirstName
	userData.LastName = user.LastName
	userData.Email = user.Email

	result, err = usecase.userRepo.Update(ctx, userData)
	if err != nil {
		return
	}

	return
}

func (usecase userUsecase) GetByID(ctx context.Context, id uint) (result domain.User, err error) {
	result, err = usecase.userRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}

func (usecase userUsecase) Delete(ctx context.Context, id uint) (err error) {
	if err = usecase.userRepo.Delete(ctx, id); err != nil {
		return
	}

	return
}
