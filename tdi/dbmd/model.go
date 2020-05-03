package main

import "time"

// Table0 ...
type Table0 struct {
	ID        uint64 `gorm:"primary_key"`
	Texto     string
	Fechahora time.Time
	Fecha     time.Time
	Hora      time.Time
}

// Table1 ...
type Table1 struct {
	ID        uint64 `gorm:"primary_key"`
	Table0Id  *uint64
	Table0    *Table0 `gorm:"foreignkey:Table0Id"`
	Texto     *string
	Tsmallint *int16
	Tinteger  *int32
	Tbiging   *int64
	Tdecimal  *float64
	Treal     *float32
	Tdouble   *float64
}

// Table2 ...
type Table2 struct {
	ID       uint64 `gorm:"primary_key"`
	Table0Id *uint64
	Table0   *Table0 `gorm:"foreignkey:Table0Id"`
	Table1Id *uint64
	Table1   *Table1 `gorm:"foreignkey:Table1Id"`
	Texto    *string
}
