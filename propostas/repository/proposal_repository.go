package repository

import (
	"database/sql"

	"github.com/pastorilps/propostas_populares/propostas/domain"
	"github.com/pastorilps/propostas_populares/propostas/entity"
	"github.com/sirupsen/logrus"
)

type proposalReposiroty struct {
	DbConn *sql.DB
}

func NewProposalRepo(DbConn *sql.DB) domain.ProposalRepository {
	return &proposalReposiroty{DbConn}
}

func (p *proposalReposiroty) CreateProposal(ep *entity.Send_Proposal_Data) (er *entity.Receive_Proposal_Data, err error) {
	query := `insert into public.proposal (title,pictures,attachments,description,status,user_id) values ($1,$2,$3,$4,$5,$6)`

	stmt, err := p.DbConn.Prepare(query)
	if err != nil {
		logrus.Error("Error in pushing data in database", err)
		return
	}

	res, err := stmt.Exec(
		ep.ProposalTitle,
		ep.ProposalPictures,
		ep.ProposalAttachments,
		ep.ProposalDescription,
		ep.ProposalStatus,
		ep.ProposalUserID,
	)
	if err != nil {
		logrus.Error("Error to attach data struct database", err)
		return
	}

	affect, err := res.RowsAffected()
	if affect != 1 {
		logrus.Error("Error in affect in database, err")
		return
	}

	return
}
