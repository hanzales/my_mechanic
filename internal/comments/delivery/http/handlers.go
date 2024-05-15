package http

import (
	"MyMechanic/internal/models"
	"github.com/opentracing/opentracing-go"
	"net/http"
	"strconv"

	"MyMechanic/config"
	"MyMechanic/internal/comments"
	"MyMechanic/pkg/logger"
	"MyMechanic/pkg/utils"
	"github.com/labstack/echo/v4"
)

// Comments handlers
type commentsHandlers struct {
	cfg    *config.Config
	comUC  comments.UseCase
	logger logger.Logger
}

// NewCommentsHandlers Comments handlers constructor
func NewCommentsHandlers(cfg *config.Config, comUC comments.UseCase, logger logger.Logger) comments.Handlers {
	return &commentsHandlers{cfg: cfg, comUC: comUC, logger: logger}
}

// GetByID
// @Summary Get comment
// @Description Get comment by id
// @Tags Comments
// @Accept  json
// @Produce  json
// @Param id path int true "comment_id"
// @Success 200 {object} models.Comment
// @Failure 500 {object} httpErrors.RestErr
// @Router /comments/{id} [get]
func (h *commentsHandlers) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "commentsHandlers.GetByID")
		defer span.Finish()

		commID, err := strconv.Atoi(c.QueryParam("id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(models.ErrorResponse(err))
		}

		comment, err := h.comUC.GetByID(ctx, commID)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(models.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, models.NewSuccessResponse(comment))
	}
}

func (h *commentsHandlers) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "commentsHandlers.Delete")
		defer span.Finish()

		commID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(models.ErrorResponse(err))
		}

		if err = h.comUC.Delete(ctx, commID); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(models.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, models.NewEmptySuccessResponse())
	}
}

// çalışmıyor düzenlenecek
func (h *commentsHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "commentsHandlers.Create")
		defer span.Finish()

		user, err := utils.GetUserFromCtx(ctx)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(models.ErrorResponse(err))
		}

		comment := &models.AddCommentRequest{}
		user.Id = 1

		if err = utils.SanitizeRequest(c, comment); err != nil {
			return utils.ErrResponseWithLog(c, h.logger, err)
			// return err
		}

		createdComment, err := h.comUC.Create(ctx, comment)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(models.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, models.NewSuccessResponse(createdComment))
	}
}

func (h *commentsHandlers) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "commentsHandlers.Update")
		defer span.Finish()

		user, err := utils.GetUserFromCtx(ctx)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(models.ErrorResponse(err))
		}

		comment := &models.UpdateCommentRequest{}
		user.Id = 1

		if err = utils.SanitizeRequest(c, comment); err != nil {
			return utils.ErrResponseWithLog(c, h.logger, err)
			// return err
		}

		updatedComment, err := h.comUC.Update(ctx, comment)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(models.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, models.NewSuccessResponse(updatedComment))
	}
}

func (h *commentsHandlers) IncreaseLikeCount() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "commentsHandlers.IncreaseLikeCount")
		defer span.Finish()

		user, err := utils.GetUserFromCtx(ctx)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(models.ErrorResponse(err))
		}

		increaseLikeRequest := &models.IncreaseLikeRequest{}
		user.Id = 1

		if err = utils.SanitizeRequest(c, increaseLikeRequest); err != nil {
			return utils.ErrResponseWithLog(c, h.logger, err)
		}

		if err = h.comUC.IncreaseLikeCount(ctx, increaseLikeRequest); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(models.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, models.NewEmptySuccessResponse())
	}
}
