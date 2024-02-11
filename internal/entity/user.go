package entity

import (
	"gorm.io/gorm"
	"time"
)

type UserRole int

const (
	Admin UserRole = iota
	Service
	Person
)

type User struct {
	gorm.Model
	Info UserCommonInfo `gorm:"embedded" json:"info"`
	Cred UserCred       `gorm:"embedded" json:"req"`
}

type UserCommonInfo struct {
	Name     string    `gorm:"default:null" json:"name"`
	Email    string    `gorm:"default:null" json:"email"`
	Birthday time.Time `gorm:"default:null" json:"birthday"`
	Role     UserRole  `gorm:"default:2" json:"role"`
}

type UserCred struct {
	Login          string `json:"login"`
	HashedPassword string `json:"password"`
}
