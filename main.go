package main

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var emailCount=0
var emailLimit=5
var totalEmailCount=0
var lastResetTime = time.Now()

func main(){
	r:=gin.Default()
	r.GET("/health",func(c*gin.Context){
		c.JSON(200,gin.H{"message":"API is running"})
	})

	r.POST("/send-email",func(c*gin.Context){
		resetLimitIfNeeded()
		if emailCount>=emailLimit{
			c.JSON(http.StatusTooManyRequests,gin.H{"error":"Sending Limit Exceeeded. Wait for 24hours!!",
			})
		return
		}
		emailCount++
		totalEmailCount++
		c.JSON(http.StatusOK,gin.H{"message":"mock-message-id-123",
		})
	})

	r.POST("/send-raw-email",func(c*gin.Context){
		resetLimitIfNeeded()
		if emailCount>=emailLimit{
			c.JSON(http.StatusTooManyRequests,gin.H{"error":"Sending Limit Exceeeded. Wait for 24hours!!",
			})
		return
		}
		emailCount++
		totalEmailCount++
		c.JSON(http.StatusOK,gin.H{"message":"mock-raw-message-id-456",
		})
	})

	r.GET("/get-send-quota",func(c*gin.Context){
		resetLimitIfNeeded()
		c.JSON(http.StatusOK,gin.H{
			"max24HourSend":emailLimit,
			"sentLast24Hours":emailCount,
			"TotalEmailCount":totalEmailCount,
			"MaxSendRate":1.0,
		})
	})

	r.GET("/get-send-statistics",func(c*gin.Context){
		resetLimitIfNeeded()
		c.JSON(http.StatusOK,gin.H{
			"DeliveryAttempts":emailCount,
			"TotalEmailsSent":totalEmailCount,
			"Bounces":0,
			"Complaints":0,
			"Rejects":0,
		})
	})

	r.Run(":8080")
}

func resetLimitIfNeeded() {
	now := time.Now()
	if now.Day() != lastResetTime.Day() {
		emailCount=0
		lastResetTime=now
		if totalEmailCount >= 25 && emailLimit == 5{
			emailLimit = 10
		} else if totalEmailCount >= 100 && emailLimit == 10{
			emailLimit = 25
		}else if totalEmailCount >= 250 && emailLimit == 25{
			emailLimit = 50
		}
	}
}