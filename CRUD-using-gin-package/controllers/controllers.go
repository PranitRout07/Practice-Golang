package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
type api struct{
	Name string `json:"name"`
	Email string `json:"email" binding:"required"`
}

var data api 

func Initialize(r *gin.Engine){
	r.GET("/get", getData)
	r.POST("/post", postData)
	r.PUT("/put", putData)
	r.DELETE("/delete", deleteData)
	r.GET("/getparams/:id/",getParams)
	r.Run(":9090") 
}




func getData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

func getParams(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"user_id": id,
		"message": data,
	})
}

func postData(c *gin.Context) {
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})

	}
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

func putData(c *gin.Context) {
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})

	}
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

func deleteData(c *gin.Context) {
	data = api{}
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}