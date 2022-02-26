package models

type TaraType int

const (
	High TaraType = iota
	Low
)

type Plank struct {
	ID     int
	Height int
	Width  int
	Length int
	Amount int
	Type   TaraType
	Zkv    bool
	Kalts  bool
	D9     bool
}
