package model

import (
	"time"

	"github.com/jketcham/vicus/Godeps/_workspace/src/gopkg.in/mgo.v2"
	"github.com/jketcham/vicus/Godeps/_workspace/src/gopkg.in/mgo.v2/bson"

	"github.com/jketcham/vicus/shared/database"
)
var db = database.Database

// User is the model of users in mongodb and vicus service
type User struct {
	ID       bson.ObjectId `bson:"_id"`
	Email    string        `bson:"email"`
	Phone    uint64        `bson:"phone"`    // best way to define phone numbers?
	Role     string        `bson:"role"`     // TODO: make Role model
	Password string        `bson:"password"` // use bcrypt here

	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	Location  string `bson:"location"`
	Bio       string `bson:"bio"`

	CreatedAt  time.Time `bson:"created_at"`
	UpdatedAt  time.Time `bson:"updated_at"`
	LastActive time.Time `bson:"last_active"`
}

// CreateUser returns a new User
func CreateUser(email, password, firstName, lastName string) *User {
	return &User{
		ID:         bson.NewObjectId(),
		Email:      email,
		FirstName:  firstName,
		LastName:   lastName,
		Password:   password,
		CreatedAt:  time.Now(),
		LastActive: time.Now(),
	}
}

// Update receives email and password strings that are used to update the User
func (u *User) Update(email, password string) (User, error) {
	u.Email = email
	u.Password = password
	err := u.save()
	if err != nil {
		return *u, err
	}
	return *u, nil
}

// FindByEmail finds and assigns the User with the given email
func (u *User) FindByEmail(email string) error {
	return u.coll().Find(bson.M{"email": email}).One(u)
}

// FindByID finds and assigns the User with the given ObjectID
func (u *User) FindByID(id bson.ObjectId) error {
	return u.coll().FindId(id).One(u)
}

// Delete removes the User from the database
func (u *User) Delete() error {
	return u.coll().RemoveId(u.ID)
}

func (u *User) save() error {
	_, err := u.coll().UpsertId(u.ID, u)
	return err
}

func (u *User) coll() *mgo.Collection {
	return db.C("user")
}
