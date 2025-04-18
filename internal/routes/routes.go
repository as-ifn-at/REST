package routes

import (
	"net/http"

	"github.com/as-ifn-at/glofox/internal/config"
	"github.com/as-ifn-at/glofox/internal/middlewares"
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
	attachMiddleWares(r.router)
	r.classesRoutes()
	r.attendClassesRoutes()

	return r.router.Handler()
}

func attachMiddleWares(router *gin.Engine) {
	router.Use(gin.Recovery())
	router.Use(middlewares.RateLimit())
}
