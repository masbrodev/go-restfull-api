package controllers

import (
	"net/http"

	"../structs"
	"github.com/gin-gonic/gin"
)

// GetPerson one (id)
func (idb *InDB) GetPerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&person).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"Conunt": 1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// GetPersons All
func (idb *InDB) GetPersons(c *gin.Context) {
	var (
		person []structs.Person
		result gin.H
	)

	idb.DB.Find(&person)
	if len(person) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  len(person),
		}
	}

	c.JSON(http.StatusOK, result)
}

// CreatePerson (tambah)
func (idb *InDB) CreatePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	nama := c.PostForm("nama")
	jurusan := c.PostForm("jurusan")

	person.Nama = nama
	person.Jurusan = jurusan

	idb.DB.Create(&person)
	result = gin.H{
		"result": person,
	}
	c.JSON(http.StatusOK, result)

}

// UpdatePerson (id)
func (idb *InDB) UpdatePerson(c *gin.Context) {
	id := c.Query("id")
	nama := c.PostForm("nama")
	jurusan := c.PostForm("jurusan")

	var (
		person    structs.Person
		newPerson structs.Person
		result    gin.H
	)

	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "data tidak ditemukan",
		}
	}
	newPerson.Nama = nama
	newPerson.Jurusan = jurusan
	err = idb.DB.Model(&person).Update(newPerson).Error
	if err != nil {
		result = gin.H{
			"result": "Update Gagal",
		}
	} else {
		result = gin.H{
			"result": "Data Berhasil diUpdate",
		}
	}
	c.JSON(http.StatusOK, result)
}

// DeletePerson (Hapus)
func (idb *InDB) DeletePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "Data tidak ditemukan",
		}
		err = idb.DB.Delete(&person).Error
		if err != nil {
			result = gin.H{
				"result": "Gagal dihapus",
			}
		} else {
			result = gin.H{
				"result": "Berhasil dihapus",
			}
		}
		c.JSON(http.StatusOK, result)

	}
}
