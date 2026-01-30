package controllers

import (
	"log/slog"
	"net/http"
	"strconv"
	"type_writer_api/services/texts"
	"type_writer_api/structures"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TextsController struct {
	TextsService texts_service.TextsServiceInterface
}

func (t *TextsController) GetTexts(ctx echo.Context) error {
	reqCtx := ctx.Request().Context()

	texts, err := t.TextsService.GetTexts(reqCtx)
	if err != nil {
		slog.ErrorContext(reqCtx, "error fetching texts", "error", err)
		return ctx.JSON(http.StatusInternalServerError, "error fetching texts")
	}

	return ctx.JSON(http.StatusOK, struct{ Texts []*structures.Text }{Texts: texts})
}

func (t *TextsController) GetText(ctx echo.Context) error {
	reqCtx := ctx.Request().Context()
	var (
		textId int
		title  string
		err    error
	)

	textId, err = strconv.Atoi(ctx.Param("text_id"))
	if err != nil {
		title = ctx.Param("text_id")
	}

	text, err := t.TextsService.GetTextByIdOrTitle(reqCtx, &textId, &title)
	if err != nil && err == gorm.ErrRecordNotFound {
		slog.ErrorContext(reqCtx, "text not found", "error", err)
		return ctx.JSON(http.StatusNotFound, "text not found")
	} else if err != nil {
		slog.ErrorContext(reqCtx, "error fetching text", "error", err)
		return ctx.JSON(http.StatusInternalServerError, "error fetching text")
	}

	return ctx.JSON(http.StatusOK, text)
}

func (t *TextsController) CreateText(ctx echo.Context) error {
	reqCtx := ctx.Request().Context()
	req := structures.TextReq{}

	err := ctx.Bind(&req)
	if err != nil {
		slog.ErrorContext(reqCtx, "error binding request body", "error", err)
		return ctx.JSON(http.StatusBadRequest, "error binding request body, incomplete or bad request")
	}

	createdText, err := t.TextsService.CreateText(reqCtx, req)
	if err != nil {
		slog.ErrorContext(reqCtx, "error creating new text", "error", err)
		return ctx.JSON(http.StatusInternalServerError, "error creating new text")
	}

	return ctx.JSON(http.StatusCreated, createdText)
}

func (t *TextsController) UpdateText(ctx echo.Context) error {
	reqCtx := ctx.Request().Context()
	var (
		req    structures.TextReq
		textId int
		err    error
	)

	textId, err = strconv.Atoi(ctx.Param("text_id"))
	if err != nil {
		slog.ErrorContext(reqCtx, "bad text id in request", "error", err)
		return ctx.JSON(http.StatusBadRequest, "bad text id in request")
	}

	err = ctx.Bind(&req)
	if err != nil {
		slog.ErrorContext(reqCtx, "error binding request body", "error", err)
		return ctx.JSON(http.StatusBadRequest, "error binding request body, incomplete or bad request")
	}

	updatedText, err := t.TextsService.UpdateText(reqCtx, req, textId)
	if err != nil {
		slog.ErrorContext(reqCtx, "error updating text", "error", err)
		return ctx.JSON(http.StatusInternalServerError, "error updating text")
	}

	return ctx.JSON(http.StatusOK, updatedText)
}

func (t *TextsController) DeleteText(ctx echo.Context) error {
	reqCtx := ctx.Request().Context()
	var (
		textId int
		err    error
	)

	textId, err = strconv.Atoi(ctx.Param("text_id"))
	if err != nil {
		slog.ErrorContext(reqCtx, "bad text id in request", "error", err)
		return ctx.JSON(http.StatusBadRequest, "bad text id in request")
	}

	text, err := t.TextsService.DeleteText(reqCtx, textId)
	if err != nil && err == gorm.ErrRecordNotFound {
		slog.ErrorContext(reqCtx, "text not found", "error", err)
		return ctx.JSON(http.StatusNotFound, "text not found")
	} else if err != nil {
		slog.ErrorContext(reqCtx, "error deleting text", "error", err)
		return ctx.JSON(http.StatusInternalServerError, "error deleting text")
	}

	return ctx.JSON(http.StatusOK, text)
}

func NewTextsController(textsService *texts_service.TextsService) *TextsController {
	return &TextsController{
		TextsService: textsService,
	}
}
