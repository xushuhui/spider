package core

import (
	"context"

	"github.com/qiniu/qmgo"
)

func initDB() {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb://localhost:27017"})
	if err != nil {
		return
	}
	db := client.Database("class")
	coll := db.Collection("user")
	println(coll)

}
