package handlers

import (
	"encoding/json"
	"go_job_test/models"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
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
	var target models.Human
	if resp := h.DB.First(&target, id); resp.Error != nil {
		slog.Error("Nothing to update")
		c.Status(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, "Nothing to update")
		return
	}
	target.Age = result.Age
	target.Patronimyc = result.Patronimyc
	target.Name = result.Name
	target.Nationality = result.Nationality
	target.Surname = result.Surname
	target.Gender = result.Gender
	h.DB.Save(&target)
	c.Status(http.StatusOK)
	c.JSON(200, target.ID)
}
