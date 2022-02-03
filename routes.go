package main

import (
	"github.com/gin-gonic/gin"
)

// LoadAPIRoutes takes the router object and adds this apps API endpoints
func LoadAPIRoutes(r *gin.Engine) {
	r.SetTrustedProxies(nil)
	r.POST("/wordoccurance", WordOccurance)
}
