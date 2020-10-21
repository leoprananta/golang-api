package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"leonanta.github.io/models"
)

type MahasiswaInput struct{
	Nim string `json:"nim"`
	Nama string `json:"nama"`
}

//get all mahasiswa
func GetAllMahasiswa(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)

	var mhs []models.Mahasiswa
	db.Find(&mhs)
	c.JSON(http.StatusOK, gin.H{"data":mhs})
}

//create mahasiswa
func CreateMahasiswa(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)

	//validation input
	var dataInput MahasiswaInput
	if err := c.ShouldBindJSON(&dataInput); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"Error":err.Error()})
		return
	}

	//input process
	mhs := models.Mahasiswa{
		NIM : dataInput.Nim,
		Nama : dataInput.Nama,
	}

	db.Create(&mhs)

	c.JSON(http.StatusOK, gin.H{"data":mhs})
}