package entity

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectProposal struct {
	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Reason      string    `json:"reason"`
	UserRefer   uuid.UUID `json:"user_refer"`
	User        UserDB    `gorm:"foreignKey:UserRefer"`
}

func (p ProjectProposal) EmailMessageTextForOwner() string {
	return fmt.Sprintf(
		"Whoosh :)"+
			"\n"+"Project proposal from user with login: %s"+
			"\n"+"Title: %s"+
			"\n"+"Description: %s"+
			"\n"+"Reason: %s",
		p.User.Login(),
		p.Title,
		p.Description,
		p.Reason,
	)
}

func (p ProjectProposal) EmailMessageTextForUser() string {
	return fmt.Sprintf(
		"Ehoooo :)"+
			"\n"+"Your project proposal for project with title %s"+
			"\n"+"was send to Dima"+
			"\n"+"Have a good day ;)",
		p.Title,
	)
}
