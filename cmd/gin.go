package main
import (
    "os"
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context){
        c.JSON(200, gin.H{
            "message":"pong",
        })
    })
    r.GET("/hello", func(c *gin.Context){
        c.JSON(http.StatusOK, gin.H{
            "message":"hello Gin Framework",
        })
    })
    r.Run(port())
}

func port() string {
    port := os.Getenv("PORT")
    if len(port) == 0 {
        port = "8080"
    }
    return ":" + port
}
