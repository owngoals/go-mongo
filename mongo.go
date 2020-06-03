package gomongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Configs struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func NewMongo(host string, port int, db string) (*mongo.Database, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	c, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port)))
	if err != nil {
		return nil, err
	}
	if err := c.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}
	return c.Database(db), nil
}
