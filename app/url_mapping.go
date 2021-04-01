package app

import (
	"github.com/balwaninitu/nitu-s_bookstore_api/controllers/ping"
	"github.com/balwaninitu/nitu-s_bookstore_api/controllers/users"
)

func mapUrls() {
	/*request will be set to get so that if we running of cloud like
	aws it will hit the endpoint at url */
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/internal/users/search", users.Search)

}
