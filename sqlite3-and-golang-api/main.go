package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)
var db *gorm.DB

type Post struct {
	gorm.Model
	Title string `json:"title"`
	Slug  string `json:"slug" gorm:"uniqueIndex"`
	Likes int		`json:"likes"`
}

func (p *Post) String() string {
	return fmt.Sprintf("Post Title: %s, Slug: %s", p.Title, p.Slug)
}

func main() {
	db, err := gorm.Open(sqlite.Open("post.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	db.AutoMigrate(&Post{})
	log.Println("DB created...")
	r := gin.Default()
	r.POST("/post",func( c *gin.Context){
		var posts Post
		err := c.ShouldBindJSON(&posts)
		if err!=nil{
			log.Println("Error while posting::",err)
			c.JSON(http.StatusBadRequest,gin.H{"message":err})
		}

		_,err = createPost(db,posts.Title,posts.Slug)
		if err != nil {
			log.Fatal("Failed to create post: ", err)
		}
		c.JSON(http.StatusOK,gin.H{"message":"Successfully created a new post."})
	})

	

	r.GET("/data/:id",func(c *gin.Context){
		id := c.Param("id")
		retrievedPost, err := getPostBySlug(db, id)
		if err != nil {
			log.Fatal("Failed to retrieve post: ", err)
		}
		c.JSON(http.StatusOK,gin.H{"message":retrievedPost})
		fmt.Println(retrievedPost)
	})

	log.Println("Listening in port 4000...")
	err = http.ListenAndServe(":4000",r)
	if err!=nil{
		log.Println("Error while starting the server!")
	}

}

func createPost(db *gorm.DB, title, slug string) (*Post, error) {
	post := &Post{Title: title, Slug: slug}
	result := db.Create(post)
	if result.Error != nil {
		return nil, result.Error
	}
	return post, nil
}

func getPostBySlug(db *gorm.DB, slug string) (*Post, error) {
	var post Post
	result := db.First(&post, "slug = ?", slug)
	if result.Error != nil {
		return nil, result.Error
	}
	return &post, nil
}
