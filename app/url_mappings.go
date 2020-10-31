package app

import (
	"github.com/vitorreis/bookstore-users-api/controllers"
	"github.com/vitorreis/bookstore-users-api/controllers/health_check"
)

func mapUrls() {
	router.GET("/health/check", health_check.HealthCheck)

	router.POST("/v1/users", controllers.CreateUser)
	// router.GET("/v1/users/search", controllers.SearchUser)
	router.GET("/v1/users/:user_id", controllers.GetUser)

}
