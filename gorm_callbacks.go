package gotemporal

import (
	"time"

	"gorm.io/gorm"
)

// WithID where id = <input value>
func WithID(id IDer) func(d *gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		return d.Where("id = (?)", id.GetID())
	}
}

// WithTID where biid = <input value>
func WithTID(tid TIDer) func(d *gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		return d.Where("tid = (?)", tid.GetTID())
	}
}

// WithValidAt where code = <input value>
func WithValidAt(validAt time.Time) func(d *gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		if validAt.IsZero() {
			return d
		}
		return d.Where("valid_from <= ? AND (valid_to IS NULL OR valid_to > ?)", validAt, validAt)
	}
}

// WithSysAt where code = <input value>
func WithSysAt(sysAt time.Time) func(d *gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		if sysAt.IsZero() {
			return d
		}
		return d.Where("created_at <= ? AND (deleted_at IS NULL OR deleted_at > ?)", sysAt, sysAt)
	}
}

// ValidToIsNull where the valid_to column is NULL
func ValidToIsNull(table string) func(d *gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		return d.Where(table + ".valid_to IS NULL")
	}
}

// SysToIsNull where code = <input value>
func SysToIsNull(table string) func(d *gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		return d.Where(table + ".deleted_at IS NULL")
	}
}
