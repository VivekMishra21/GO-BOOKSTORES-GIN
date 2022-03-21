package routes

import (
	Controllers "book/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() gin.Engine {
	r := gin.Default()
	{
		r.GET("/books", Controllers.GetBooks)
		r.POST("/book", Controllers.CreateABook)
		r.GET("/book/:id", Controllers.GetABook)
		r.PUT("/book:id", Controllers.UpdateABook)
		r.DELETE("/book:id", Controllers.DeleteABook)
	}

	return *r
}
