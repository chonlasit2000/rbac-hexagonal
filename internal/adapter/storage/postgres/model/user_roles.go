package model

type UserRole struct {
	Model
	UserID uint `gorm:"primaryKey"`
	RoleID uint `gorm:"primaryKey"`
}
