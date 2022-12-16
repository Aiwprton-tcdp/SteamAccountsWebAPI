package app

import (
	"log"
	"net/http"
	"os"
	"sawa/app/services"
	"sawa/controllers"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Initialize() {
	gin.SetMode(os.Getenv("GIN_MODE"))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// router.Use(middleware.AuthorizeJWT())
	router.SetTrustedProxies(nil)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	initializeRoutes()

	router.Run()
}

func initializeRoutes() {
	var loginService services.LoginService = services.StaticLoginService()
	var jwtService services.JWTService = services.JWTAuthService()
	var loginController controllers.LoginController = controllers.LoginHandler(loginService, jwtService)

	router.GET("", func(c *gin.Context) {
		log.Printf("ClientIP: %s\n", c.ClientIP())
	})

	router.NoRoute(func(c *gin.Context) {
		c.Status(http.StatusNotFound)
	})

	router.POST("/login", func(c *gin.Context) {
		token := loginController.Login(c)
		if token != "" {
			c.JSON(http.StatusOK, gin.H{"data": token})
		} else {
			c.Status(http.StatusUnauthorized)
		}
	})

	accounts := router.Group("/accounts")
	accounts.GET("", controllers.GetAccounts)
	accounts.POST("", controllers.AddAccounts)
	accounts.GET("/lps", controllers.GetLPS)

	account := accounts.Group("/:id")
	account.GET("", controllers.GetAccount)
	account.PATCH("", controllers.UpdateAccount)
	account.DELETE("", controllers.DeleteAccount)
	account.GET("/guardcode", controllers.GetGuardCode)
	account.GET("/summaries", controllers.GetAccountSummaries)
	account.GET("/check", controllers.CheckAvailability)
}
