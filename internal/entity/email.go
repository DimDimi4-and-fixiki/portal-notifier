package entity

import "fmt"

type EmailMsg struct {
	ToEmail string `json:"to_email"`
	Title   string `json:"title"`
	Text    string `json:"text"`
}

func NewEmailMsg() *EmailMsg {
	return &EmailMsg{}
}

func (e *EmailMsg) FromProposalForOwner(proposal *ProjectProposal, toEmail string) EmailMsg {
	e.Title = fmt.Sprintf("Project proposal from user with email=%s", proposal.Email)
	e.Text = proposal.EmailMessageForOwner()
	e.ToEmail = toEmail
	return *e
}

func (e *EmailMsg) FromProposalForUser(proposal *ProjectProposal, toEmail string) EmailMsg {
	e.Title = "Meow from Dima 🐱"
	e.Text = proposal.EmailMessageForUser()
	e.ToEmail = toEmail
	return *e
}
