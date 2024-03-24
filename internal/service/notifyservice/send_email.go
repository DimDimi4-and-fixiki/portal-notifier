package notifyservice

import (
	"context"
	e "notify/internal/entity"
)

//func appendError(mu *sync.Mutex, errs []error, err error) []error {
//	mu.Lock()
//	result := append(errs, err)
//	mu.Unlock()
//	return result
//
//}

func (s *Service) sendProposal(ctx context.Context, msg e.EmailMsg, email string) error {
	res, err := s.courierClient.SendTextMessage(ctx, msg)
	if err != nil {
		s.logger.Errorf("eror sending email to user %s, resp: %s, error: %s ", email, res, err)
		return err
	}
	return nil
}

func (s *Service) SendProposalToOwner(ctx context.Context, proposal *e.ProjectProposal, ownerEmail string) error {
	msg := e.NewEmailMsg().FromProposalForOwner(proposal, ownerEmail)
	return s.sendProposal(ctx, msg, ownerEmail)
}

func (s *Service) SendProposalToUser(ctx context.Context, proposal *e.ProjectProposal) error {
	msg := e.NewEmailMsg().FromProposalForUser(proposal, proposal.Email)
	return s.sendProposal(ctx, msg, proposal.Email)
}

func (s *Service) CreateProposal(ctx context.Context, proposal *e.ProjectProposalDB) error {
	_, err := s.proposalRepo.Create(ctx, proposal)
	return err
}
