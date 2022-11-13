package http

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/pastorilps/propostas_populares/app/docs"
	middleware "github.com/pastorilps/propostas_populares/middlewares"
	"github.com/pastorilps/propostas_populares/propostas/domain"
	"github.com/pastorilps/propostas_populares/propostas/entity"
)

type Response struct {
	Message string `json:"message"`
}

type proposalHandler struct {
	AUsecase domain.ProposalUseCase
}

func NewProposalHandler(e echo.Echo, pu domain.ProposalUseCase) {
	handler := &proposalHandler{
		AUsecase: pu,
	}

	e.POST("/v1/proposal/create", handler.CreateProposal)
}

// CreateProposal godoc
// @Summary Create Proposal.
// @Description Create UsProposaler.
// @Tags Proposal
// @Accept json
// @Param Body body entity.Send_Proposal_Data true "The body to create a proposal"
// @Produce json
// @Success 200 {object} entity.Receive_Proposal_Data
// @Router /v1/peoposal/create [post]
func (ph *proposalHandler) CreateProposal(c echo.Context) error {
	p := new(entity.Send_Proposal_Data)
	if err := c.Bind(p); err != nil {
		return err
	}

	_, err := ph.AUsecase.CreateProposal(p)
	if err != nil {
		return c.JSON(getStatusCode(err), Response{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, p)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	log.Fatalln(err)

	switch err {
	case middleware.ErrInternalServerError:
		return http.StatusInternalServerError
	case middleware.ErrorNotFound:
		return http.StatusNotFound
	case middleware.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
