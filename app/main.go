package main

import (
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "fmt"
  "strconv"
)

type User struct {
  TenantId int  `json:"-"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Sex    string `json:"sex"`
  Skills Skills `json:"skills"`
}

type Skills struct {
  Golang bool `json:"golang"`
		Docker bool `json:"docker"`
		Rust   bool `json:"rust"`
}

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/", hello)
  e.POST("/user/:tenantId", func(c echo.Context) (err error) {
    tenantId := c.Param("tenantId")
    u := new(User)
    if err = c.Bind(u); err != nil {
      return c.JSON(http.StatusBadRequest, u)
    }
    u.TenantId, _ = strconv.Atoi(tenantId)
    fmt.Print(u)
    return c.JSON(http.StatusOK, u)
  })
  // Start server
  e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}
