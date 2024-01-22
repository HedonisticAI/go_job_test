package handlers

import (
	"go_job_test/models"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.Status(http.StatusOK)
	slog.Debug("Request to delete data with id " + c.Param("id"))
	h.DB.Delete(&models.Human{}, id)
}
