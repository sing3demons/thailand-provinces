package entity

type ZipCode struct {
	ID           uint   `gorm:"primarykey"`
	LocationCode uint   `gorm:"not null"`
	Province     string `gorm:"not null"`
	Amphoe       string `gorm:"not null"`
	Subdistrict  string `gorm:"not null"`
	ZipCode      uint   `gorm:"not null"`
	Node         string
}
