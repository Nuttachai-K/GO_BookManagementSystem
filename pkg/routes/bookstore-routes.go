package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/korn/learning/book_store/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *gin.Engine) {
	router.POST("/book/", controllers.CreateBook)
	router.GET("/book", controllers.GetBook)
	router.GET("/book/:bookId", controllers.GetBookById)
	router.PUT("/book/:bookId", controllers.UpdateBook)
	router.DELETE("/book/:bookId", controllers.DeleteBook)
}
