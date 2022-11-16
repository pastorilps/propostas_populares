package domain

import "github.com/pastorilps/propostas_populares/propostas/entity"

type ProposalUseCase interface {
	CreateProposal(ep *entity.Send_Proposal_Data) (pr *entity.Receive_Proposal_Data, err error)
	GetAllProposal() ([]*entity.Proposal_Data, error)
}

type ProposalRepository interface {
	CreateProposal(ep *entity.Send_Proposal_Data) (er *entity.Receive_Proposal_Data, err error)
	GetAllProposal() ([]*entity.Proposal_Data, error)
}
