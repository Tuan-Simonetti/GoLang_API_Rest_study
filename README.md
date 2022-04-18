# GoLang_API_Rest_study

Execução: go run main.go


Metodos criados encontram-se no arquivo main.go:

GET "/books"

POST "/books"   deve ser passado o body de modelo abaixo:      

{
  "title": "The Infinite Game"
  "author": "J. K. Rowling"
}



GET "/books/:id"            exemplo: localhost:8080/books/2

PATCH "/books/:id"          exemplo: localhost:8080/books/2

DELETE "/books/:id"         exemplo: localhost:8080/books/2
