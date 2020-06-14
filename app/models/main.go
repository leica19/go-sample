package models

import (
	"time"
)

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

type User struct {
	Model
	TenantId int    `json:"-"`
	Name     string `json:"name"`
	Age      uint   `json:"age"`
	Sex      string `json:"sex"`
	SkillId  uint   `json:"-"`
	Skill    Skill  `json:"skill"`
}

type Skill struct {
	Model
	Golang bool `json:"golang"`
	Docker bool `json:"docker"`
	Rust   bool `json:"rust"`
}