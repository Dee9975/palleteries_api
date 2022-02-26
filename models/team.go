package models

import "time"

type Team struct {
	ID         int
	CreateTime time.Time
	Members    []TeamMember
	Planks     []Plank
}
