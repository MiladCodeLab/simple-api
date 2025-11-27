package application

import (
	"errors"
	"github.com/MiladCodeLab/simple-api/repository"
	"net/http"

	"github.com/MiladCodeLab/simple-api/dto"
	"github.com/MiladCodeLab/simple-api/entity"
	"github.com/MiladCodeLab/simple-api/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
)

type UserHandler struct {
	logger  *slog.Logger
	service service.UserService
}

func NewUserHandler(logger *slog.Logger, svc service.UserService) *UserHandler {
	return &UserHandler{
		logger:  logger.With("application", "user"),
		service: svc,
	}
}

func (h *UserHandler) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	users := api.Group("/users")

	users.GET("/", h.GetAll)
	users.GET("/:id", h.GetByID)
	users.POST("/", h.Add)
	users.DELETE("/:id", h.DeleteByID)
}

func (h *UserHandler) GetAll(c *gin.Context) {
	lg := h.logger.With("method", "GetAll")

	users, err := h.service.GetAll()
	if err != nil {
		lg.Error("failed fetching users", "error", err)
		JSONError(c, http.StatusInternalServerError, "failed to fetch users")
		return
	}

	JSONSuccess(c, http.StatusOK, users)
}

func (h *UserHandler) GetByID(c *gin.Context) {
	lg := h.logger.With("method", "GetByID")

	id := c.Param("id")
	if id == "" {
		JSONError(c, http.StatusBadRequest, "id required")
		return
	}

	// UUID validation
	if _, err := uuid.Parse(id); err != nil {
		JSONError(c, http.StatusBadRequest, "invalid uuid format")
		return
	}

	user, err := h.service.GetByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFoundUser) {
			lg.Warn("user not found", "id", id, "error", err)
			JSONError(c, http.StatusNotFound, "user not found")
			return
		}
		lg.Error("failed fetching user", "id", id, "error", err)
		JSONError(c, http.StatusInternalServerError, "failed to fetch user")
		return
	}

	JSONSuccess(c, http.StatusOK, user)
}

func (h *UserHandler) Add(c *gin.Context) {
	lg := h.logger.With("method", "Add")

	var body dto.UserDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		lg.Warn("validation failed", "error", err)
		JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	user := entity.User{
		ID:    uuid.New().String(),
		Name:  body.Name,
		Phone: body.Phone,
	}

	if err := h.service.Add(user); err != nil {
		lg.Error("failed to add user", "error", err)
		JSONError(c, http.StatusInternalServerError, "failed to create user")
		return
	}

	JSONSuccess(c, http.StatusCreated, user)
}

func (h *UserHandler) DeleteByID(c *gin.Context) {
	lg := h.logger.With("method", "DeleteByID")

	id := c.Param("id")
	if id == "" {
		JSONError(c, http.StatusBadRequest, "id required")
		return
	}

	// UUID validation
	if _, err := uuid.Parse(id); err != nil {
		JSONError(c, http.StatusBadRequest, "invalid uuid format")
		return
	}

	err := h.service.DeleteByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFoundUser) {
			lg.Warn("delete failed", "id", id, "error", err)
			JSONError(c, http.StatusNotFound, "user not found")
			return
		}
		lg.Error("failed to delete user", "error", err)
		JSONError(c, http.StatusInternalServerError, "failed to delete user")
		return
	}

	c.Status(http.StatusNoContent)
}
