package proposalrepo

import (
	"context"
	"fmt"
	e "notify/internal/entity"
	"notify/internal/repo/pgrepo"
	"notify/pkg/pgerr"
)

func (r *Repository) Create(_ context.Context, proposal *e.ProjectProposal) (*e.ProjectProposal, error) {
	res := r.conn.Create(&proposal)
	if res.Error != nil {
		pgErr := pgerr.GetPgErr(res.Error)
		if pgErr != nil {
			switch pgErr.Code {
			case pgrepo.UNIQUE_CONSTRAINT_VIOLATED.Str():
				return nil, fmt.Errorf("proposal title=%s, user login=%s; %w: %w",
					proposal.Title, proposal.User.Cred.Login, ErrCreateProposal, ErrProposalAlreadyExists)
			}
		} else {
			return nil, fmt.Errorf(
				"proposal title=%s, user login=%s; %w: %w",
				proposal.Title, proposal.User.Cred.Login,
				ErrCreateProposal, res.Error)
		}
	}
	return proposal, res.Error
}
