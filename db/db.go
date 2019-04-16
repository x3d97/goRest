package db

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID    string `json:"id" bson:"_id,omitempty"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone int    `json:"phone"`
}

var db *mgo.Database

func init() {
	session, err := mgo.Dial("mongodb://heroku_0dl5d90k:dii2h930p13v6l8luskli2ojgt@ds125526.mlab.com:25526/heroku_0dl5d90k")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db = session.DB("heroku_0dl5d90k")
}

func collection() *mgo.Collection {
	return db.C("users")
}

func GetAll() ([]User, error) {
	res := []User{}

	if err := collection().Find(nil).All(&res); err != nil {
		return nil, err
	}

	return res, nil
}

func GetOne(id string) (*User, error) {
	res := User{}

	if err := collection().Find(bson.M{"_id": id}).One(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

func Save(user User) error {
	return collection().Insert(user)
}

func Remove(id string) error {
	return collection().Remove(bson.M{"_id": id})
}
