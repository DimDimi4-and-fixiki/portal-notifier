package userrepo

import (
	"context"
	"gorm.io/gorm"
	"notify/internal/entity"
)

func (r *Repository) GetByID(_ context.Context, id uint) {
	r.conn.First(&entity.User{Model: gorm.Model{ID: id}})
}
