package bitemporal

import (
	"time"

	"github.com/guregu/null"
)

// Model is the temporal model with 2-axes of time
type Model struct {
	ID        interface{}
	TID       interface{}
	SysFrom   time.Time
	SysTo     null.Time
	ValidFrom time.Time
	ValidTo   null.Time
}

// Clean the information related with history
func (t *Model) Clean() {
	t.ID = nil
}
