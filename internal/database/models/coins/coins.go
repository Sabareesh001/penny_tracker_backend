package coins

import "github.com/Sabareesh001/penny_tracker_backend/internal/database/models/user"


type TrackingStatus string

const (
    StatusDisabled  TrackingStatus = "0"
    StatusEnabled TrackingStatus = "1"
)


type Coins struct {
	Id     int `gorm:"primaryKey"`
	Name   string
	Image    string
	Symbol string
}

type UserCoinTracking struct {
	Id     int `gorm:"primaryKey"`
	Quantity float64

	User    int
	UserKey user.User `gorm:"foreignKey:UserKey"`

	Coin    int
	CoinKey Coins `gorm:"foreignKey:CoinKey"`

	Status TrackingStatus
}