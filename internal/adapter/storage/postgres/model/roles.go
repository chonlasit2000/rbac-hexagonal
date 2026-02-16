package model

type Role struct {
	Model
	Name        string        `gorm:"unique;not null"`
	Permissions []*Permission `gorm:"many2many:role_permissions;"`
	Users       []*User       `gorm:"many2many:user_roles;"`
}
