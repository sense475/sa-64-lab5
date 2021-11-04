package main


import (

  "github.com/sense475/sa-64-lab5/controller"

  "github.com/sense475/sa-64-lab5/entity"
  "github.com/sense475/sa-64-lab5/middlewares"
  "github.com/gin-gonic/gin"

)


func main() {

  entity.SetupDatabase()

  r := gin.Default()
	r.Use(CORSMiddleware())
	api := r.Group("")
	{
    protected := api.Use(middlewares.Authorizes())
		{
	// User Routes
	protected.GET("/users", controller.ListUser)
	protected.GET("/users/:id", controller.GetUser)
	protected.POST("/user", controller.CreateUser)


	// Patient Routes
	protected.GET("/patients", controller.ListPatient)
	protected.GET("/patient/:id", controller.GetPatient)
	protected.POST("/patients", controller.CreatePatient)
	protected.PATCH("/patients", controller.UpdatePatient)
	protected.DELETE("/patients/:id", controller.DeletePatient)

	// treatment Routes
	protected.GET("/treatments", controller.ListTreatment)
	protected.GET("/treatment/:id", controller.GetTreatment)
	protected.POST("/treatments", controller.CreateTreatment)
	protected.PATCH("/treatments", controller.UpdateTreatment)
	protected.DELETE("/treatments/:id", controller.DeleteTreatment)

	// payment Routes
	protected.GET("/payments", controller.ListPayment)
	protected.GET("/payment/:id", controller.GetPayment)
	protected.POST("/payments", controller.CreatePayment)
	protected.PATCH("/payments", controller.UpdatePayment)
	protected.DELETE("/payments/:id", controller.DeletePayment)

	// role
	protected.GET("/roles", controller.ListRole)
	protected.POST("/role", controller.CreateRole)

	//remedy
	protected.GET("/remedy_types", controller.ListRemedyType)
	protected.POST("/remedy_type", controller.CreateRemedyType)
      }
  }

	// Run the server

		// Authentication Routes
		r.POST("/login", controller.Login)
	r.Run()



}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}