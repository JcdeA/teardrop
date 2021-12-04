package storage

import (
	"context"
	"time"

	"github.com/fosshostorg/teardrop/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBClient struct {
	*mongo.Client
}

func NewDBClient(options *options.ClientOptions) (*DBClient, error) {
	mongoClient, err := mongo.NewClient(options)
	if err != nil {
		return nil, err
	}
	return &DBClient{mongoClient}, nil
}

func (d *DBClient) Connect(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, err
}

func (d *DBClient) NewUser(ctx context.Context, user models.User) {
	users := d.Database("teardrop-v1").Collection("users")
	users.InsertOne(ctx, user, nil)
}
