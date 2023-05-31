package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // ルートパスのハンドラ
    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello, World!",
        })
    })

    // /users パスのハンドラ
    router.GET("/users", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "User list",
        })
    })

    router.Run(":8080")
}
