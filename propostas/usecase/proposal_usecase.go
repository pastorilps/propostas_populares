package usecase

import (
	"github.com/pastorilps/propostas_populares/propostas/domain"
	"github.com/pastorilps/propostas_populares/propostas/entity"
	"github.com/sirupsen/logrus"
)

type proposalUseCase struct {
	proposalRepo domain.ProposalRepository
}

func NewProposalUseCase(dp domain.ProposalRepository) domain.ProposalUseCase {
	return &proposalUseCase{
		proposalRepo: dp,
	}
}

func (pu *proposalUseCase) CreateProposal(ep *entity.Send_Proposal_Data) (pr *entity.Receive_Proposal_Data, err error) {
	res, err := pu.proposalRepo.CreateProposal(ep)
	if err != nil {
		logrus.Error("Error return data for saving in database", err)
	}

	return res, nil
}
