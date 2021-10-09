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
	clientOptions := options.Client().ApplyURI("mongodb+srv://amit:9mnYIInwK7yDu0IF@cluster0.l0avf.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	repoClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return errors.Wrap(err, "error-connecting-to-mongodb-repo")
	}
	return nil
}
