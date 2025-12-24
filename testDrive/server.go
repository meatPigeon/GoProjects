package main

import (
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Template renderer for HTML templates.
type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// User used for binding examples.
type User struct {
	Name  string `json:"name"  xml:"name"  form:"name"  query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func main() {
	e := echo.New()

	// Root-level middleware (from quick-start)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Template renderer (for HTML responses)
	e.Renderer = &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	// ---- Hello, World! ----
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// ---- Routing & path parameters ----
	// /users        - POST  → saveUser
	// /users/:id    - GET   → getUser
	// /users/:id    - PUT   → updateUser
	// /users/:id    - DELETE→ deleteUser
	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	// ---- Query parameters ----
	// Example URL: /show?team=x-men&member=wolverine
	e.GET("/show", showTeam)

	// ---- Forms: application/x-www-form-urlencoded ----
	// Example curl:
	// curl -d "name=Joe Smith" -d "[email protected]" http://localhost:1323/save
	e.POST("/save", saveForm)

	// ---- Forms: multipart/form-data (file upload) ----
	// Example curl:
	// curl -F "name=Joe Smith" -F "avatar=@/path/to/avatar.png" http://localhost:1323/upload
	e.POST("/upload", uploadAvatar)

	// ---- JSON/XML binding & response ----
	// POST /bind with JSON, form, query, etc.
	e.POST("/bind", bindUser)

	// ---- Static files ----
	// Serves files under ./static at /static/*
	e.Static("/static", "static")

	// ---- Template rendering example ----
	e.GET("/hello-template", func(c echo.Context) error {
		data := map[string]interface{}{
			"Title": "Echo Quickstart",
			"Name":  c.QueryParam("name"),
		}
		if data["Name"] == "" {
			data["Name"] = "Anonymous Gopher"
		}
		return c.Render(http.StatusOK, "hello.html", data)
	})

	// ---- Group-level middleware (basic auth) ----
	admin := e.Group("/admin")
	admin.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))

	admin.GET("/dashboard", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to the admin dashboard!")
	})

	// ---- Route-level middleware example ----
	track := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("request to /tracked/users")
			return next(c)
		}
	}
	e.GET("/tracked/users", func(c echo.Context) error {
		return c.String(http.StatusOK, "/tracked/users")
	}, track)

	// Start server on 1323 like in docs
	e.Logger.Fatal(e.Start(":1323"))
}

// ---------------- Handlers ----------------

// saveUser demonstrates binding JSON/form/query into a User.
func saveUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bind error: "+err.Error())
	}
	return c.JSON(http.StatusCreated, u)
}

// getUser reads :id from path.
func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "user id from path: "+id)
}

// updateUser just echos the id (placeholder for real logic).
func updateUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "update user "+id)
}

// deleteUser just echos the id (placeholder for real logic).
func deleteUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "delete user "+id)
}

// showTeam demonstrates query parameters: /show?team=x-men&member=wolverine
func showTeam(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

// saveForm handles application/x-www-form-urlencoded forms.
func saveForm(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

// uploadAvatar handles multipart/form-data with file upload.
func uploadAvatar(c echo.Context) error {
	name := c.FormValue("name")

	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")
}

// bindUser shows generic binding + JSON response.
func bindUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}
