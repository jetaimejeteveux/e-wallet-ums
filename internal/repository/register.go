package repository

import (
	"context"

	"github.com/jetaimejeteveux/e-wallet-ums/models"
	"gorm.io/gorm"
)

type RegisterRepository struct {
	DB *gorm.DB
}

func (r *RegisterRepository) InsertUser(ctx context.Context, user *models.User) error {
	return r.DB.Create(&user).Error
}
