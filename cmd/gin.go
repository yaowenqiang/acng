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
    r.GET("/api/books", func(c *gin.Context){
        c.JSON(http.StatusOK, AllBooks())
    })

    r.POST("/api/books", func(c *gin.Context){
        var book Book
        if c.BindJSON(&book) == nil {
            isbn, created := CreateBook(book)
            if created {
                c.Header("location", "/api/books/" + isbn)
                c.Status(http.StatusCreated)
            } else {
                c.Status(http.StatusConflict)
            }
        }
    })

    r.GET("/api/books/:isbn", func(c *gin.Context){
        isbn := c.Params.ByName("isbn")
        book, found := GetBook(isbn)

        if found {
            c.JSON(http.StatusOK, book)
        } else {
            c.AbortWithStatus(http.StatusNotFound)
        }
    })

    r.PUT("/api/books/:isbn", func(c *gin.Context){
        isbn := c.Params.ByName("isbn")
        var book Book
        if c.BindJSON(&book) != nil {
            exists := UpdateBook(isbn, book)
            if exists {
                c.Status(http.StatusOK)
            } else {
                c.Status(http.StatusNotFound)
            }
        }
    })
    r.DELETE("/api/books/:isbn", func(c *gin.Context){
        isbn := c.Params.ByName("isbn")
        DeleteBook(isbn)
        c.Status(http.StatusOK)
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
type Book struct {
    Title string `json:"title"`
    Author string `json:"author"`
    ISBN string `json:"isbn"`
    Description string `json:"description"`
}

var books = map[string]Book{
    "0000000000001":Book{Title:"Glaxy",Author:"glaxy", ISBN:"abc", Description: "desc" },
    "0000000000002":Book{Title:"University",Author:"University", ISBN:"def", Description: "asc" },
}

func AllBooks() []Book {
    values := make([]Book, len(books))
    idx := 0
    for _, book := range books {
        values[idx] = book
        idx++
    }
    return values
}

func GetBook(isbn string) (Book, bool) {
    book, found := books[isbn]
    return book, found
}

func CreateBook(book Book) (string, bool) {
    _, exists := books[book.ISBN]
    if exists {
        return "", false
    }

    books[book.ISBN] = book
    return book.ISBN, true
}

func UpdateBook(isbn string, book Book) bool {
    _, exists := books[isbn]
    if exists {
        books[isbn] = book
    }
    return exists
}

func DeleteBook(isbn string) {
    delete(books, isbn)
}
