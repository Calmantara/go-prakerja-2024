package model

type Order struct {
	ID         uint64 `json:"id" gorm:"column:id"`
	UserID     uint64 `json:"user_id" gorm:"column:user_id"`
	OrderName  string `json:"order_name" gorm:"column:order_name"`
	UserDetail *User  `json:"user_detail" gorm:"foreignKey:UserID"` // preloading
}
