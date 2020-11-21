package gotemporal

import (
	"time"

	"github.com/guregu/null"
)

type creator interface {
	Create(model interface{}) error
}

type updater interface {
	UpdateObject(model interface{}) error
	UpdateRecord(model interface{}) error
}

type getter interface {
	GetRecord(model TemporalModel, systemAt time.Time, validAt time.Time) error
	ListRecords(model TemporalModel, at time.Time, to interface{}) error
}

type deleter interface {
	DeleteObject(tid TIDer) error
	HardDeleteObject(tid TIDer) error
	DeleteRecord(record TemporalModel, modifyValid null.Time) error
	DeleteRecordWithoutCollection(record TemporalModel, modifyTime null.Time) error
}

type Temporaler interface {
	creator
	updater
	getter
	deleter
}
