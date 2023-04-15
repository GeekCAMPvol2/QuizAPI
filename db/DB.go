package db

import (
	"context"
	"log"
	"os"
	"serv/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	coll mongo.Collection
}

func InitDB() *Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_DB_URI")))
	if err != nil {
		log.Fatalln(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(ctx)

	return &Client{coll: *client.Database("Quizs").Collection("col")}
}

func (c *Client) InsertDB(QuizData models.ReturnData) error {
	_, err := c.coll.InsertOne(context.Background(), QuizData)
	return err
}

func (c *Client) FindDB() []bson.M {
	collection, err := c.coll.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal("a", err)
	}
	var results []bson.M
	if err = collection.All(context.TODO(), &results); err != nil {
		log.Fatal("b", err)
	}
	return results
}

func Mainfind(hits int) []models.ReturnData {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://192.168.100.231:27017"))
	if err != nil {
		log.Fatal(err)
	}
	if err = client.Connect(context.Background()); err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	col := client.Database("quiz").Collection("col")
	cur, err := col.Aggregate(context.Background(), bson.A{
		bson.M{
			"$sample": bson.M{
				"size": hits,
			},
		},
	}, options.MergeAggregateOptions())

	if err != nil {
		log.Fatal(err)
	}
	var results []models.ReturnData

	for cur.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem models.ReturnData
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return results
}

func insertStruct(col *mongo.Collection, quizdata models.ReturnData) error {

	_, err := col.InsertOne(context.Background(), quizdata)
	return err
}

func Maininsert(quizdata models.ReturnData) error {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://192.168.100.231:27017"))
	if err != nil {
		return err
	}
	if err = client.Connect(context.Background()); err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	col := client.Database("quiz").Collection("col")

	if err = insertStruct(col, quizdata); err != nil {
		return err
	}
	return nil
}

func MainfindAll() []models.ReturnData {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://192.168.100.231:27017"))
	if err != nil {
		log.Fatal(err)
	}
	if err = client.Connect(context.Background()); err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	col := client.Database("quiz").Collection("col")
	cur, err := col.Find(context.Background(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}
	var results []models.ReturnData

	for cur.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem models.ReturnData
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return results
}
