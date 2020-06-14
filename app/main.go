package main

import (
	"app/db"
	"app/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/tenants/:tenantId/users", func(c echo.Context) (err error) {
		tenantId := c.Param("tenantId")
		user := new(models.User)
		if err = c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, user)
		}
		user.TenantId, _ = strconv.Atoi(tenantId)
		db := db.GetConnection()
		db.Create(&user)
		defer db.Close()
		return c.JSON(http.StatusOK, user)
	})
	e.GET("/tenants/:tenantId/users", func(c echo.Context) (err error) {
		tenantId, _ := strconv.Atoi(c.Param("tenantId"))
		users := []models.User{}
		db := db.GetConnection()
		db.Preload("Skill").Where("tenant_id = ?", tenantId).Find(&users)
		defer db.Close()
		return c.JSON(http.StatusOK, users)
	})
	e.DELETE("/tenants/:tenantId/users", func(c echo.Context) (err error) {
		tenantId, _ := strconv.Atoi(c.Param("tenantId"))
		db := db.GetConnection()
		var skillIds []string
		// get skill.id to delete relation table
		db.Model(&models.User{}).Where("tenant_id = ?", tenantId).Pluck("skill_id", &skillIds)
		db.Where("tenant_id = ?", tenantId).Delete(models.User{})
		// delete relation table
		db.Where("id IN (?)", skillIds).Delete(models.Skill{})
		defer db.Close()
		return c.JSON(http.StatusOK, nil)
	})
	e.GET("/users", func(c echo.Context) (err error) {
		users := []models.User{}
		db := db.GetConnection()
		db.Preload("Skill").Find(&users)
		defer db.Close()
		return c.JSON(http.StatusOK, users)
	})
	e.GET("/users/:userId", func(c echo.Context) (err error) {
		userId, _ := strconv.Atoi(c.Param("userId"))
		user := models.User{}
		db := db.GetConnection()
		db.Preload("Skill").Where("id = ?", userId).Find(&user)
		defer db.Close()
		return c.JSON(http.StatusOK, user)
	})
	e.DELETE("/users/:userId", func(c echo.Context) (err error) {
		userId, _ := strconv.Atoi(c.Param("userId"))
		db := db.GetConnection()
		var skillId string
		// get skill.id to delete relation table
		db.Table("users").Select("skill_id").Where("id = ?", userId).Row().Scan(&skillId)
		db.Where("id = ?", userId).Delete(models.User{})
		// delete relation table
		db.Where("id = ?", skillId).Delete(models.Skill{})
		defer db.Close()
		return c.JSON(http.StatusOK, nil)
	})
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
