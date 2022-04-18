# GoLang_API_Rest_study

Execução: go run main.go


Metodos criados encontram-se no arquivo main.go:

r.GET("/books", controllers.FindBooks)

r.POST("/books", controllers.CreateBook)

r.GET("/books/:id", controllers.FindBook)

r.PATCH("/books/:id", controllers.UpdateBook)

r.DELETE("/books/:id", controllers.DeleteBook)
