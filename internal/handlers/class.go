package handlers

import (
	"net/http"

	"github.com/as-ifn-at/glofox/common"
	"github.com/as-ifn-at/glofox/internal/config"
	"github.com/as-ifn-at/glofox/models"
	"github.com/gin-gonic/gin"
)

var Classes = make(map[string]models.Class, 0)

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
	if _, ok := Classes[id]; !ok {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, Classes[id])
}

func (h *classHandler) Save(ctx *gin.Context) {
	var newClass models.Class
	if err := ctx.BindJSON(&newClass); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := common.CheckValidStartEndDate(newClass.StartDate, newClass.EndDate); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	Classes[newClass.ClassName] = newClass

	ctx.IndentedJSON(http.StatusCreated, gin.H{"class created": newClass})
}
