package models

type Province struct {
	ID           uint   `gorm:"primarykey"`
	Code         string `gorm:"not null"`
	NameTH       string `gorm:"not null"`
	NameEN       string `gorm:"not null"`
	GeographieID uint   `gorm:"not null"`
	Geographie   Geographie
}
