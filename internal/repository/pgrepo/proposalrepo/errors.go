package proposalrepo

import "errors"

var ErrCreateProposal = errors.New("create proposal error")
var ErrGetUser = errors.New("get proposal error")
var ErrUpdateUser = errors.New("update proposal error")
var ErrProposalAlreadyExists = errors.New("proposal with this details already exists")
