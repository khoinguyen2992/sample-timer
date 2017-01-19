package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/jasonlvhit/gocron"
	gin "gopkg.in/gin-gonic/gin.v1"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	session *mgo.Session
	once    sync.Once
)

type task struct {
	OrderId string `json:"order_id" bson:"order_id"`
	State   string `json:"state" bson:"state"`
}

func init() {
	once.Do(func() {
		s, err := mgo.Dial("mongodb:27017")
		if err != nil {
			panic(err)
		}

		session = s
		session.SetMode(mgo.Monotonic, true)
	})
}

func processing(componentName string, orderId string) {
	time.Sleep(time.Duration(rand.Int31n(3)) * time.Second)
	fmt.Printf("Component: %s processed task: %s \n", componentName, orderId)
}

func getCollection(dbName string) *mgo.Collection {
	return session.Copy().DB(dbName).C("tasks")
}

func cronjob() {
	collection := getCollection("app2")
	s := gocron.NewScheduler()
	s.Every(1).Second().Do(func(componentName string) {
		tasks := []task{}
		if err := collection.Find(bson.M{"state": "processing"}).All(&tasks); err != nil {
			panic(err)
		}

		fmt.Printf("%s will process: %v\n", componentName, tasks)

		for _, t := range tasks {
			processing(componentName, t.OrderId)
			if err := collection.Update(bson.M{"order_id": t.OrderId, "state": "processing"},
				bson.M{"$set": bson.M{"state": "done"}}); err != nil {
				fmt.Println(err)
			}
		}
	}, os.Getenv("COMPONENT_NAME"))

	<-s.Start()
	defer s.Clear()
}

func main() {
	go cronjob()

	r := gin.Default()
	r.GET("/add_processing_task/:id", func(c *gin.Context) {
		id := c.Param("id")
		collection := getCollection("app2")
		collection.Insert(task{OrderId: id, State: "processing"})
	})
	r.Run()
}
