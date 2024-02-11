package entity

import "gorm.io/gorm"

type ProjectProposal struct {
	gorm.Model
	title       string
	description string
	reason      string
	user        User
}
