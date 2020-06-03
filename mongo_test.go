package gomongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
	"time"
)

func TestNewMongo(t *testing.T) {
	host := "127.0.0.1"
	port := 27017
	db := "gomongo"

	d, err := NewMongo(host, port, db)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	r, err := d.Collection("applog").InsertOne(context.TODO(), bson.M{
		"name": "gomongo_test",
		"ts":   time.Now().UnixNano(),
	})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(r.InsertedID)
}
