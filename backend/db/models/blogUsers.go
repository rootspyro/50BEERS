package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type BlogUser struct {
	ID        primitive.ObjectID `bson:"_id"`
	Username  string             `bson:"username"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	Origin    string             `bson:"origin"`
	CreatedAt string             `bson:"created_at"`
	UpdatedAt string             `bson:"updated_at"`
}

type NewBlogUser struct {
	Username string `bson:"username"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Origin   string `bson:"origin"`
	CreatedAt string             `bson:"created_at"`
	UpdatedAt string             `bson:"updated_at"`
}
