package models

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	id      int      `json:"id"`
	name    string   `json:"name"`
	salary  int      `json:"salary"`
	project *Project `json:"id"`
	manager *Manager `json:"manager"`
}

type Project struct {
	id          int    `json:"id"`
	projectname string `json:"projectname"`
	billing     int    `json:"billing"`
	duration    int    `json:"duration"`
}

type Manager struct {
	id   int    `json:"id"`
	name string `json:"name"`
}
