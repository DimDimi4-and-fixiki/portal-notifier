package userrepo

import (
	"gorm.io/gorm"
	"notify/internal/entity"
)

type Repository struct {
	conn *gorm.DB
}

func New(conn *gorm.DB) *Repository {
	return &Repository{conn}
}

func (r *Repository) WithMigrate() (*Repository, error) {
	for _, e := range r.Entities() {
		err := r.Conn().AutoMigrate(e)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

func (r *Repository) Conn() *gorm.DB {
	return r.conn
}

func (r *Repository) Entities() []interface{} {
	return []interface{}{
		&entity.UserDB{},
	}
}
