package main

import (
	"net/http"
	"strconv"
    "fmt"
    "github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// User
type User struct {
  Name  string `json:"name" form:"name" query:"name"`
  Email string `json:"email" form:"email" query:"email"`
}

type Article struct {
  ID  string `json:"id" form:"name" query:"name"`
  Title string `json:"title" form:"email" query:"email"`
  Body string `json:"body" form:"email" query:"email"`
}

type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

type (
    article struct {
        id       string    `json:"id"`
        title    string    `json:"title"`
        body     string    `json:"body"`
    }
)

var (
	users    = map[int]*user{}
	articles = map[string]*Article{}
	seq   = 1
)

func NewArticle() article {
    return article{}
}

//----------
// Handlers
//----------

func createArticle(c echo.Context) error {
    id := uuid.New().String()
    fmt.Println(id)
    u := &Article{ID: id}
	if err := c.Bind(u); err != nil {
		return err
	}
	articles[u.ID] = u
	return c.JSON(http.StatusCreated, u)
}

func getArticle(c echo.Context) error {
    id := c.Param("id")
	return c.JSON(http.StatusOK, articles[id])
}

func createUser(c echo.Context) error {
	u := &user{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

func updateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/users", createUser)
	e.POST("/articles", createArticle)
	e.GET("/articles/:id", getArticle)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
