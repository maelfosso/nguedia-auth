package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Query interface {
	FindUserByEmail(email string)
	SaveUser()

	DeleteResetPassword(email string)
	SaveResetPasswordToken(email, token string)
	FindResetPassword(email, token string)
	UpdateUserPassword()
}

type DBQuery struct {
	database *mongo.Database
}

func NewQuery(database *mongo.Database) *DBQuery {
	return &DBQuery{
		database: database,
	}
}

func (db DBQuery) FindUserByEmail(email string) (*User, error) {
	var user User

	if err := db.database.Collection("users").FindOne(context.Background(), bson.M{"email": email}).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &user, nil
}

func (db DBQuery) SaveUser(user *User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	var usersColl = db.database.Collection("users")
	user.Password = string(hash)
	insertResult, err := usersColl.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	user.ID = insertResult.InsertedID.(primitive.ObjectID)
	return nil
}

func (db DBQuery) DeleteResetPassword(email string)           {}
func (db DBQuery) SaveResetPasswordToken(email, token string) {}
func (db DBQuery) FindResetPassword(email, token string)      {}
func (db DBQuery) UpdateUserPassword()                        {}
