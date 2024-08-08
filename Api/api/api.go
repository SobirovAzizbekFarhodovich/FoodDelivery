package api

import (
	handlerC "api/api/handler"

	_ "api/docs"
	_ "api/genprotos/delivery"
	_ "api/genprotos/user"

	"api/api/middlerware"
	"log"

	"github.com/casbin/casbin/v2"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

// @title Delivery System API
// @version 1.0
// @description API for Delivery System resources
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin( /*AuthConn, */ companyConn *grpc.ClientConn) *gin.Engine {
	delivery := handlerC.NewHandler(companyConn)

	router := gin.Default()

	enforcer, err := casbin.NewEnforcer("/home/sobirov/go/src/gitlab.com/FoodDelivery/Api/api/model.conf", "/home/sobirov/go/src/gitlab.com/FoodDelivery/Api/api/policy.csv")
	if err != nil {
		log.Fatal(err)
	}

	sw := router.Group("/")
	sw.Use(middlerware.NewAuth(enforcer))

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Adjust for your specific origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	cart := router.Group("/cart")
	{
		cart.POST("/create",delivery.CreateCart)
		cart.DELETE("/delete/:id",delivery.DeleteCart)
		cart.PUT("/update",delivery.UpdateCart)
		cart.GET("/get/:id",delivery.GetByIdCart)
		cart.GET("/get",delivery.GetAllCart)
	}

	orderitems := router.Group("/order_items")
	{
		orderitems.POST("/create",delivery.CreateOrderItems)
		orderitems.PUT("/update",delivery.UpdateOrderItems)
		orderitems.DELETE("/delete/:id", delivery.DeleteOrderItems)
		orderitems.GET("/get/:id",delivery.GetByIdOrderItems)
		orderitems.GET("/get",delivery.GetAllOrderItems)
	}
	order := router.Group("/order")
	{
		order.POST("/create",delivery.CreateOrder)
		order.PUT("/update",delivery.UpdateOrder)
		order.DELETE("/delete/:id", delivery.DeleteOrder)
		order.GET("/get/:id",delivery.GetByIdOrder)
		order.GET("/get",delivery.GetAllOrders)
		order.GET("/:location",delivery.GetByLocation)
		order.PUT("/status",delivery.UpdateStatus)
		order.GET("/history",delivery.GetHistory)
	}

	product := router.Group("/product")
	{
		product.POST("/create",delivery.CreateProduct)
		product.PUT("/update",delivery.UpdateProduct)
		product.DELETE("/delete/:id", delivery.DeleteProduct)
		product.GET("/get/:id",delivery.GetByIdProduct)
		product.GET("/get",delivery.GetAllProducts)
	}

	task := router.Group("/task")
	{
		task.POST("/create",delivery.CreateTask)
		task.PUT("/update",delivery.UpdateTask)
		task.DELETE("/delete/:id", delivery.DeleteTask)
		task.GET("/get/:id",delivery.GetByIdTask)
		task.GET("/get",delivery.GetAllTasks)
	}
	return router
}
