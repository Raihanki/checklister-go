package dto

import "time"

type ChecklistItemRequest struct {
	ItemName    string `json:"item_name"`
	ChecklistId int    `json:"-"`
}

type UpdateChecklistItemRequest struct {
	ItemName    string `json:"item_name"`
	ChecklistId int    `json:"-"`
}

type ChecklistItemResponse struct {
	Id          int        `json:"id"`
	ItemName    string     `json:"item_name"`
	ChecklistId int        `json:"checklist_id"`
	CompletedAt *time.Time `json:"completed_at"`
	CreatedAt   time.Time  `json:"created_at"`
}
