package post

import (
	"myshelf/internal/modules/post/dto"
	"myshelf/pkg/logger"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Handler struct {
	Service PostService
}

func (h *Handler) CreatePost(c echo.Context) error {

	var req dto.CreatePostRequest

	if err := c.Bind(&req); err != nil {
		logger.Log.Warn("Failed to parse request", zap.Error(err))
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	logger.Log.Info("Received request to create post", zap.Int32("userID", req.UserID))

	err := h.Service.CreatePost(c.Request().Context(), req.UserID, req.Title, req.Content)
	if err != nil {
		logger.Log.Error("Post creation failed", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create post")
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Post created successfully",
	})
}

func (h *Handler) GetPostByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid post ID")
	}

	post, err := h.Service.GetPostByID(c.Request().Context(), int32(id))
	if err != nil {

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, post)
}

func (h *Handler) UpdatePost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid post ID")
	}

	var req struct {
		Content string `json:"content"`
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = h.Service.UpdatePost(c.Request().Context(), int32(id), req.Content)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Post updated successfully",
	})
}

func (h *Handler) DeletePost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid post ID")
	}

	err = h.Service.DeletePost(c.Request().Context(), int32(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Post deleted successfully",
	})
}
