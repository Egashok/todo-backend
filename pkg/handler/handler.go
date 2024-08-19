package handler

import (
	"github/Egashok/todo-backend/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service

}
func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
 }

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth:=router.Group("/auth")
	{
		auth.POST("/sign-in",h.signIn)
		auth.POST("/sign-up",h.signUp)
	}
	api:=router.Group("/api")
	{
		api.GET("/",h.getAllLists)
		api.POST("/",h.createList)
		api.GET("/:id",h.getListById)
		api.PUT("/:id",h.updateList)
		api.DELETE("/:id",h.deleteList)
		items:=api.Group("/items")
		{
		items.GET("/",h.getAllItems)
		items.POST("/",h.createItem)
		items.GET("/:items_id",h.getItemById)
		items.PUT("/:items_id",h.updateItem)
		items.DELETE("/:items_id",h.deleteItem)
		}
	}
	return router
}