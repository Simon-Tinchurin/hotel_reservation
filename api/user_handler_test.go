package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"hotel-reservation/customTypes"
	"hotel-reservation/db"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	testDbUri = "mongodb://localhost:27017"
	dbname    = "hotel-reservation-test"
)

type testDb struct {
	db.UserStore
}

func (tdb *testDb) teardown(t *testing.T) {
	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testDb {
	// connection to the mongodb
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testDbUri))
	if err != nil {
		log.Fatal(err)
	}
	return &testDb{
		UserStore: db.NewMongoUserStore(client, dbname),
	}
}

func TestPostUser(t *testing.T) {
	// simple setup for the test DB
	tdb := setup(t)
	// drop test DB after test
	defer tdb.teardown(t)

	app := fiber.New()
	userHandler := NewUserHandler(tdb.UserStore)
	// simple route for testing POST request
	app.Post("/", userHandler.HandlePostUser)

	params := customTypes.CreateUserParams{
		FirstName: "James",
		LastName:  "Foo",
		Email:     "some@foo.com",
		Password:  "123asdfaed43",
	}
	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	response, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}
	// decoding the response body to the user struct
	var user customTypes.User
	json.NewDecoder(response.Body).Decode(&user)
	if len(user.Id) == 0 {
		t.Errorf("expected a user id to be set")
	}
	if len(user.EncryptedPassword) > 0 {
		t.Errorf("expected an encrypted password not to be included in the json response")
	}
	// checking if decoded response params equals original params
	if user.FirstName != params.FirstName {
		t.Errorf("expected firstname %s but got %s", params.FirstName, user.FirstName)
	}
	if user.LastName != params.LastName {
		t.Errorf("expected lastname %s but got %s", params.LastName, user.LastName)
	}
	if user.Email != params.Email {
		t.Errorf("expected email %s but got %s", params.Email, user.Email)
	}
	fmt.Println(user)
}
