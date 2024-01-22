package handlers

import (
	"encoding/json"
	"go_job_test/models"
	"io"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		slog.Error(err.Error())
	}
	var result models.Human
	if err := json.Unmarshal(body, &result); err != nil {
		slog.Debug("Error during Unmarshalling" + err.Error())
		c.Status(http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !is_valid(result) {
		c.Status(http.StatusBadRequest)
		slog.Debug("Cannot create with empty params")
		c.JSON(http.StatusBadRequest, "Empty name of surname")
		return
	}
	result.Enrich()
	h.DB.Create(&result)
	c.Status(http.StatusOK)
	c.JSON(200, result.ID)
	slog.Info("New User added")
}

func is_valid(h models.Human) bool {
	if h.Name == "" {
		return false
	}
	if h.Surname == "" {
		return false
	}
	return true
}
