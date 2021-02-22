package main

type Book struct {
    Title string `json:"title"`
    Author string `json:"author"`
    ISBN string `json:"isbn"`
    Description string `json:"description"`
}

var books = []stringBook{
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

func GetBook(isbn string) (Book, book) {
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
