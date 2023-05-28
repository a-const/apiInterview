package repo

import (
	user "apiInterview/pkg/type"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
}

func NewMongoDB() *MongoDB {
	return &MongoDB{}
}

func (mdb *MongoDB) Connect(port string) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://db:%s", port)).
		SetConnectTimeout(20 * time.Second).
		SetServerSelectionTimeout(2 * time.Second).
		SetSocketTimeout(1 * time.Hour).
		SetMaxPoolSize(10)
	var err error
	mdb.client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		logrus.Fatalf(err.Error())
	}
	err = mdb.client.Ping(ctx, nil)
	if err != nil {
		logrus.Fatalf(err.Error())
	}
	mdb.db = mdb.client.Database("usrdb")
	mdb.collection = mdb.db.Collection("users")

	indexOptions := options.Index().SetUnique(true)
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"username": 1},
		Options: indexOptions,
	}
	mdb.collection.Indexes().CreateOne(ctx, indexModel)
}

func (mdb *MongoDB) CreateUser(user *user.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := mdb.collection.InsertOne(ctx, *user); err != nil {
		return err
	}
	return nil

}

func (mdb *MongoDB) DeleteUser(username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"username": username}
	res, err := mdb.collection.DeleteOne(ctx, filter)
	if res.DeletedCount == 0 {
		err = errors.New("username not found")
		return err
	}
	return err
}

func (mdb *MongoDB) UpdateUser(username string, password string, description string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"username": username}
	update := bson.M{"$set": bson.M{
		"description": description,
		"password":    password,
	}}
	res, err := mdb.collection.UpdateOne(ctx, filter, update)
	if res.MatchedCount == 0 {
		err = errors.New("username not found")
		return err
	}
	return err
}

func (mdb *MongoDB) GetUser(username string) (*user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"username": username}
	user := &user.User{}
	if err := mdb.collection.FindOne(ctx, filter).Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (mdb *MongoDB) GetAllUsers() (*[]user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var users []user.User
	findOpt := options.Find()
	cur, err := mdb.collection.Find(ctx, bson.M{}, findOpt)
	for cur.Next(ctx) {
		var usr user.User
		cur.Decode(&usr)
		users = append(users, usr)
	}
	return &users, err
}
