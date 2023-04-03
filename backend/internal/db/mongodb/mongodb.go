package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type configDB struct {
	HOST     string `json:"host"`
	PORT     string `json:"port"`
	DATABASE string `json:"database"`
}

func NewClient(ctx context.Context, config configDB) (*mongo.Database, error) {

	mongoDBURL := fmt.Sprintf("mongdb://%s:%s", config.HOST, config.PORT)

	clientOptions := options.Client().ApplyURI(mongoDBURL)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongoDB dur to error %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping mongoDB due to error %v", err)
	}

	return client.Database(config.DATABASE), nil

}
