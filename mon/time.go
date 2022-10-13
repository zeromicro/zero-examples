package main

import (
	"context"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Roster struct {
	Id          primitive.ObjectID `bson:"_id"`
	CreateTime  time.Time          `bson:"createTime"`
	DisplayName string             `bson:"displayName"`
}

func main() {
	model := mon.MustNewModel("mongodb://root:example@127.0.0.1:27017", "db", "user")

	r := &Roster{
		Id:          primitive.NewObjectID(),
		CreateTime:  time.Now(),
		DisplayName: "Hello",
	}
	ctx := context.Background()
	_, err := model.InsertOne(ctx, r)
	if err != nil {
		panic(err)
	}

	update := bson.M{"$set": bson.M{
		"displayName": "Hello world",
		"createTime":  time.Now(),
	}}
	_, err = model.UpdateByID(ctx, r.Id, update)
	if err != nil {
		panic(err)
	}

	r.DisplayName = "Hello world!"
	_, err = model.ReplaceOne(ctx, bson.M{"_id": r.Id}, r)
	if err != nil {
		panic(err)
	}

	var tr Roster
	err = model.FindOne(ctx, &tr, bson.M{"_id": r.Id})
	if err != nil {
		panic(err)
	}
}
