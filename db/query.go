package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Query interface {
	FindUserByEmail(email string) (*User, error)
	SaveUser(user *User)

	DeleteResetPassword(email string) error
	SaveResetPasswordToken(email, token string) (*ResetPassword, error)
	FindResetPassword(email, token string) (*ResetPassword, error)
	UpdateUserPassword(user User, password string) error
}

type DBQuery struct {
	database *mongo.Database
}

func NewQuery(database *mongo.Database) *DBQuery {
	return &DBQuery{
		database: database,
	}
}

func (db DBQuery) collection(coll string) *mongo.Collection {
	return db.database.Collection(coll)
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

func (db DBQuery) SaveResetPasswordToken(email, token string) (*ResetPassword, error) {
	now := time.Now()
	ttl := 10 * time.Minute
	data := ResetPassword{
		Email:     email,
		Token:     token,
		ExpiredAt: now.Add(ttl),
		CreatedAt: now,
	}

	insertResult, err := db.collection("passwords").InsertOne(context.Background(), data)
	if err != nil {
		return nil, err
	}
	data.ID = insertResult.InsertedID.(primitive.ObjectID)

	return &data, nil
}

func (db DBQuery) DeleteResetPassword(email string) error {
	if _, err := db.collection("passwords").DeleteOne(context.Background(), bson.M{"email": email}); err != nil {
		return err
	}

	return nil
}

func (db DBQuery) FindResetPassword(email, token string) (*ResetPassword, error) {
	var rp ResetPassword
	var rpColl = db.database.Collection("passwords")

	if err := rpColl.FindOne(context.Background(), bson.M{"email": email, "token": token}).Decode(&rp); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &rp, nil
}

func (db DBQuery) UpdateUserPassword(u User, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	_, err = db.collection("users").UpdateByID(
		context.Background(),
		u.ID,
		bson.M{
			"$set": bson.M{
				"password": string(hash),
			},
		},
	)
	if err != nil {
		return err
	}

	return nil
}
