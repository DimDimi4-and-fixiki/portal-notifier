package entity

import (
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
	"time"
)

type UserRole string

const (
	Admin   UserRole = "admin"
	Service UserRole = "service"
	Person  UserRole = "person"
)

// UserDB corresponds to rows from users table in DB
type UserDB struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Info      UserCommonInfo `gorm:"embedded" json:"info"`
	Cred      HashedUserCred `gorm:"embedded" json:"req"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u UserDB) Name() string {
	return u.Info.Name
}

func (u UserDB) Login() string {
	return u.Cred.Login
}

func (UserDB) TableName() string {
	return "users"
}

func (u UserDB) Email() string { return u.Info.Email }

// User contains non-hashed data to pass through the app
type User struct {
	Info UserCommonInfo `json:"info"`
	Cred UserCred       `json:"cred"`
}

type ServiceUser struct {
	Info ServiceUserCommonInfo `json:"common_info"`
	Cred ServiceCred           `json:"cred"`
}

type ServiceUserCommonInfo struct {
	Name string `gorm:"default:null" json:"name" binding:"required"`
}

type ServiceCred struct {
	ApiKey string `json:"api_key"`
	Login  string `json:"login"`
}

type UserCommonInfo struct {
	Name     string    `gorm:"default:null" json:"name" binding:"required"`
	Email    string    `gorm:"default:null" json:"email" binding:"email"`
	Birthday time.Time `gorm:"default:null" json:"birthday"`
	Role     UserRole  `gorm:"default:person" json:"role"`
}

type UserCred struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// HashedUserCred represents User credentials stored in DB
type HashedUserCred struct {
	Login          string `gorm:"unique" json:"login"`
	HashedPassword string `json:"password"`
}

type GetUserInput struct {
	ID    string `json:"user_id"`
	Login string `json:"user_login"`
}
