# Go-Bookstore

This is my CRUD bookstore managment api

## Endpoints
# Get all books
GET "http://localhost:8080/book/" 

# Get book by id
GET "http://localhost:8080/book/{id}" 

# Create book
POST "http://localhost:8080/book/" 
Body : {
 {
    "Name": string,
    "Author": string,
    "Publication": string
}
}

# Delete book
DELETE "http://localhost:8080/book/{id}" 

# Update book 
PUT "http://localhost:8080/book/{id}" 
Body : {
 {
    "Name": string,
    "Author": string,
    "Publication": string
}
}
