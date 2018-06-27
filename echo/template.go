package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
  e := echo.New()
  renderer := &TemplateRenderer{
      templates: template.Must(template.ParseGlob("*.html")),
  }
  e.Renderer = renderer

  // Named route "foobar"
  e.GET("/something", func(c echo.Context) error {
      return c.Render(http.StatusOK, "something.html", map[string]interface{}{
          "name": "Dolly!",
      })
  }).Name = "foobar"

  log.Fatal(e.Start(":8000"))
}