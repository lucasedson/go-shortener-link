package schemas

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	Url       string
	ShortLink string `gorm:"unique"`
}
