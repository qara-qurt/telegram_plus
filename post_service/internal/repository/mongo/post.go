package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Post struct {
	mongodb *mongo.Database
}

func NewPost(mongodb *mongo.Database) *Post {
	return &Post{
		mongodb: mongodb,
	}
}
