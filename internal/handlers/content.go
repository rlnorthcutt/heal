package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rlnorthcutt/heal/internal/models"
)

func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to the HELP Framework!")
}

func ListContent(c echo.Context) error {
	content, err := models.GetAllContent()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, content)
}

func CreateContent(c echo.Context) error {
	var content models.Content
	if err := c.Bind(&content); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := models.InsertContent(content); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, content)
}
