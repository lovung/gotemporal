package entity

import "github.com/lovung/gotemporal/bitemporal"

type Example struct {
	bitemporal.Model
	Name string
}

func (e *Example) GetTID() string {
	return e.TID
}

func (e *Example) SetTID(s string) {
	e.TID = s
}

func (e *Example) GetID() uint64 {
	return e.ID
}

func (e *Example) Clean() {
	e.ID = 0
}
