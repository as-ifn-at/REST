package handlers

import (
	"net/http"

	"github.com/as-ifn-at/glofox/internal/config"
	"github.com/as-ifn-at/glofox/models"
	"github.com/gin-gonic/gin"
)

var bokingsArr = []models.ClassBooking{}

type bookingHandler struct {
	Handler
	config config.Config
	// logger
	// db
}

func NewBookingHandler(config config.Config) *bookingHandler {
	return &bookingHandler{
		config: config,
	}
}

func (h *bookingHandler) Get(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, class := range bokingsArr {
		if class.MemberName == id {
			ctx.IndentedJSON(http.StatusOK, class)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "member not found"})
}

func (h *bookingHandler) Save(ctx *gin.Context) {
	var newClassBooking models.ClassBooking
	if err := ctx.BindJSON(&newClassBooking); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	bokingsArr = append(bokingsArr, newClassBooking)

	ctx.IndentedJSON(http.StatusCreated, newClassBooking)
}
