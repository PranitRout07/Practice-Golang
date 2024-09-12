package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string
	Slug  string `gorm:"uniqueIndex"`
	Likes int
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

	post, err := createPost(db, "The Short Story", "slug")
	if err != nil {
		log.Fatal("Failed to create post: ", err)
	}
	fmt.Println(post)

	retrievedPost, err := getPostBySlug(db, "new-slug")
	if err != nil {
		log.Fatal("Failed to retrieve post: ", err)
	}
	fmt.Println(retrievedPost)
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
