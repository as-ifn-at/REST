package db

import "github.com/as-ifn-at/glofox/internal/config"

type DbHandler struct {
}

func NewDbHandler(config *config.Config) *DbHandler {

	switch (config.DatabaseName) {
	case "badgerdb":
	default:
		panic("no database selected")
	}

	return &DbHandler{}
}
