package models

import (
	"github.com/kamva/mgm/v3"
)

type Team struct {
	mgm.DefaultModel `bson:",inline"`
	ID               int          `json:"id" bson:"id"`
	Created          int          `json:"created" bson:"created"`
	Members          []TeamMember `json:"members" bson:"members"`
	Planks           []Plank      `json:"planks" bson:"planks"`
}
