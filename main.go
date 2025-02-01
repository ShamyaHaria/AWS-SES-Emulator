package main

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

var emailCount=0

func main(){
	r:=gin.Default()
	r.GET("/health",func(c*gin.Context){
		c.JSON(200,gin.H{"message":"API is running"})
	})

	r.POST("/send-email",func(c*gin.Context){
		emailCount++
		c.JSON(http.StatusOK,gin.H{"message":"mock-message-id-123",
		})
	})

	r.POST("/send-raw-email",func(c*gin.Context){
		emailCount++
		c.JSON(http.StatusOK,gin.H{"message":"mock-raw-message-id-456",
		})
	})

	r.GET("/get-send-quota",func(c*gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"max24HourSend":200,
			"sentLast24Hours":emailCount,
			"MaxSendRate":1.0,
		})
	})

	r.GET("/get-send-statistics",func(c*gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"DeliveryAttempts":emailCount,
			"Bounces":0,
			"Complaints":0,
			"Rejects":0,
		})
	})

	r.Run(":8080")
}