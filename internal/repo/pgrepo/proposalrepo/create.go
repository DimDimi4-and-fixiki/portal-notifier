package proposalrepo

import (
	"context"
	"fmt"
	e "notify/internal/entity"
	"notify/internal/repo/pgrepo"
	"notify/pkg/pgerr"
)

func (r *Repository) Create(_ context.Context, p *e.ProjectProposalDB) (*e.ProjectProposalDB, error) {
	res := r.conn.Create(&p)
	if res.Error != nil {
		pgErr := pgerr.GetPgErr(res.Error)
		if pgErr != nil {
			switch pgErr.Code {
			case pgrepo.UNIQUE_CONSTRAINT_VIOLATED.Str():
				return nil, fmt.Errorf("p title=%s, user login=%s; %w: %w",
					p.Proposal.Title, p.User.Cred.Login, ErrCreateProposal, ErrProposalAlreadyExists)
			}
		} else {
			return nil, fmt.Errorf(
				"p title=%s, user login=%s; %w: %w",
				p.Proposal.Title, p.User.Cred.Login,
				ErrCreateProposal, res.Error)
		}
	}
	return p, res.Error
}
