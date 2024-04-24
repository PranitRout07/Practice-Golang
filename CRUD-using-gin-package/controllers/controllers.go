package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
type Api struct{
	Name string `json:"name"`
	Email string `json:"email" binding:"required"`
}

var data Api 

func (t *Api)Initialize(r *gin.Engine){
	r.GET("/get", t.getData)
	r.POST("/post", t.postData)
	r.PUT("/put", t.putData)
	r.DELETE("/delete", t.deleteData)
	r.GET("/getparams/:id/",t.getParams)
	r.Run(":9090") 
}




func (t *Api) getData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

func (t *Api) getParams(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"user_id": id,
		"message": data,
	})
}

func (t *Api)postData(c *gin.Context) {
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

func (t *Api) putData(c *gin.Context) {
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

func (t *Api) deleteData(c *gin.Context) {
	data = Api{}
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}
