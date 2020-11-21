package bitemporal

import (
	"time"

	"github.com/guregu/null"
	"github.com/lovung/gotemporal"
	"gorm.io/gorm"
)

var _ gotemporal.Temporaler = (*BiGormer)(nil)

type BiGormer struct {
	db gorm.DB
}

func (g *BiGormer) Create(model interface{}) error {
	return g.db.Create(model).Error
}

func (g *BiGormer) UpdateObject(model interface{}) error {
	panic("not implemented") // TODO: Implement
}

func (g *BiGormer) UpdateRecord(model interface{}) error {
	panic("not implemented") // TODO: Implement
}

func (g *BiGormer) GetRecord(model gotemporal.TemporalModel, systemAt time.Time, validAt time.Time) error {
	return g.db.Scopes(
		gotemporal.WithTID(model),
		gotemporal.WithSysAt(systemAt),
		gotemporal.WithValidAt(validAt),
	).Take(model).Error
}

func (g *BiGormer) ListRecords(model gotemporal.TemporalModel, at time.Time, to interface{}) error {
	return g.db.Scopes(
		gotemporal.WithTID(model),
		gotemporal.WithSysAt(at),
	).Find(to).Error
}

func (g *BiGormer) DeleteObject(tid gotemporal.TIDer) error {
	return g.db.Delete("tid = ?", tid.GetTID()).Error
}

func (g *BiGormer) HardDeleteObject(tid gotemporal.TIDer) error {
	return g.db.Unscoped().Delete("tid = ?", tid.GetTID()).Error
}

func (g *BiGormer) DeleteRecordWithoutCollection(record gotemporal.TemporalModel, modifyTime null.Time) error {
	// TODO: Update here follow bitemporal
	return g.db.Delete("id = ?", record.GetID()).Error
}

func (g *BiGormer) DeleteRecord(record gotemporal.TemporalModel, modifyTime null.Time) error {
	collection, err := g.loadCollection(record)
	if err != nil {
		return err
	}

	collection.SortByValidTime()
	newRecords, err := collection.DeleteByID(record, modifyTime)
	if err != nil {
		return err
	}
	record.Clean()
	err = g.db.Model(record).Updates(collection).Error
	if err != nil {
		return err
	}
	err = g.db.Model(record).Create(newRecords).Error
	if err != nil {
		return err
	}
	return nil
}

func (g *BiGormer) loadCollection(tid gotemporal.TIDer) (Collection, error) {
	var list Collection
	err := g.db.Scopes(
		gotemporal.WithTID(tid),
	).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}
