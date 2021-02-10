package main

import (
	"net/http"

	"github.com/mm0070/goAuthAPI/controller"
	"github.com/mm0070/goAuthAPI/middleware"
	"github.com/mm0070/goAuthAPI/service"

	"github.com/gin-gonic/gin"
)

func getTestEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func main() {
	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	server := gin.New()

	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	// server.GET("/test", getTestEndpoint)
	server.GET("/test", middleware.AuthorizeJWT(), getTestEndpoint)

	// server.GET("/test", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	port := "8080"
	server.Run(":" + port)

}
