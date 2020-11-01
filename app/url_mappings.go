package app

import (
	"github.com/vitorreis/bookstore-users-api/controllers/health_check"
	"github.com/vitorreis/bookstore-users-api/controllers/users"
)

func mapUrls() {
	router.GET("/health/check", health_check.HealthCheck)

	router.POST("/v1/users", users.CreateUser)
	// router.GET("/v1/users/search", controllers.SearchUser)
	router.GET("/v1/users/:user_id", users.GetUser)

}
