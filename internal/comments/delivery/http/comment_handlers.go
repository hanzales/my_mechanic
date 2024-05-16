package http

import (
	"MyMechanic/internal/models"
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
	cfg         *config.Config
	commService comments.Service
	logger      logger.Logger
}

// CommentsHandlers Comments handlers constructor
func CommentsHandlers(cfg *config.Config, commService comments.Service, logger logger.Logger) comments.Handlers {
	return &commentsHandlers{cfg: cfg, commService: commService, logger: logger}
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
		commID, err := strconv.Atoi(c.QueryParam("id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(models.ErrorResponse(err))
		}

		comment, err := h.commService.GetByID(c.Request().Context(), commID)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(models.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, models.NewSuccessResponse(comment))
	}
}

func (h *commentsHandlers) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		commID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(models.ErrorResponse(err))
		}

		if err = h.commService.Delete(c.Request().Context(), commID); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(models.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, models.NewEmptySuccessResponse())
	}
}

func (h *commentsHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {

		comment := &models.AddCommentRequest{}
		err := utils.SanitizeRequest(c, comment)

		if err != nil {
			return utils.ErrResponseWithLog(c, h.logger, err)
		}

		createdComment, err := h.commService.Create(c.Request().Context(), comment)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(models.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, models.NewSuccessResponse(createdComment))
	}
}

func (h *commentsHandlers) Update() echo.HandlerFunc {
	return func(c echo.Context) error {

		comment := &models.UpdateCommentRequest{}
		err := utils.SanitizeRequest(c, comment)

		if err != nil {
			return utils.ErrResponseWithLog(c, h.logger, err)
		}

		updatedComment, err := h.commService.Update(c.Request().Context(), comment)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(models.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, models.NewSuccessResponse(updatedComment))
	}
}

func (h *commentsHandlers) IncreaseLikeCount() echo.HandlerFunc {
	return func(c echo.Context) error {

		increaseLikeRequest := &models.IncreaseLikeRequest{}
		err := utils.SanitizeRequest(c, increaseLikeRequest)

		if err != nil {
			return utils.ErrResponseWithLog(c, h.logger, err)
		}

		if err = h.commService.IncreaseLikeCount(c.Request().Context(), increaseLikeRequest); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(models.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, models.NewEmptySuccessResponse())
	}
}
