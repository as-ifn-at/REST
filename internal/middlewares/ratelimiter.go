package middlewares

import (
	"net/http"
	"time"

	"github.com/as-ifn-at/glofox/common"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimit() gin.HandlerFunc {

	limiter := rate.NewLimiter(per(time.Second, common.MaxNoOfRequestAllowed), common.MaxNoOfRequestAllowed)
	return func(ctx *gin.Context) {
		if !limiter.Allow() {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"message": "too many requests"})
		}
	}
}

func per(duration time.Duration, eventCount int) rate.Limit {
	return rate.Every(duration / time.Duration(eventCount))
}
