package metals

import "github.com/Sabareesh001/penny_tracker_backend/internal/database/models/user"

type TrackingStatus string

const (
    StatusDisabled  TrackingStatus = "0"
    StatusEnabled TrackingStatus = "1"
)


type Metals struct {
	Id     int `gorm:"primaryKey"`
	Name   string
	Image    string
	Symbol string
	Color string
}

type UserMetalTracking struct {
	Id     int `gorm:"primaryKey"`
	Weight float64

	User    int
	UserKey user.User `gorm:"foreignKey:UserKey"`

	Metal    int
	MetalKey Metals `gorm:"foreignKey:MetalKey"`

	Status TrackingStatus
}