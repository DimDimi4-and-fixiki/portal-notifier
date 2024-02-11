package pgrepo

import "gorm.io/gorm"

// I use Gorm implementation for Repositories
// Entity and Connector types could be changed later for using pgx

type Repository interface {
	Entities() []interface{}
	Conn() *gorm.DB
}
