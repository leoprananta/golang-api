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