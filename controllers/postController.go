package controllers

import (
	"test/go-basic-project/initializers"
	"test/go-basic-project/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	var body struct {
		Title string
		Body string 
	}
	c.Bind(&body)

	post := models.Post{
		Title: body.Title, 
		Body: body.Body,
	}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {

	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c* gin.Context) {
	
	id := c.Param("id")
 
	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c* gin.Context) {
	id := c.Param("id")

	var body struct {
		Title string
		Body string
	}

	c.Bind(&body)

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c* gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.Status(200)
}
