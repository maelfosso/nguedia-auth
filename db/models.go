package db

import (
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FullName string             `json:"fullname,omitempty" bson:"fullname,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitemtpy"`
	Phone    string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
}

func (u *User) MarshalJSON() ([]byte, error) {
	type Alias User

	u.Password = ""
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(u),
	})
}

type ResetPassword struct {
	ID        primitive.ObjectID
	Email     string    `json:"email,omitempty" bson:"email,omitempty"`
	Token     string    `json:"token,omitempty" bson:"token,omitempty"`
	ExpiredAt time.Time `json:"expiredAt,omitempty" bson:"expiredAt,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}
