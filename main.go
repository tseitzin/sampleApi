package main

import (
	"sampleApi/controllers"
	_ "sampleApi/docs"
	"sampleApi/middlewares"
	"sampleApi/models"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title     Tire Tread Application API
// @version         1.0
// @description     A Customer, Vehicle and Tire API in Go using Gin framework.

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @contact.name   Tim Seitzinger
// @contact.email  tseitzinger@gmail.com

// @host      localhost:8080
// @BasePath  /
func main() {

	r := setupRouter()

	r.Run()

	// r := gin.Default()

	// r.POST("/register", controllers.Register)
	// r.POST("/login", controllers.Login)

	// r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// r.Use(middlewares.JwtAuthMiddleware())

	// models.ConnectDatabase()

	// r.GET("/user", controllers.CurrentUser)

	// r.GET("/customers", controllers.FindCustomers)
	// r.GET("/vehicles", controllers.FindVehicles)
	// r.GET("/tires", controllers.FindTires)

	// r.GET("/customers/:id", controllers.FindCustomer)
	// r.GET("/vehicles/:v_id", controllers.FindVehicle)
	// r.GET("/tires/:t_id", controllers.FindTire)

	// r.POST("/customers", controllers.CreateCustomer)
	// r.POST("/vehicles", controllers.CreateVehicle)
	// r.POST("/tires", controllers.CreateTire)

	// r.DELETE("/customers/:id", controllers.DeleteCustomer)
	// r.DELETE("/tires/:t_id", controllers.DeleteTire)
	// r.DELETE("/vehicles/:v_id", controllers.DeleteVehicle)

	// r.PATCH("/customers/:id", controllers.UpdateCustomer)
	// r.PATCH("/tires/:t_id", controllers.UpdateTire)
	// r.PATCH("/vehicles/:v_id", controllers.UpdateVehicle)

	// err1 := r.Run()
	// if err1 != nil {
	// 	return
	// }
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(middlewares.JwtAuthMiddleware())

	models.ConnectDatabase()

	r.GET("/user", controllers.CurrentUser)

	r.GET("/companies", controllers.FindCompanies)
	r.GET("/customers", controllers.FindCustomers)
	r.GET("/vehicles", controllers.FindVehicles)
	r.GET("/tires", controllers.FindTires)

	r.GET("/customers/:id", controllers.FindCustomer)
	r.GET("/vehicles/:v_id", controllers.FindVehicle)
	r.GET("/tires/:t_id", controllers.FindTire)

	r.POST("/companies", controllers.CreateCompany)
	r.POST("/customers", controllers.CreateCustomer)
	r.POST("/vehicles", controllers.CreateVehicle)
	r.POST("/tires", controllers.CreateTire)

	r.DELETE("/customers/:id", controllers.DeleteCustomer)
	r.DELETE("/tires/:t_id", controllers.DeleteTire)
	r.DELETE("/vehicles/:v_id", controllers.DeleteVehicle)

	r.PATCH("/customers/:id", controllers.UpdateCustomer)
	r.PATCH("/tires/:t_id", controllers.UpdateTire)
	r.PATCH("/vehicles/:v_id", controllers.UpdateVehicle)

	return r
}
