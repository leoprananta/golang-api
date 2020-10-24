package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"leonanta.github.io/models"
	"leonanta.github.io/controllers"
)

func main(){
	r := gin.Default();
	
	db := models.SetupModels() 
	
	r.Use(func (c *gin.Context){
		c.Set("db", db)
		c.Next()
	})

	r.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"data":"Succesful"})
	})

	r.GET("/mahasiswa", controllers.GetAllMahasiswa)
	r.POST("/mahasiswa", controllers.CreateMahasiswa)
	r.PUT("/mahasiswa/:nim", controllers.UpdateMahasiswa)
	r.DELETE("/mahasiswa/:nim", controllers.DeleteMahasiswa)

	r.Run()
}