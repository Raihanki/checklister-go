package dto

import "time"

type CreateChecklist struct {
	Name string `json:"name"`
}

type ChecklistResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
