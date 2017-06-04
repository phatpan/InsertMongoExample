package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main(){
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	
	collection := session.DB("income").C("users")
	err = collection.Insert(bson.M{"name": "Phatcharaphan"})
	if err != nil {
		panic(err)
	}
}
