package main

import (
    "github.com/gin-gonic/gin"
    "github.com/PMaini532/Sprint_Planner/database"
    "github.com/PMaini532/Sprint_Planner/routes"
)

func main() {
    database.Connect()

    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Backend up and running!",
        })
    })
    routes.AuthRoutes(r)
    r.Run(":8080")
}
