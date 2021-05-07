package db

import "go.mongodb.org/mongo-driver/mongo"

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

func (db DBQuery) FindUserByEmail(email string) {}
func (db DBQuery) SaveUser()                    {}

func (db DBQuery) DeleteResetPassword(email string)           {}
func (db DBQuery) SaveResetPasswordToken(email, token string) {}
func (db DBQuery) FindResetPassword(email, token string)      {}
func (db DBQuery) UpdateUserPassword()                        {}
