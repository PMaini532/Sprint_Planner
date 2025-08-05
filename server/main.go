package main

import (
    "github.com/gin-gonic/gin"
    "github.com/PMaini532/Sprint_Planner/database"
    "github.com/PMaini532/Sprint_Planner/routes"
    "github.com/gin-contrib/cors"
)

func main() {
    database.Connect()

    r := gin.Default()

    r.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:3000"},
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
    AllowCredentials: true,
    }))

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Backend up and running!",
        })
    })
    routes.AuthRoutes(r)
    routes.UserRoutes(r)
    r.Run(":8080")
}
