package controllers

import (
	"net/http"

	"github.com/PranitRout07/Practice-Golang/CRUD-using-gin-package/services"
	"github.com/gin-gonic/gin"
)
type Api struct{
	svc_data services.Svc
}


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
		"message": t.svc_data.GetDataService(t.svc_data),
	})
}

func (t *Api) getParams(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"user_id": id,
		"message": &t.svc_data,
	})
}

func (t *Api)postData(c *gin.Context) {
	err := c.BindJSON(&t.svc_data)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})

	}
	c.JSON(http.StatusOK, gin.H{
		"message": t.svc_data.PostDataService(t.svc_data),
	})
}

func (t *Api) putData(c *gin.Context) {
	err := c.BindJSON(&t.svc_data)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})

	}
	c.JSON(http.StatusOK, gin.H{
		"message": t.svc_data,
	})
}

func (t *Api) deleteData(c *gin.Context) {
	t.svc_data.DeleteDataService(t.svc_data)
	c.JSON(http.StatusOK, gin.H{
		"message": t.svc_data,
	})
}
