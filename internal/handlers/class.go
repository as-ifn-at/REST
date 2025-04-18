package handlers

import (
	"net/http"

	"github.com/as-ifn-at/glofox/internal/config"
	"github.com/as-ifn-at/glofox/models"
	"github.com/gin-gonic/gin"
)

var classesArr = []models.Class{}

type classHandler struct {
	Handler
	config config.Config
	// logger
	// db
}

func NewClassHandler(config config.Config) *classHandler {
	return &classHandler{
		config: config,
	}
}

func (h *classHandler) Get(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, class := range classesArr {
		if class.ClassName == id {
			ctx.IndentedJSON(http.StatusOK, class)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func (h *classHandler) Save(ctx *gin.Context) {
	var newClass models.Class
	if err := ctx.BindJSON(&newClass); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	classesArr = append(classesArr, newClass)

	ctx.IndentedJSON(http.StatusCreated, newClass)
}
