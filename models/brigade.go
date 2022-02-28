package models

import "github.com/kamva/mgm/v3"

type Brigade struct {
	mgm.DefaultModel `bson:",inline"`
	ID               int          `json:"id" bson:"id"`
	Name             string       `json:"name" bson:"name"`
	Members          []TeamMember `json:"members" bson:"members"`
}
