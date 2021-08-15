package models

type Geographie struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"not null"`
}
