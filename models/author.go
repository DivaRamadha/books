package models

type Author struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:255"`
	Birth string `gorm:"localDate"`
	Books []Book `gorm:"foreignKey:AuthorId;constraint:OnDelete:CASCADE;"`
}
