package handlers

import (
	"go_job_test/models"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAll(c *gin.Context) {
	var tests []models.Human
	c.Status(http.StatusOK)
	slog.Debug("Request to Get All Data")
	h.DB.Find(&tests)
	c.JSON(200, tests)
}
func (h *Handler) GetByName(c *gin.Context) {

	var tests []models.Human
	c.Status(http.StatusOK)
	slog.Debug("Request to get data with name " + c.Param("Name"))
	h.DB.Where("name = ?", c.Param("name")).Find(&tests)
	c.JSON(200, tests)
}

func (h *Handler) GetBySurname(c *gin.Context) {

	var tests []models.Human
	c.Status(http.StatusOK)
	slog.Debug("Request to get data with surname " + c.Param("Name"))
	h.DB.Where("surname = ?", c.Param("surname")).Find(&tests)
	c.JSON(200, tests)
}

func (h *Handler) GetByAge(c *gin.Context) {

	var tests []models.Human
	c.Status(http.StatusOK)
	slog.Debug("Request to get data with  age" + c.Param("condition") + c.Param("age"))
	switch considition := c.Param("condition"); considition {
	case "=":
		h.DB.Where("age = ?", c.Param("age")).Find(&tests)
		c.JSON(200, tests)
	case ">":
		h.DB.Where("age > ?", c.Param("age")).Find(&tests)
		c.JSON(200, tests)
	case "<":
		h.DB.Where("age < ?", c.Param("age")).Find(&tests)
		c.JSON(200, tests)
	case "=>":
		h.DB.Where("age >= ?", c.Param("age")).Find(&tests)
		c.JSON(200, tests)
	case "<=":
		h.DB.Where("age =< ?", c.Param("age")).Find(&tests)
		c.JSON(200, tests)
	default:
		c.JSON(200, "No data found or bad condition")
	}
}
