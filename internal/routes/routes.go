package routes

import (
	"net/http"

	"github.com/as-ifn-at/glofox/internal/config"
	"github.com/gin-gonic/gin"
)

type router struct {
	router    *gin.Engine
	appConfig config.Config
	// logger
	// db
}

func NewRouter(config *config.Config) *router {
	return &router{
		router:    gin.Default(),
		appConfig: *config,
	}
}

func (r *router) SetRouters() http.Handler {
	r.classesRoutes()
	r.attendClassesRoutes()

	return r.router.Handler()
}
