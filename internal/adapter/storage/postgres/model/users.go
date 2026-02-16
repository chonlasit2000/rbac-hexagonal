package model

type User struct {
	Model
	Username string `gorm:"unique;not null"`
	Password string
	// Relationship (GORM จะรู้เองว่าต้องไปหาที่ตารางเชื่อม)
	Roles []*Role `gorm:"many2many:user_roles;"`
}
