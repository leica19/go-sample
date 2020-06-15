package main

import (
	"app/db"
	"app/models"
)

func main() {
	db := db.GetConnection()
	pTrue := &[]bool{true}[0]
	pFalse := &[]bool{false}[0]
	db.Create(&models.User{
		TenantId: 1,
		Name:     "Taro",
		Age:      10,
		Sex:      "man",
		Skill: models.Skill{
			Golang: pTrue,
			Docker: pTrue,
			Rust:   pFalse,
			PHP:    pTrue,
		},
	})
	db.Create(&models.User{
		TenantId: 1,
		Name:     "Jiro",
		Age:      11,
		Sex:      "man",
		Skill: models.Skill{
			Golang: pTrue,
			Docker: pTrue,
			Rust:   pTrue,
			PHP:    pFalse,
		},
	})
	db.Create(&models.User{
		TenantId: 1,
		Name:     "Mika",
		Age:      12,
		Sex:      "female",
		Skill: models.Skill{
			Golang: pFalse,
			Docker: pTrue,
			Rust:   pFalse,
		},
	})
	db.Create(&models.User{
		TenantId: 2,
		Name:     "Noriko",
		Age:      10,
		Sex:      "female",
		Skill: models.Skill{
			Golang: pTrue,
			Docker: pTrue,
			Rust:   pFalse,
		},
	})
	db.Create(&models.User{
		TenantId: 3,
		Name:     "Noriko",
		Age:      22,
		Sex:      "man",
		Skill: models.Skill{
			Golang: pFalse,
			Docker: pTrue,
			Rust:   pTrue,
		},
	})
	defer db.Close()
}
