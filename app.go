package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		requestID := uuid.New()
		logger.WithFields(logrus.Fields{
			"time":    time.Now().Format("2006-01-02 15:04:05"),
			"level":   "INFO",
			"message": "New request - " + requestID.String(),
		}).Info()

		time.Sleep(10 * time.Second)

		c.JSON(200, gin.H{"message": "success"})
	})

	router.Run(":8000")
}
