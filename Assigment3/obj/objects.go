package obj

import (
	"gorm.io/gorm "
)

type Listofbook struct {
	gorm.Model
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Cost        int    `json:"cost"`
}

type JsonResponse struct {
	Type    string       `json:"type"`
	Data    Listofbook `json:"data"`
	Message string       `json:"message"`
	Datas []Listofbook `json:"datas"`
}
