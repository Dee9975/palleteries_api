package models

import "github.com/kamva/mgm/v3"

type Employee struct {
	mgm.DefaultModel `bson:",inline"`
	ID               int    `json:"id" bson:"id"`
	Name             string `json:"name" bson:"name"`
}
