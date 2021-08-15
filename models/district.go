package models

type District struct {
	ID        uint   `gorm:"primarykey"`
	ZipCode   uint   `gorm:"not null"`
	NameTH    string `gorm:"not null"`
	NameEN    string `gorm:"not null"`
	AmphureID uint   `gorm:"not null"`
	Amphue    Amphue `gorm:"foreignKey:AmphureID"`
}
