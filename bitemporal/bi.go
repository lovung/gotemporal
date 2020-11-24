package bitemporal

import (
	"time"

	"github.com/guregu/null"
)

// Model is the temporal model with 2-axes of time
type Model struct {
	ID        uint64
	TID       string `gorm:"column:tid"`
	SysFrom   time.Time
	SysTo     null.Time
	ValidFrom time.Time
	ValidTo   null.Time
}

// Clean the information related with history
func (t *Model) Clean() {
	t.ID = 0
}
