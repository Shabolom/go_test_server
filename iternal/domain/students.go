package domain

type Student struct {
	Base
	Name    string `gorm:"type:text"`
	Surname string `gorm:"type:text"`
	Gender  string `gorm:"type:text"`
	Age     int    `gorm:"type:int"`
	Country string `gorm:"type:text"`
	Email   string `gorm:"type:text"`
}
