package repository

import (
	"log"

	"github.com/Calmantara/go-prakerja-2024/sesi7/model"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func (u *UserRepo) Migrate() {
	// migration
	err := u.DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal(err)
	}
	// seeding
}

func (u *UserRepo) Create(user *model.User) error {
	err := u.DB.Debug().Model(&model.User{}).Create(user).Error
	return err
}

func (u *UserRepo) Get() ([]*model.User, error) {
	users := []*model.User{}
	err := u.DB.Debug().Model(&model.User{}).Find(&users).Error
	return users, err
}

func (u *UserRepo) GetByUsername(username string) (*model.User, error) {
	user := &model.User{}
	err := u.DB.Debug().Model(&model.User{}).Where("username = ?", username).First(&user).Error
	return user, err
}
