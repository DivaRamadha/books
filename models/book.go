package models

type Book struct {
	ID       uint   `gorm:"primaryKey"`
	Title    string `gorm:"size:255"`
	Isbn     string `gorm:"size:13;unique"`
	AuthorId uint
	Author   Author `gorm:"foreignKey:AuthorId"`
}
