package models

type Amphue struct {
	ID         uint     `gorm:"primarykey"`
	Code       string   `gorm:"not null"`
	NameTH     string   `gorm:"not null"`
	NameEN     string   `gorm:"not null"`
	ProvinceID uint     `gorm:"not null"`
	Province   Province `gorm:"foreignKey:ProvinceID"`
}
