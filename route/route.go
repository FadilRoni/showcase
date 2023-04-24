package route

import (
	"challange-8_2/handler"
	"challange-8_2/service"

	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, app service.ServiceInterface) {
	server := handler.NewHttpServer(app)
	api := r.Group("/books")
	{
		api.POST("", server.AddBook)
		api.GET("", server.GetAllBook)
		api.GET("/:id", server.GetBookId)
		api.PUT("/:id", server.UpdateBook)
		api.DELETE("/:id", server.DeleteBook)
	}
}
