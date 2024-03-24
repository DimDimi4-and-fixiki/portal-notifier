package entity

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectProposal struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Reason      string `json:"reason"`
	Email       string `json:"email"`
}

type ProjectProposalDB struct {
	gorm.Model
	Proposal  ProjectProposal `gorm:"embedded" json:"proposal"`
	UserRefer uuid.UUID       `json:"user_refer"`
	User      UserDB          `gorm:"foreignKey:UserRefer"`
}

func (p *ProjectProposal) EmailMessageForOwner() string {
	return fmt.Sprintf(
		"Whoosh :)"+
			"\n"+"Project proposal from user with email: %s"+
			"\n"+"Title: %s"+
			"\n"+"Description: %s"+
			"\n"+"Reason: %s",
		p.Email,
		p.Title,
		p.Description,
		p.Reason,
	)
}

func (p *ProjectProposal) EmailMessageForUser() string {
	return fmt.Sprintf(
		"Ehoooo :)"+
			"\n"+"Your project proposal for project with title %s"+
			"\n"+"was send to Dima"+
			"\n"+"Have a good day ;)",
		p.Title,
	)
}
