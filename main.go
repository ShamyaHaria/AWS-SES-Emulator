package main

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var emailCount=0
var emailLimit=5
var firstEmailTime time.Time

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
		c.JSON(http.StatusOK,gin.H{"message":"mock-raw-message-id-456",
		})
	})

	r.GET("/get-send-quota",func(c*gin.Context){
		resetLimitIfNeeded()
		c.JSON(http.StatusOK,gin.H{
			"max24HourSend":200,
			"sentLast24Hours":emailCount,
			"MaxSendRate":1.0,
		})
	})

	r.GET("/get-send-statistics",func(c*gin.Context){
		resetLimitIfNeeded()
		c.JSON(http.StatusOK,gin.H{
			"DeliveryAttempts":emailCount,
			"Bounces":0,
			"Complaints":0,
			"Rejects":0,
		})
	})

	r.Run(":8080")
}

func resetLimitIfNeeded() {
	if firstEmailTime.IsZero(){
		return
	}
	if time.Since(firstEmailTime)>24*time.Hour{
		emailCount=0
		if emailLimit == 5{
			emailLimit=10
		}
		else if emailLimit == 10{
			emailLimit=20
		}
		else if emailLimit == 20{
			emailLimit=50
		}
		firstEmailTime=time.Now()
	}
}