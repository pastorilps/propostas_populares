package http

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/pastorilps/propostas_populares/app/docs"
	"github.com/pastorilps/propostas_populares/middlewares"
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

func NewProposalHandler(e *echo.Echo, pu domain.ProposalUseCase) {
	handler := &proposalHandler{
		AUsecase: pu,
	}

	e.POST("/v1/proposal/create", handler.CreateProposal)
	e.GET("/v1/proposal", handler.GetAllProposal)
}

// GetAllUsers godoc
// @Summary Show all users.
// @Description Get all users list.
// @Tags Proposal
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Accept */*
// @Produce json
// @Success 200 {object} entity.Proposal_Data
// @Router /v1/proposal [get]
func (ph *proposalHandler) GetAllProposal(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	data, err := ph.AUsecase.GetAllProposal()
	if err != nil {
		return c.JSON(getStatusCode(err), Response{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

// CreateProposal godoc
// @Summary Create Proposal.
// @Description Create UsProposaler.
// @Tags Proposal
// @Accept json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param Body body entity.Send_Proposal_Data true "The body to create a proposal"
// @Produce json
// @Success 200 {object} entity.Receive_Proposal_Data
// @Router /v1/proposal/create [post]
func (ph *proposalHandler) CreateProposal(c echo.Context) error {
	token := middlewares.GetToken(c)
	userId := middlewares.GetUserIDJWT(token)
	p := new(entity.Send_Proposal_Data)
	if err := c.Bind(p); err != nil {
		return err
	}

	p.ProposalUserID = userId

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
