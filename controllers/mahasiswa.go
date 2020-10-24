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

	//input validation
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

//update mahasiswa
func UpdateMahasiswa(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)

	// err != nil = error tidak sama dengan nil atau ada error

	// check exsisting data
	var mhs models.Mahasiswa
	if err := db.Where("nim = ?", c.Param("nim")).First(&mhs).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"Error":"Data not found"})
		return
	}

	//input validation
	var dataInput MahasiswaInput
	if err := c.ShouldBindJSON(&dataInput); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"Error":err.Error()})
		return
	}

	//update process
	db.Model(&mhs).Update(dataInput)

	c.JSON(http.StatusOK, gin.H{"data":mhs})
}

//delete mahasiswa
func DeleteMahasiswa(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)

	// err != nil = error tidak sama dengan nil atau ada error

	// check exsisting data
	var mhs models.Mahasiswa
	if err := db.Where("nim = ?", c.Param("nim")).First(&mhs).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"Error":"Data not found"})
		return
	}

	//delete process
	db.Delete(&mhs)

	c.JSON(http.StatusOK, gin.H{"data":true})
}