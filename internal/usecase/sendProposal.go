package usecase

import (
	"context"
	"notify/internal/entity"
	"sync"
)

func appendError(mu *sync.Mutex, errs []error, err error) []error {
	mu.Lock()
	result := append(errs, err)
	mu.Unlock()
	return result
}

func (u *UseCase) SendProposal(ctx context.Context, proposal *entity.ProjectProposal, ownerEmail string) (*entity.ProjectProposal, []error) {
	err := u.notifyService.SendProposalToOwner(ctx, proposal, ownerEmail)
	if err != nil {
		return nil, []error{err}
	}

	var wg sync.WaitGroup
	errs := []error{}
	var mu sync.Mutex

	wg.Add(2)
	go func() {
		defer wg.Done()
		err = u.notifyService.SendProposalToUser(ctx, proposal, proposal.UserEmail())
		appendError(&mu, errs, err)
	}()

	go func() {
		defer wg.Done()
		err = u.notifyService.SaveProposal(ctx, proposal)
		appendError(&mu, errs, err)
	}()

	wg.Wait()
	return proposal, errs
}
