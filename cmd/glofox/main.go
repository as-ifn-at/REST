package main

import (
	"fmt"
	"net/http"

	"github.com/as-ifn-at/glofox/internal/config"
	"github.com/as-ifn-at/glofox/internal/routes"
)

func main() {
	config := config.Load()
	router := routes.NewRouter(config).SetRouters()
	listenPort := fmt.Sprintf(":%v", config.Port)

	httpServer := &http.Server{
		Addr:    listenPort,
		Handler: router,
	}

	fmt.Println("listening on port: ", config.Port)
	if err := httpServer.ListenAndServe(); err != nil {
		panic(err)
	}
}
