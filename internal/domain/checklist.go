package domain

import "time"

type Checklist struct {
	Id        int
	Name      string
	UserId    int
	CreatedAt time.Time
}
