package routes

import (
	"github.com/as-ifn-at/glofox/internal/handlers"
)

func (r *router) classesRoutes() {

	classHandler := handlers.NewClassHandler(r.appConfig, r.logger, r.db)
	routerG := r.router.Group("/classes/v1")
	routerG.POST("/create", classHandler.Save)
	routerG.GET("/:id", classHandler.Get)
}
