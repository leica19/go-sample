package main

import (
	"app/db"
	"app/models"
)

func main() {
	db := db.GetConnection()
	db.AutoMigrate(&models.User{}, &models.Skill{})
	db.Create(&models.User{
		TenantId: 1,
		Name:     "Taro",
		Age:      10,
		Sex:      "man",
		Skill: models.Skill{
			Golang: true,
			Docker: true,
			Rust:   false,
		},
	})
	db.Create(&models.User{
		TenantId: 1,
		Name:     "Jiro",
		Age:      11,
		Sex:      "man",
		Skill: models.Skill{
			Golang: true,
			Docker: true,
			Rust:   true,
		},
	})
	db.Create(&models.User{
		TenantId: 1,
		Name:     "Mika",
		Age:      12,
		Sex:      "female",
		Skill: models.Skill{
			Golang: false,
			Docker: false,
			Rust:   false,
		},
	})
	db.Create(&models.User{
		TenantId: 2,
		Name:     "Noriko",
		Age:      10,
		Sex:      "female",
		Skill: models.Skill{
			Golang: true,
			Docker: true,
			Rust:   true,
		},
	})
	db.Create(&models.User{
		TenantId: 3,
		Name:     "Noriko",
		Age:      22,
		Sex:      "man",
		Skill: models.Skill{
			Golang: true,
			Docker: true,
			Rust:   true,
		},
	})
	defer db.Close()
}
