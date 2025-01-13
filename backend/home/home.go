package home

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"runtime"

	"globalbans/backend/auth"
	"globalbans/backend/bans"
	"globalbans/backend/logs"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	baseTemplate := template.New("base.html")
	baseTemplate, err := baseTemplate.ParseFiles("frontend/views/base.html", fmt.Sprintf("frontend/views/%s.html", name))
	if err != nil {
		return err
	}
	return baseTemplate.ExecuteTemplate(w, "base.html", data)
}

func NewTemplateRenderer(glob string) *TemplateRenderer {
	tmpl := template.Must(template.ParseGlob(glob))
	return &TemplateRenderer{
		templates: tmpl,
	}
}

func (t *TemplateRenderer) LoadTemplates() {
	t.templates = template.Must(template.ParseGlob("frontend/views/*.html"))
}

func NewRenderer() *TemplateRenderer {
	return &TemplateRenderer{}
}

var baseurl string

func init() {
	renderer := NewRenderer()
	renderer.LoadTemplates()
	baseurl = os.Getenv("BASE_URL")
}
func RenderPage(c echo.Context, page string, data map[string]interface{}) error {
	renderer := c.Get("renderer").(*TemplateRenderer)
	err := renderer.Render(c.Response().Writer, page, data, c)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}
	return nil
}

func HomeHandler(c echo.Context) error {
	data := make(map[string]interface{})
	data["title"] = "Home"
	data = globaldata(c)
	return RenderPage(c, "home", data)
}

func AdminHandler(c echo.Context) error {
	data := make(map[string]interface{})
	data["title"] = "Admin"
	data = globaldata(c)
	return RenderPage(c, "admin", data)
}

func globaldata(c echo.Context) map[string]interface{} {
	data := make(map[string]interface{})
	data["User"], _ = auth.GetCurrentUser(c)
	data["IsAdmin"] = auth.AdminCheck(c)
	data["IsBanned"] = bans.BannedCheck(c)
	data["BASEURL"] = baseurl
	data["csrf"] = c.Get("csrf")
	return data
}
