package entity

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ProjectProposal struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Reason      string         `json:"reason"`
	UserID      uuid.UUID      `json:"user_id"`
	User        UserDB         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (p *ProjectProposal) UserEmail() string {
	return p.User.Info.Email
}

func (p *ProjectProposal) EmailMessageTextForOwner() string {
	return fmt.Sprintf(
		"Whoosh :)"+
			"\n"+"Project proposal from user with login: %s and email: %s"+
			"\n"+"Title: %s"+
			"\n"+"Description: %s"+
			"\n"+"Reason: %s",
		p.User.Login(),
		p.UserEmail(),
		p.Title,
		p.Description,
		p.Reason,
	)
}

func (p *ProjectProposal) EmailMessageTextForUser() string {
	return fmt.Sprintf(
		"Ehoooo :)"+
			"\n"+"Your project proposal for project with title %s"+
			"\n"+"was send to Dima"+
			"\n"+"Have a good day ;)",
		p.Title,
	)
}

type SendProposalReq struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Reason      string    `json:"reason"`
	UserID      uuid.UUID `json:"user_id"`
}

type SendProposalResp struct {
	ProposalID  uuid.UUID `json:"proposal_id"`
	Title       string    `json:"proposal_title"`
	Description string    `json:"proposal_description"`
	Reason      string    `json:"proposal_reason"`
}

func NewSendProposalResp() *SendProposalResp {
	return &SendProposalResp{}
}

func (resp *SendProposalResp) FromProjectProposal(p *ProjectProposal) *SendProposalResp {
	resp.ProposalID = p.ID
	resp.Title = p.Title
	resp.Reason = p.Reason
	resp.Description = p.Description
	return resp
}

func (req *SendProposalReq) ToProjectProposal() *ProjectProposal {
	return &ProjectProposal{
		Title:       req.Title,
		Description: req.Description,
		Reason:      req.Reason,
		UserID:      req.UserID,
	}
}
