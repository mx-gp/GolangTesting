package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

// HandleError: Error Handler Function
func HandleError(c *gin.Context, err error, context string, code int) {
	log.WithFields(logrus.Fields{
		"error": err,
	}).Errorln(context)

	c.JSON(code, gin.H{
		"code":    code,
		"error":   err,
		"message": context,
		"success": false,
	})
}
