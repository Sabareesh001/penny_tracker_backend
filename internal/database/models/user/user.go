package user

type User struct {
	Id           int    `gorm:"primaryKey"`
	Username     string
	Password     string
	First_Name   string
	Last_Name    string
	Age          int
	mail         string
	phone        string

	Occupation int
	OccupationKey   Occupation `gorm:"foreignKey:OccupationKey"`

	Gender int
	GenderKey   Gender `gorm:"foreignKey:GenderKey"`

	Country int
	CountryKey   Country `gorm:"foreignKey:CountryKey"`
}


type Country struct {
	Id   int    `gorm:"primaryKey"`
	Name string
	Code string
}

type Gender struct {
	Id   int    `gorm:"primaryKey"`
	Name string
}

type Occupation struct {
	Id   int    `gorm:"primaryKey"`
	Name string
}
