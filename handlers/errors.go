package handlers

import (
	"errors"
	"log/slog"
	"net/http"
	"stats/services"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	logAttrs := []any{
		slog.String("path", c.Request.URL.Path),
		slog.String("method", c.Request.Method),
		slog.String("err_message", err.Error()),
	}

	switch {
	// 400 Bad Request
	case errors.Is(err, services.ErrInvalidLevelID),
		errors.Is(err, services.ErrInvalidSteps),
		errors.Is(err, services.ErrInvalidTime),
		errors.Is(err, services.ErrInvalidUsername):
		slog.Warn("client bad request", logAttrs...)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 500 Internal Server Err
	default:
		slog.Warn("internal server error detected", logAttrs...)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
	}
}
