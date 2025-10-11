package user

type User struct {
	Id           int    `gorm:"primaryKey"`
	Username     string
	Password     string
	First_Name   string
	Last_Name    string
	Age          int
	Email         string
	Phone        string

	IsEmailVerified string `gorm:"column:isEmailVerified"`
    IsPhoneVerified string `gorm:"column:isPhoneVerified"`

	Occupation int
	OccupationKey   Occupation `gorm:"foreignKey:OccupationKey"`

	Gender int
	GenderKey   Gender `gorm:"foreignKey:GenderKey"`

	Monthly_Income float32
	Monthly_Income_Currency string
	Country string
}



type Gender struct {
	Id   int    `gorm:"primaryKey"`
	Name string
	Symbol string
}

type Occupation struct {
	Id   int    `gorm:"primaryKey"`
	Name string
	Symbol string
}
