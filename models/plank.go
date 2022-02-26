package models

import "github.com/kamva/mgm/v3"

type TaraType int

const (
	High TaraType = iota
	Low
)

type Plank struct {
	mgm.DefaultModel `bson:",inline"`
	ID               int      `json:"id" bson:"id"`
	Height           int      `json:"height" bson:"height"`
	Width            int      `json:"width" bson:"width"`
	Length           int      `json:"length" bson:"length"`
	Amount           int      `json:"amount" bson:"amount"`
	Volume           float64  `json:"volume" bson:"volume"`
	Type             TaraType `json:"type" bson:"type"`
	Zkv              bool     `json:"zkv" bson:"zkv"`
	Kalts            bool     `json:"kalts" bson:"kalts"`
	D9               bool     `json:"d9" bson:"d9"`
}
