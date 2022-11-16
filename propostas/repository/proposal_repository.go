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

func (u *proposalReposiroty) GetAllProposal() ([]*entity.Proposal_Data, error) {
	query := `select * from public.proposal order by id desc`

	return u.getAllProposal(query)
}

func (u *proposalReposiroty) getAllProposal(query string, args ...interface{}) ([]*entity.Proposal_Data, error) {
	rows, err := u.DbConn.Query(query, args...)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()

	result := make([]*entity.Proposal_Data, 0)
	for rows.Next() {
		gap := new(entity.Proposal_Data)
		err = rows.Scan(
			&gap.ProposalId,
			&gap.ProposalTitle,
			&gap.ProposalPictures,
			&gap.ProposalAttachments,
			&gap.ProposalDescription,
			&gap.ProposalStatus,
			&gap.ProposalUserID,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, gap)
	}

	return result, nil
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
