package main

import "github.com/gin-gonic/gin"

func init() {
	utilities.EnableSQLDatabasesConfiguration()
}

func main() {
	router := gin.New()
	router.Use(gin.Logger())
}
