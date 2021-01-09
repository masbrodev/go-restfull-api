package structs

import "github.com/jinzhu/gorm"

type Person struct {
	gorm.Model
	Nama    string
	Jurusan string
}
