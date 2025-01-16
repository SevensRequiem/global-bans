package home

import (
	"fmt"
	"io"
	"os"

	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"

	"globalbans/backend/auth"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplateRenderer(glob string) *TemplateRenderer {
	tmpl := template.Must(template.ParseGlob(glob))
	return &TemplateRenderer{
		templates: tmpl,
	}
}

func LoginHandler(c echo.Context) error {
	tmpl, err := template.ParseFiles("frontend/views/base.html", "frontend/views/login.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	data := map[string]interface{}{}
	data = globaldata(c)
	data["Pagename"] = "Login"

	err = tmpl.ExecuteTemplate(c.Response().Writer, "base.html", data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func HomeHandler(c echo.Context) error {
	tmpl, err := template.ParseFiles("frontend/views/base.html", "frontend/views/home.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	data := map[string]interface{}{}
	data = globaldata(c)
	data["Pagename"] = "Home"

	err = tmpl.ExecuteTemplate(c.Response().Writer, "base.html", data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func BansHandler(c echo.Context) error {
	tmpl, err := template.ParseFiles("frontend/views/base.html", "frontend/views/bans.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	data := map[string]interface{}{}
	data = globaldata(c)
	data["Pagename"] = "Home"

	err = tmpl.ExecuteTemplate(c.Response().Writer, "base.html", data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func ServersHandler(c echo.Context) error {
	tmpl, err := template.ParseFiles("frontend/views/base.html", "frontend/views/servers.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	data := map[string]interface{}{}
	data = globaldata(c)
	data["Pagename"] = "Home"

	err = tmpl.ExecuteTemplate(c.Response().Writer, "base.html", data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func AppealsHandler(c echo.Context) error {
	tmpl, err := template.ParseFiles("frontend/views/base.html", "frontend/views/appeals.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	data := map[string]interface{}{}
	data = globaldata(c)
	data["Pagename"] = "Home"

	err = tmpl.ExecuteTemplate(c.Response().Writer, "base.html", data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func DocsHandler(c echo.Context) error {
	tmpl, err := template.ParseFiles("frontend/views/base.html", "frontend/views/docs.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	data := map[string]interface{}{}
	data = globaldata(c)
	data["Pagename"] = "Home"

	err = tmpl.ExecuteTemplate(c.Response().Writer, "base.html", data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return nil
}
func ErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	data := map[string]interface{}{
		"code":  code,
		"error": err.Error(),
	}

	if err := c.Render(code, "error.html", data); err != nil {
		c.Logger().Error(err)
	}
}

func globaldata(c echo.Context) map[string]interface{} {
	data := map[string]interface{}{}

	data["Auth"] = auth.AuthCheck(c)
	data["IsAdmin"] = auth.IsAdmin(c)
	data["IsMod"] = auth.IsMod(c)
	data["IP"] = c.RealIP()
	data["csrf"] = c.Get("csrf")
	data["BASEURL"] = os.Getenv("BASEURL")
	return data
}

func AdminHandler(c echo.Context) error {
	tmpl, err := template.ParseFiles("frontend/views/base.html", "frontend/admin/admin.html", "frontend/admin/home.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	data := map[string]interface{}{}
	data = globaldata(c)
	data["Pagename"] = "Admin"

	err = tmpl.ExecuteTemplate(c.Response().Writer, "base.html", data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func AdminDashboardHandler(c echo.Context) error {
	tmpl, err := template.ParseFiles("frontend/views/base.html", "frontend/admin/admin.html", "frontend/admin/dashboard.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	data := globaldata(c)
	data["Pagename"] = "Dashboard"
	err = tmpl.ExecuteTemplate(c.Response().Writer, "base.html", data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func AdminBansHandler(c echo.Context) error {
	tmpl, err := template.ParseFiles("frontend/views/base.html", "frontend/admin/admin.html", "frontend/admin/bans.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	data := globaldata(c)
	data["Pagename"] = "Bans"
	err = tmpl.ExecuteTemplate(c.Response().Writer, "base.html", data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func AdminServersHandler(c echo.Context) error {
	tmpl, err := template.ParseFiles("frontend/views/base.html", "frontend/admin/admin.html", "frontend/admin/servers.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	data := globaldata(c)
	data["Pagename"] = "Servers"
	err = tmpl.ExecuteTemplate(c.Response().Writer, "base.html", data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func AdminSettingsHandler(c echo.Context) error {
	tmpl, err := template.ParseFiles("frontend/views/base.html", "frontend/admin/admin.html", "frontend/admin/settings.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	data := globaldata(c)
	data["Pagename"] = "Settings"
	err = tmpl.ExecuteTemplate(c.Response().Writer, "base.html", data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return nil
}
