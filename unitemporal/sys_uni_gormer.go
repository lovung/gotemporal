package gotemporal

import (
	"gorm.io/gorm"
)

// SysUniGormer implemented the uni-temporal interface
type SysUniGormer struct {
	db gorm.DB
}

// Create a new object
