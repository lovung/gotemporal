package bitemporal

import (
	"time"

	"github.com/guregu/null"
	"github.com/lovung/gotemporal"
	"gorm.io/gorm"
)

var _ gotemporal.Temporal = (*BiGormer)(nil)

type BiGormer struct {
	db        *gorm.DB
	tableName string
}

func NewManager(
	db *gorm.DB,
	tableName string,
) *BiGormer {
	return &BiGormer{
		db:        db,
		tableName: tableName,
	}
}

func (g *BiGormer) Create(model gotemporal.TIDer) error {
	model.SetTID(gotemporal.GenTID())
	return g.db.Table(g.tableName).Create(model).Error
}

func (g *BiGormer) UpdateObject(model interface{}) error {
	panic("not implemented") // TODO: Implement
}

func (g *BiGormer) UpdateRecord(model interface{}) error {
	panic("not implemented") // TODO: Implement
}

func (g *BiGormer) GetRecord(model gotemporal.TemporalModel, systemAt time.Time, validAt time.Time) error {
	return g.db.Table(g.tableName).Scopes(
		gotemporal.WithTID(model),
		gotemporal.WithSysAt(systemAt),
		gotemporal.WithValidAt(validAt),
	).Take(model).Error
}

func (g *BiGormer) ListRecords(model gotemporal.TemporalModel, at time.Time, to interface{}) error {
	return g.db.Table(g.tableName).Scopes(
		gotemporal.WithTID(model),
		gotemporal.WithSysAt(at),
	).Find(to).Error
}

func (g *BiGormer) DeleteObject(object gotemporal.TIDer) error {
	return g.db.Table(g.tableName).Where("tid = ?", object.GetTID()).Update("sys_to", time.Now()).Error
}

func (g *BiGormer) HardDeleteObject(tid gotemporal.TIDer) error {
	return g.db.Table(g.tableName).Where("tid = ?", tid.GetTID()).Delete(tid).Error
}

func (g *BiGormer) DeleteRecordWithoutCollection(record gotemporal.TemporalModel, modifyTime null.Time) error {
	// TODO:
	// Find the previous/next record to update the valid_from and valid_to
	// Set the sys_to for target record
	return g.db.Table(g.tableName).Where("id = ?", record.GetID()).Update("sys_to", time.Now()).Error
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
	for _, e := range collection.ToSlice() {
		err = g.db.Table(g.tableName).Select("sys_from", "sys_to", "valid_from", "valid_to").Updates(e).Error
	}
	if err != nil {
		return err
	}
	for _, e := range newRecords.ToSlice() {
		// FIXME: Create the whole model instead of only bitemporal.Model
		err = g.db.Table(g.tableName).Create(e).Error
	}
	if err != nil {
		return err
	}
	return nil
}

func (g *BiGormer) loadCollection(tid gotemporal.TIDer) (Collection, error) {
	var list Collection
	err := g.db.Table(g.tableName).Scopes(
		gotemporal.WithTID(tid),
	).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}
