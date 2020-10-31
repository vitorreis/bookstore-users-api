package app

import "github.com/vitorreis/bookstore-users-api/controllers"

func mapUrls() {
	router.GET("/health/check", controllers.HealthCheck)
}
