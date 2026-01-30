package controllers

import (
	"log/slog"
	"net/http"
	"strconv"
	"type_writer_api/services/activites"
	"type_writer_api/structures"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ActivitiesController struct {
	ActivitiesService activities_service.ActivitiesServiceInterface
}

func (t *ActivitiesController) GetActivities(ctx echo.Context) error {
	reqCtx := ctx.Request().Context()

	activities, err := t.ActivitiesService.GetActivities(reqCtx)
	if err != nil {
		slog.ErrorContext(reqCtx, "error fetching activities", "error", err)
		return ctx.JSON(http.StatusInternalServerError, "error fetching activities")
	}

	return ctx.JSON(http.StatusOK, struct{ Activities []*structures.Activity }{Activities: activities})
}

func (t *ActivitiesController) GetActivity(ctx echo.Context) error {
	reqCtx := ctx.Request().Context()
	var (
		activityId int
		name       string
		err        error
	)

	activityId, err = strconv.Atoi(ctx.Param("activity_id"))
	if err != nil {
		name = ctx.Param("activity_id")
	}

	activity, err := t.ActivitiesService.GetActivityByIdOrName(reqCtx, &activityId, &name)
	if err != nil && err == gorm.ErrRecordNotFound {
		slog.ErrorContext(reqCtx, "activity not found", "error", err)
		return ctx.JSON(http.StatusNotFound, "activity not found")
	} else if err != nil {
		slog.ErrorContext(reqCtx, "error fetching activity", "error", err)
		return ctx.JSON(http.StatusInternalServerError, "error fetching activity")
	}

	return ctx.JSON(http.StatusOK, activity)
}

func (t *ActivitiesController) CreateActivity(ctx echo.Context) error {
	reqCtx := ctx.Request().Context()
	req := structures.ActivityReq{}

	err := ctx.Bind(&req)
	if err != nil {
		slog.ErrorContext(reqCtx, "error binding request body", "error", err)
		return ctx.JSON(http.StatusBadRequest, "error binding request body, incomplete or bad request")
	}

	createdActivity, err := t.ActivitiesService.CreateActivity(reqCtx, req)
	if err != nil {
		slog.ErrorContext(reqCtx, "error creating new activity", "error", err)
		return ctx.JSON(http.StatusInternalServerError, "error creating new activity")
	}

	return ctx.JSON(http.StatusCreated, createdActivity)
}

func (t *ActivitiesController) UpdateActivity(ctx echo.Context) error {
	reqCtx := ctx.Request().Context()
	var (
		req        structures.ActivityReq
		activityId int
		err        error
	)

	activityId, err = strconv.Atoi(ctx.Param("activity_id"))
	if err != nil {
		slog.ErrorContext(reqCtx, "bad activity id in request", "error", err)
		return ctx.JSON(http.StatusBadRequest, "bad activity id in request")
	}

	err = ctx.Bind(&req)
	if err != nil {
		slog.ErrorContext(reqCtx, "error binding request body", "error", err)
		return ctx.JSON(http.StatusBadRequest, "error binding request body, incomplete or bad request")
	}

	updatedActivity, err := t.ActivitiesService.UpdateActivity(reqCtx, req, activityId)
	if err != nil {
		slog.ErrorContext(reqCtx, "error updating activity", "error", err)
		return ctx.JSON(http.StatusInternalServerError, "error updating activity")
	}

	return ctx.JSON(http.StatusOK, updatedActivity)
}

func (t *ActivitiesController) DeleteActivity(ctx echo.Context) error {
	reqCtx := ctx.Request().Context()
	var (
		activityId int
		err        error
	)

	activityId, err = strconv.Atoi(ctx.Param("activity_id"))
	if err != nil {
		slog.ErrorContext(reqCtx, "bad activity id in request", "error", err)
		return ctx.JSON(http.StatusBadRequest, "bad activity id in request")
	}

	activity, err := t.ActivitiesService.DeleteActivity(reqCtx, activityId)
	if err != nil && err == gorm.ErrRecordNotFound {
		slog.ErrorContext(reqCtx, "activity not found", "error", err)
		return ctx.JSON(http.StatusNotFound, "activity not found")
	} else if err != nil {
		slog.ErrorContext(reqCtx, "error deleting activity", "error", err)
		return ctx.JSON(http.StatusInternalServerError, "error deleting activity")
	}

	return ctx.JSON(http.StatusOK, activity)
}

func NewActivitiesController(activitiesService *activities_service.ActivitiesService) *ActivitiesController {
	return &ActivitiesController{
		ActivitiesService: activitiesService,
	}
}
