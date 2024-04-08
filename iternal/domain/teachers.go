package domain

type Teachers struct {
	Base
	Name      string `gorm:"type:text"`
	Surname   string `gorm:"type:text"`
	MidleName string `gorm:"type:text"`
}
