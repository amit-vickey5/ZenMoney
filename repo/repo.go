package repo

import (
	"context"

	"amit.vickey/ZenMoney/common"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckHealth(ctx *gin.Context) error {
	err := repoClient.Ping(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "error-pinging-database")
	}
	return nil
}

func AddConsentForUser(ctx *gin.Context, userId string, consentInfo UserConsentInfo) error {
	userCollection := repoClient.Database(userId).Collection("consents")
	_, err := userCollection.InsertOne(ctx, consentInfo)
	if err != nil {
		return errors.Wrap(err, "error-adding-consent-for-user")
	}
	return nil
}

func AddFIDataForUser(ctx *gin.Context, userId string, fiDataType string, fiData interface{}) error {
	userDataTypeCollection := repoClient.Database(userId).Collection(fiDataType)
	result, err := userDataTypeCollection.InsertOne(ctx, fiData)
	if err != nil {
		return errors.Wrap(err, "error-adding-user-fi-data")
	}
	common.LogInfo(ctx, result.InsertedID.(primitive.ObjectID).Hex())
	return nil
}

func GetUserFIDataForType(ctx *gin.Context, userId, dataType string) ([]interface{}, error) {
	var finalDataList []interface{}
	userDataTypeCollection := repoClient.Database(userId).Collection(dataType)
	cur, err := userDataTypeCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, errors.Wrap(err, "error-fetching-fi-data-for-user")
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		rawData := cur.Current
		var s map[string]interface{}
		if err := bson.Unmarshal(rawData, &s); err != nil {
			return nil, err
		}
		finalDataList = append(finalDataList, s)
	}
	return finalDataList, nil
}
