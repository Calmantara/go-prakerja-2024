package repository

import (
	"log"

	"github.com/Calmantara/go-prakerja-2024/deploy/model"
	"gorm.io/gorm"
)

type UserRepo interface {
	Migrate()
	Create(user *model.User) error
	Get() ([]*model.User, error)
	GetByUsername(username string) (*model.User, error)
}

type UserOrderRepo interface {
	GetByUsername(username string) (*model.User, error)
}

// SOLID PRINCIPLE PROGRAMMING GOLANG
// interface segregation

// constructor function
func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepoImpl{DB: db}
}

type userRepoImpl struct {
	DB *gorm.DB
}

func (u *userRepoImpl) Migrate() {
	// migration
	err := u.DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal(err)
	}
	// seeding
}

func (u *userRepoImpl) Create(user *model.User) error {
	err := u.DB.Debug().Model(&model.User{}).Create(user).Error
	return err
}

func (u *userRepoImpl) Get() ([]*model.User, error) {
	users := []*model.User{}
	err := u.DB.Debug().Model(&model.User{}).Find(&users).Error
	return users, err
}

func (u *userRepoImpl) GetByUsername(username string) (*model.User, error) {
	user := &model.User{}
	err := u.DB.Debug().Model(&model.User{}).Where("username = ?", username).First(&user).Error
	return user, err
}
