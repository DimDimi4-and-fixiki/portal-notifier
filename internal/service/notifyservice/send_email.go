package notifyservice

import (
	"context"
	e "notify/internal/entity"
)

func (s *Service) sendProposal(ctx context.Context, msg e.EmailMsg, email string) error {
	res, err := s.courierClient.SendTextMessage(ctx, msg)
	if err != nil {
		s.logger.Errorf("%w: error sending email to user %s, resp: %s, error: %s ", ErrSendProposal, email, res, err)
		return err
	}
	return nil
}

func (s *Service) SendProposalToOwner(ctx context.Context, proposal *e.ProjectProposal, ownerEmail string) error {
	msg := e.NewEmailMsg().FromProposalForOwner(*proposal, ownerEmail)
	return s.sendProposal(ctx, msg, ownerEmail)
}

func (s *Service) SendProposalToUser(ctx context.Context, proposal *e.ProjectProposal, userEmail string) error {
	msg := e.NewEmailMsg().FromProposalForUser(*proposal, userEmail)
	return s.sendProposal(ctx, msg, userEmail)
}

func (s *Service) SaveProposal(ctx context.Context, proposal *e.ProjectProposal) error {
	_, err := s.proposalRepo.Create(ctx, proposal)
	return err

}
