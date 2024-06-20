package main

import (
	"github.com/gin-gonic/gin"
	"rest-api/controllers/infopelanggancontroller"
	"rest-api/models"
)

func main() {
	route := gin.Default()

	models.ConnectDatabase()

	route.GET("/api/infopelanggan", infopelanggancontroller.Index)         //Index
	route.GET("/api/infopelanggan/:id", infopelanggancontroller.Show)      //Show
	route.POST("/api/infopelanggan", infopelanggancontroller.Create)       //Create
	route.PUT("/api/infopelanggan/:id", infopelanggancontroller.Update)    //Update
	route.DELETE("/api/infopelanggan/:id", infopelanggancontroller.Delete) //Delete

	route.Run()
}
