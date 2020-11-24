package main

import (
	"os"
	"time"

	"github.com/guregu/null"
	"github.com/lovung/gotemporal"
	"github.com/lovung/gotemporal/bitemporal"
	"github.com/lovung/gotemporal/example/entity"
	"github.com/lovung/gotemporal/example/pkg/gormutil"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var manager gotemporal.Temporal

func main() {
	logMode := logger.Default.LogMode(logger.Error)
	connStr := os.Getenv("DB_CONNECTION")
	_, err := gormutil.OpenDBConnection(
		connStr,
		gorm.Config{
			Logger: logMode,
		},
	)
	if err != nil {
		panic(err)
	}

	manager = bitemporal.NewManager(gormutil.GetDB(), "examples")
	ent, err := CreateNew()
	if err != nil {
		panic(err)
	}

	DeleteRecord(ent.ID, ent.TID)
}

func CreateNew() (entity.Example, error) {
	newExample := entity.Example{
		Model: bitemporal.Model{
			SysFrom:   time.Now(),
			ValidFrom: time.Date(2020, 10, 01, 0, 0, 0, 0, time.UTC),
		},
		Name: "Name",
	}
	err := manager.Create(&newExample)
	return newExample, err
}

func DeleteObject(tid string) error {
	wantDlt := entity.Example{
		Model: bitemporal.Model{
			TID: tid,
		},
	}
	err := manager.DeleteObject(&wantDlt)
	return err
}

func DeleteRecord(id uint64, tid string) error {
	wantDlt := entity.Example{
		Model: bitemporal.Model{
			ID:  id,
			TID: tid,
		},
	}
	err := manager.DeleteRecord(&wantDlt, null.TimeFrom(time.Now()))
	return err

}
