package routes

import (
	"net/http"

	"github.com/as-ifn-at/REST/internal/config"
	"github.com/as-ifn-at/REST/internal/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type router struct {
	router    *gin.Engine
	appConfig config.Config
	logger    zerolog.Logger
}

func NewRouter(config *config.Config, logger zerolog.Logger) *router {
	return &router{
		logger:    logger,
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
