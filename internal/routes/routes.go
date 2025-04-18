package routes

import (
	"net/http"

	"github.com/as-ifn-at/glofox/internal/config"
	"github.com/as-ifn-at/glofox/internal/db"
	"github.com/as-ifn-at/glofox/internal/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type router struct {
	router    *gin.Engine
	appConfig config.Config
	logger    zerolog.Logger
	db        *db.DbHandler
}

func NewRouter(config *config.Config, logger zerolog.Logger, db *db.DbHandler) *router {
	return &router{
		logger:    logger,
		router:    gin.Default(),
		appConfig: *config,
		db:        db,
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
