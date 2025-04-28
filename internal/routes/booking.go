package routes

import (
	"github.com/as-ifn-at/REST/internal/handlers"
)

func (r *router) attendClassesRoutes() {

	bookingHandler := handlers.NewBookingHandler(r.appConfig, r.logger)
	routerG := r.router.Group("/bookings/v1")
	routerG.POST("/book", bookingHandler.Save)
	routerG.GET("/:id", bookingHandler.Get)
}
