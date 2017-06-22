package db

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
)

var (
	Session *mgo.Session
	Mongo *mgo.DialInfo

	Todos *mgo.Collection
)

func Connect(MongoDBUrl string) {

	uri := os.Getenv("MONGODB_URL")

	if len(uri) == 0 {
		uri = MongoDBUrl
	}

	mongo, err := mgo.ParseURL(uri)
	session, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	session.SetSafe(&mgo.Safe{})
	session.SetMode(mgo.Monotonic, true)
	fmt.Println("Connected to", uri)
	Session = session
	Mongo = mongo

	//Collections
	Todos = session.DB(mongo.Database).C("todos")
}