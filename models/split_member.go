package models

type SplitMember struct {
	ID      int          `json:"id" bson:"id"`
	Members []TeamMember `json:"members" bson:"members"`
}
