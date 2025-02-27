package models

import (
	"time"
)

type Category string

const (
	Electronics   Category = "Electronics"
	Books         Category = "Books"
	Clothing      Category = "Clothing"
	Furniture     Category = "Furniture"
	Tickets       Category = "Tickets"
	Sports        Category = "Sports"
	Appliances    Category = "Appliances"
	Miscellaneous Category = "Miscellaneous"
)

type Product struct {
	ID              int       `gorm:"primaryKey;autoIncrement"`
	Pid             string    `gorm:"type:uuid;unique;not null"`
	Name            string    `gorm:"not null"`
	Description     string    `gorm:"type:text"`
	Price           float64   `gorm:"not null"`
	Category        Category  `gorm:"type:text;not null"`
	Quantity        int       `gorm:"not null;default:0"`
	PopularityScore float64   `gorm:"default:0"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`

	Images []ProductImage `gorm:"foreignKey:Pid;references:Pid"`
}

type ProductImage struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Pid       string    `gorm:"type:varchar(36);not null;index"`
	MimeType  string    `gorm:"type:varchar(100);not null"`
	Url 	  string    `gorm:"type:text;not null;default:'https://placehold.co/600x400'"`
	IsMain    bool      `gorm:"not null;default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
