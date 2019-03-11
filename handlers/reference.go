package handlers

import (
	"kw-shortener/helpers"
	"kw-shortener/models"

	"github.com/labstack/echo"
)

// GetAllReference list of all url reference
func GetAllReference(c echo.Context) error {
	var reference []models.Reference

	err := models.SelectAllReference(&reference)

	if err != nil {
		return c.JSON(500, "Server may error")
	}

	return c.JSON(200, reference)
}

// InserReference basic API to insert the reference
// It return bad request if the request not same as the structure
// It return server errror if the database isn't connected
// It return created if the record sucess
func InsertReference(c echo.Context) (err error) {
	r := new(models.Reference)

	if err = c.Bind(r); err != nil {
		return c.JSON(400, "Bad Request")
	}

	// Change the pointer of bind json to random code
	r.Destination = helpers.GenerateRandomCode()

	insertReference := models.StoreReference(r)

	if insertReference != nil {
		return c.JSON(500, "Sorry server error")
	}

	response := helpers.Response{r, "Reference created"}

	return c.JSON(201, response)
}
