package db

import (
	"context"
	"log"
	"time"

	"github.com/GeekCAMPvol2/QuizAPI/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	coll *mongo.Collection
}
type MongoDB struct {
	client *mongo.Client
}

func NewClient(uri string) *Client {
	return &Client{
		coll: NewDB(uri).client.Database("quiz").Collection("col"),
	}
}
func NewDB(uri string) *MongoDB {
	return &MongoDB{
		client: ConnectDB(uri),
	}
}

func ConnectDB(uri string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// 任意数のデータを取得する
func (client *Client) GetQuiz(size int64) (returnData []models.ReturnData, err error) {
	limitStage := bson.D{{"$limit", size}}

	cursor, err := client.coll.Aggregate(context.TODO(), mongo.Pipeline{limitStage})

	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {

		var elem models.ReturnData

		if err = cursor.Decode(&elem); err != nil {
			return
		}

		returnData = append(returnData, elem)
	}

	if err = cursor.Err(); err != nil {
		return
	}
	return
}
