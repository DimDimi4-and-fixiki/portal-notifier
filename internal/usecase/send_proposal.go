package usecase

import (
	"context"
	"golang.org/x/sync/errgroup"
	cfg "notify/internal/config"
	e "notify/internal/entity"
)

func (u *UseCase) SendProposal(ctx context.Context, proposal e.ProjectProposal) error {
	user, err := u.authService.GetUserByEmail(ctx, proposal.Email)
	if err != nil {
		return err
	}
	proposalDB := e.ProjectProposalDB{Proposal: proposal, User: *user}
	err = u.notifyService.CreateProposal(ctx, &proposalDB)
	if err != nil {
		return err
	}

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return u.notifyService.SendProposalToUser(ctx, &proposal)
	})

	g.Go(func() error {
		return u.notifyService.SendProposalToOwner(ctx, &proposal, cfg.Get().Courier.OwnerEmail)
	})

	err = g.Wait()
	return err
}
