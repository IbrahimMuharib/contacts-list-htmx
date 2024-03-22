package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)

}

func newTemplate() *Templates {
	return &Templates{templates: template.Must(template.ParseGlob("views/*.html"))}
}

type Contact struct {
	Name  string
	Email string
}

func newContact(name, email string) Contact {
	return Contact{Name: name, Email: email}
}

type Contacts = []Contact

type Data struct {
	Contacts Contacts
}

func newData() Data {
	return Data{Contacts: []Contact{newContact("a", "a@g.com")}}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/css", "css")

	e.Renderer = newTemplate()

	data := newData()
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", data)
	})

	e.POST("/contacts", func(c echo.Context) error {
		data.Contacts = append(data.Contacts, newContact(c.FormValue("name"), c.FormValue("email")))
		fmt.Println(data)
		return c.Render(http.StatusOK, "contacts", data)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
