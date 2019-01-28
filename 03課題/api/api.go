package main

import (
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type CalcResult struct {
	Input   int    `json:"input"`
	Output  int    `json:"output"`
	Message string `json:"message"`
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("./*.tpl")),
	}
	e.Renderer = t

	e.GET("/form", func(c echo.Context) error {
		s := "test"
		return c.Render(http.StatusOK, "form", s)
	})

	e.GET("/twice/:input", func(c echo.Context) error {
		input, err := strconv.Atoi(c.Param("input"))
		if err != nil {
			r := &CalcResult{
				Message: "input type error.",
			}
			return c.JSON(http.StatusBadRequest, r)
		}
		return c.JSON(http.StatusOK, twice(input))
	})

	e.Logger.Fatal(e.Start(":9090"))
}

func twice(in int) *CalcResult {
	out := in * 2
	r := &CalcResult{
		Input:  in,
		Output: out,
	}
	return r
}
