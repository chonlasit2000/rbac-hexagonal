package model

type Permission struct {
	Model
	Code  string  `gorm:"unique;not null"`
	Roles []*Role `gorm:"many2many:role_permissions;"`
}
