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

type TextController struct {
	TextService services.TextServiceInterface
}

func (t *TextController) GetTexts(ctx echo.Context) error {
	reqCtx := ctx.Request().Context()

	texts, err := t.TextService.GetTexts(reqCtx)
	if err != nil {
		slog.ErrorContext(reqCtx, "error fetching texts", "error", err)
		return ctx.JSON(http.StatusInternalServerError, "error fetching texts")
	}

	return ctx.JSON(http.StatusOK, struct{ Texts []*structures.Text }{Texts: texts})
}

func (t *TextController) GetText(ctx echo.Context) error {
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

	text, err := t.TextService.GetTextByIdOrTitle(reqCtx, textId, title)
	if err != nil && err == gorm.ErrRecordNotFound {
		slog.ErrorContext(reqCtx, "text not found", "error", err)
		return ctx.JSON(http.StatusNotFound, "text not found")
	} else if err != nil {
		slog.ErrorContext(reqCtx, "error fetching text", "error", err)
		return ctx.JSON(http.StatusInternalServerError, "error fetching text")
	}

	return ctx.JSON(http.StatusOK, text)
}

func (t *TextController) CreateText(ctx echo.Context) error {
	reqCtx := ctx.Request().Context()
	req := structures.TextReq{}

	err := ctx.Bind(&req)
	if err != nil {
		slog.ErrorContext(reqCtx, "error binding request body", "error", err)
		return ctx.JSON(http.StatusBadRequest, "error binding request body, incomplete or bad request")
	}

	createdText, err := t.TextService.CreateText(reqCtx, req)
	if err != nil {
		slog.ErrorContext(reqCtx, "error creating new text", "error", err)
		return ctx.JSON(http.StatusInternalServerError, "error creating new text")
	}

	return ctx.JSON(http.StatusCreated, createdText)
}

func (t *TextController) UpdateText(ctx echo.Context) error {
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

	updatedText, err := t.TextService.UpdateText(reqCtx, req, textId)
	if err != nil {
		slog.ErrorContext(reqCtx, "error updating text", "error", err)
		return ctx.JSON(http.StatusInternalServerError, "error updating text")
	}

	return ctx.JSON(http.StatusOK, updatedText)
}

func (t *TextController) DeleteText(ctx echo.Context) error {
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

	text, err := t.TextService.DeleteText(reqCtx, textId)
	if err != nil && err == gorm.ErrRecordNotFound {
		slog.ErrorContext(reqCtx, "text not found", "error", err)
		return ctx.JSON(http.StatusNotFound, "text not found")
	} else if err != nil {
		slog.ErrorContext(reqCtx, "error deleting text", "error", err)
		return ctx.JSON(http.StatusInternalServerError, "error deleting text")
	}

	return ctx.JSON(http.StatusOK, text)
}

func NewTextController(textService *services.TextService) *TextController {
	return &TextController{
		TextService: textService,
	}
}
