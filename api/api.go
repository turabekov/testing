package api

import (
	_ "app/api/docs"
	"app/api/handler"

	"app/config"
	"app/pkg/logger"
	"app/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func NewApi(r *gin.Engine, cfg *config.Config, store storage.StorageI, logger logger.LoggerI) {
	handler := handler.NewHandler(cfg, store, logger)

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

	// user api
	r.POST("/user", handler.CreateUser)
	r.GET("/user/:id", handler.GetByIdUser)
	r.GET("/user", handler.GetListUser)
	r.PUT("/user/:id", handler.UpdateUser)
	r.DELETE("/user/:id", handler.DeleteUser)

	// category api
	r.POST("/category", handler.CreateCategory)
	r.GET("/category/:id", handler.GetByIdCategory)
	r.GET("/category", handler.GetListCategory)
	r.PUT("/category/:id", handler.UpdateCategory)
	r.DELETE("/category/:id", handler.DeleteCategory)

	// product api
	r.POST("/product", handler.CreateProduct)
	r.GET("/product/:id", handler.GetByIdProduct)
	r.GET("/product", handler.GetListProduct)
	r.PUT("/product/:id", handler.UpdateProduct)
	r.DELETE("/product/:id", handler.DeleteProduct)

	// client api
	r.POST("/client", handler.CreateClient)
	r.GET("/client/:id", handler.GetByIdClient)
	r.GET("/client", handler.GetListClient)
	r.PUT("/client/:id", handler.UpdateClient)
	r.DELETE("/client/:id", handler.DeleteClient)

	// order api
	r.POST("/order", handler.CreateOrder)
	r.GET("/order/:id", handler.GetByIdOrder)
	r.GET("/order", handler.GetListOrder)
	r.PUT("/order/:id", handler.UpdateOrder)
	r.DELETE("/order/:id", handler.DeleteOrder)
	r.POST("/order_item/", handler.CreateOrderItem)
	r.DELETE("/order_item/:id", handler.DeleteOrderItem)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
