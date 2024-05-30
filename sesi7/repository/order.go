package repository

import (
	"log"

	"github.com/Calmantara/go-prakerja-2024/sesi7/model"
	"gorm.io/gorm"
)

type OrderRepo struct {
	DB *gorm.DB
}

func (u *OrderRepo) Migrate() {
	// migration
	err := u.DB.AutoMigrate(&model.Order{})
	if err != nil {
		log.Fatal(err)
	}
	// seeding
}

func (u *OrderRepo) Create(order *model.Order) error {
	err := u.DB.Debug().Model(&model.Order{}).Create(order).Error
	return err
}

func (u *OrderRepo) Get() ([]*model.Order, error) {
	orders := []*model.Order{}
	err := u.DB.Debug().Model(&model.Order{}).Preload("UserDetail").Find(&orders).Error
	return orders, err
}
