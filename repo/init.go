package repo

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var repoClient *mongo.Client

func InitRepoClient() error {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	repoClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return errors.Wrap(err, "error-connecting-to-mongodb-repo")
	}
	return nil
}
