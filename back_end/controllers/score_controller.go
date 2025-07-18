package controllers

import (
	"log/slog"
	"net/http"
	"strconv"
	"type_writer_api/services"
	"type_writer_api/structures"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ScoreController struct {
	ScoreService services.ScoreServiceInterface
}

func (t *ScoreController) GetScores(ctx echo.Context) error {
	reqCtx := ctx.Request().Context()

	scores, err := t.ScoreService.GetScores(reqCtx)
	if err != nil {
		slog.ErrorContext(reqCtx, "error fetching scores", "error", err)
		return ctx.JSON(http.StatusInternalServerError, "error fetching scores")
	}

	return ctx.JSON(http.StatusOK, struct{ Scores []*structures.Score }{Scores: scores})
}

func (t *ScoreController) GetScore(ctx echo.Context) error {
	reqCtx := ctx.Request().Context()
	var (
		scoreId int
		err     error
	)

	scoreId, err = strconv.Atoi(ctx.Param("score_id"))
	if err != nil {
		slog.ErrorContext(reqCtx, "bad score id in request", "error", err)
		return ctx.JSON(http.StatusBadRequest, "bad score id in request")
	}

	score, err := t.ScoreService.GetScoreById(reqCtx, scoreId)
	if err != nil && err == gorm.ErrRecordNotFound {
		slog.ErrorContext(reqCtx, "score not found", "error", err)
		return ctx.JSON(http.StatusNotFound, "score not found")
	} else if err != nil {
		slog.ErrorContext(reqCtx, "error fetching score", "error", err)
		return ctx.JSON(http.StatusInternalServerError, "error fetching score")
	}

	return ctx.JSON(http.StatusOK, score)
}

func (t *ScoreController) CreateScore(ctx echo.Context) error {
	reqCtx := ctx.Request().Context()
	req := structures.ScoreReq{}

	err := ctx.Bind(&req)
	if err != nil {
		slog.ErrorContext(reqCtx, "error binding request body", "error", err)
		return ctx.JSON(http.StatusBadRequest, "error binding request body, incomplete or bad request")
	}

	createdScore, err := t.ScoreService.CreateScore(reqCtx, req)
	if err != nil {
		slog.ErrorContext(reqCtx, "error creating new score", "error", err)
		return ctx.JSON(http.StatusInternalServerError, "error creating new score")
	}

	return ctx.JSON(http.StatusCreated, createdScore)
}

func (t *ScoreController) UpdateScore(ctx echo.Context) error {
	reqCtx := ctx.Request().Context()
	var (
		req     structures.ScoreReq
		scoreId int
		err     error
	)

	scoreId, err = strconv.Atoi(ctx.Param("score_id"))
	if err != nil {
		slog.ErrorContext(reqCtx, "bad score id in request", "error", err)
		return ctx.JSON(http.StatusBadRequest, "bad score id in request")
	}

	err = ctx.Bind(&req)
	if err != nil {
		slog.ErrorContext(reqCtx, "error binding request body", "error", err)
		return ctx.JSON(http.StatusBadRequest, "error binding request body, incomplete or bad request")
	}

	updatedScore, err := t.ScoreService.UpdateScore(reqCtx, req, scoreId)
	if err != nil {
		slog.ErrorContext(reqCtx, "error updating score", "error", err)
		return ctx.JSON(http.StatusInternalServerError, "error updating score")
	}

	return ctx.JSON(http.StatusOK, updatedScore)
}

func (t *ScoreController) DeleteScore(ctx echo.Context) error {
	reqCtx := ctx.Request().Context()
	var (
		scoreId int
		err     error
	)

	scoreId, err = strconv.Atoi(ctx.Param("score_id"))
	if err != nil {
		slog.ErrorContext(reqCtx, "bad score id in request", "error", err)
		return ctx.JSON(http.StatusBadRequest, "bad score id in request")
	}

	score, err := t.ScoreService.DeleteScore(reqCtx, scoreId)
	if err != nil && err == gorm.ErrRecordNotFound {
		slog.ErrorContext(reqCtx, "score not found", "error", err)
		return ctx.JSON(http.StatusNotFound, "score not found")
	} else if err != nil {
		slog.ErrorContext(reqCtx, "error deleting score", "error", err)
		return ctx.JSON(http.StatusInternalServerError, "error deleting score")
	}

	return ctx.JSON(http.StatusOK, score)
}

func NewScoreController(scoreService *services.ScoreService) *ScoreController {
	return &ScoreController{
		ScoreService: scoreService,
	}
}
