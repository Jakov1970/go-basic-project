package main

import (
	"test/go-basic-project/initializers"
	"test/go-basic-project/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
