package gotemporal

import (
	"time"

	"github.com/guregu/null"
)

type ValidUnier interface {
	Base() BaseValidUni
}

// BaseValidUni is the temporal model with 1-axis (real) of time
type BaseValidUni struct {
	ID        interface{}
	TID       interface{}
	ValidFrom time.Time
	ValidTo   null.Time
}

// Clean the information related with history
func (t *BaseValidUni) Clean() {
	t.ID = nil
}

type SysUnier interface {
	Base() BaseSysUni
}

// BaseSysUni is the temporal model with 1-axis (system) of time
type BaseSysUni struct {
	ID      interface{}
	TID     interface{}
	SysFrom time.Time
	SysTo   null.Time
}

// Clean the information related with history
func (t *BaseSysUni) Clean() {
	t.ID = nil
}
