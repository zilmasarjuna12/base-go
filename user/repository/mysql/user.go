package repository

import (
	"base-go/domain"
	"context"

	"gorm.io/gorm"
)

type mysqlUserRepo struct {
	DB *gorm.DB
}

func NewMysqlUserRepo(DB *gorm.DB) domain.UserRepository {
	return mysqlUserRepo{DB}
}

func (repo mysqlUserRepo) Fetch(ctx context.Context) (result []domain.User, err error) {
	if err = repo.DB.WithContext(ctx).Find(&result).Error; err != nil {
		return
	}

	return
}

func (repo mysqlUserRepo) Create(ctx context.Context, user domain.User) (result domain.User, err error) {
	if err = repo.DB.WithContext(ctx).Create(&user).Error; err != nil {
		return
	}

	result = user
	return
}

func (repo mysqlUserRepo) Update(ctx context.Context, user domain.User) (result domain.User, err error) {
	if err = repo.DB.WithContext(ctx).Save(&user).Error; err != nil {
		return
	}

	result = user
	return
}

func (repo mysqlUserRepo) GetByID(ctx context.Context, id uint) (result domain.User, err error) {
	if err = repo.DB.WithContext(ctx).Where("id = ?", id).First(&result).Error; err != nil {
		return
	}

	return
}

func (repo mysqlUserRepo) Delete(ctx context.Context, id uint) (err error) {
	if err = repo.DB.WithContext(ctx).Where("id = ?", id).Delete(&domain.User{}).Error; err != nil {
		return
	}

	return
}
