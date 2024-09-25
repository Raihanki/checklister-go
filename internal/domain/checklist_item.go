package domain

import "time"

type ChecklistItem struct {
	Id          int
	ChecklistId int
	ItemName    string
	CompletedAt *time.Time
	CretedAt    *time.Time
}
