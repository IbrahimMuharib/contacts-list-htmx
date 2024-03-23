package main

import (
	"htmx/cmd/types"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/css", "css")
	e.Static("/js", "js")
	e.Static("/images", "images")

	e.Renderer = types.NewTemplate()

	page := types.NewPage()
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", page)
	})

	e.POST("/contacts", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		page.Form = types.NewFormData()
		page.Form.Values["name"] = name
		page.Form.Values["email"] = email
		if page.Data.HasEmail(email) {
			page.Form.Errors["email"] = "Email already used"
			return c.Render(http.StatusUnprocessableEntity, "form", page.Form)
		}
		contact := types.NewContact(name, email)
		page.Data.Contacts = append(page.Data.Contacts, contact)
		c.Render(http.StatusOK, "form", page.Form)
		return c.Render(http.StatusOK, "oob-contact", contact)
	})

	e.DELETE("/contacts/:id", func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid id")
		}
		index := page.Data.IndexOf(id)
		if index == -1 {
			return c.JSON(http.StatusNotFound, "Contact not found")
		}
		page.Data.Contacts = append(page.Data.Contacts[:index], page.Data.Contacts[index+1:]...)
		return c.NoContent(http.StatusOK)

	})

	e.Logger.Fatal(e.Start(":8080"))
}
