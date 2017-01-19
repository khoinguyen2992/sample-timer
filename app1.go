package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	gin "gopkg.in/gin-gonic/gin.v1"
	redis "gopkg.in/redis.v5"
)

func timerJob(orderId string) {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	pubsub, err := client.Subscribe("fake")
	if err != nil {
		panic(err)
	}
	defer pubsub.Close()

	for i := 0; i < 2; i++ { //subscribe, then timeout
		if _, err := pubsub.ReceiveTimeout(5 * time.Second); err != nil {
			processing(os.Getenv("COMPONENT_NAME"), orderId)
			break
		}
	}
}

func processing(componentName string, orderId string) {
	time.Sleep(time.Duration(rand.Int31n(3)) * time.Second)
	fmt.Printf("Component: %s processed task: %s \n", componentName, orderId)
}

func main() {
	r := gin.Default()
	r.GET("/add_processing_task/:id", func(c *gin.Context) {
		id := c.Param("id")
		go timerJob(id)
	})
	r.Run()
}
