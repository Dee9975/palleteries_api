package models

import "github.com/kamva/mgm/v3"

type TeamMember struct {
	mgm.DefaultModel `bson:",inline"`
	ID               int     `json:"id" bson:"id"`
	Name             string  `json:"name" bson:"name"`
	Salary           float64 `json:"salary" bson:"salary"`
	Forklift         bool    `json:"forklift" bson:"forklift"`
	Kalts            bool    `json:"kalts" bson:"kalts"`
	Hours            int     `json:"hours" bson:"hours"`
	ExtraHours       int     `json:"extra_hours" bson:"extra_hours"`
}
