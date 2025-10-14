package payments

import (
	"time"

	"github.com/Sabareesh001/penny_tracker_backend/internal/database/models/user"
)

type Payments struct {
	Id int
	User    int
	UserKey user.User `gorm:"foreignKey:UserKey"`
	Category int
	CategoryKey   Payment_Categories `gorm:"foreignKey:CategoryKey"`
	Label string
	Amount float64
	Datetime time.Time
	Currency string
}