package main

import "time"

type CommonModelFields struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type Division struct {
	CommonModelFields `gorm:"embedded"`
	Name              string   `gorm:"not null;uniqueIndex" json:"name"`
	Drivers           []Driver `json:"drivers"`
	Events            []Event  `json:"event"`
}

type Driver struct {
	CommonModelFields `gorm:"embedded"`
	LastName          string   `gorm:"not null" json:"lastName"`
	FirstName         string   `gorm:"not null" json:"firstName"`
	CarNumber         string   `gorm:"not null;uniqueIndex:idx_driver" json:"carNumber"`
	DivisionID        uint     `json:"divisionId"`
	Division          Division `json:"division"`
}

type Event struct {
	CommonModelFields `gorm:"embedded"`
	Date              int64    `gorm:"not null" json:"date"`
	Name              string   `gorm:"not null;unique" json:"name"`
	Track             string   `gorm:"not null" json:"track"`
	DivisionID        uint     `json:"divisionId"`
	Division          Division `json:"division"`
	ModID             uint     `json:"modId"`
	Mod               Mod      `json:"mod"`
}

type EventResult struct {
	CommonModelFields `gorm:"embedded"`
	EventID           uint   `gorm:"not null" json:"eventId"`
	Event             Event  `json:"event"`
	DriverID          uint   `gorm:"not null" json:"driverId"`
	Driver            Driver `json:"driver"`
	Heat1             int64  `json:"heat1"`
	Heat2             int64  `json:"heat2"`
	Feature           int64  `json:"feature"`
}

type Mod struct {
	CommonModelFields `gorm:"embedded"`
	Name              string  `gorm:"not null;uniqueIndex" json:"name"`
	Events            []Event `json:"event"`
}

type DriverEventPoints struct {
	ID       uint   `json:"id"`
	DriverID uint   `json:"driverId"`
	Driver   Driver `json:"driver"`
	EventID  uint   `json:"eventId"`
	Event    Event  `json:"event"`
	Points   int64  `json:"points"`
}
