package main

// Next 22

import (
	"context"
	"flag"
	"fmt"
	"hotel-reservation/api"
	"hotel-reservation/customTypes"
	"hotel-reservation/db"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"
const dbname = "hotel-reservation"
const userCollection = "users"

var config = fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return ctx.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	// connection to the mongodb
	client, err := mongo.Connect(context.TODO(),
		options.Client().ApplyURI(dburi))

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	// connect to the user collection
	userCol := client.Database(dbname).Collection(userCollection)
	// create user
	// user := customTypes.User{
	// 	FirstName: "James",
	// 	LastName:  "At the watercooler",
	// }
	// inserrt the user
	// result, err := userCol.InsertOne(ctx, user)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// find first user, decode, find
	// and inject the values in the fields of custom type User
	var james customTypes.User
	if err := userCol.FindOne(ctx, bson.M{}).Decode(&james); err != nil {
		log.Fatal(err)
	}

	fmt.Println(james)

	listenAddr := flag.String("listenAddr", ":5000",
		"The listen address of the API server")
	flag.Parse()

	app := fiber.New(config)
	apiv1 := app.Group("/api/v1")

	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)

	app.Listen(*listenAddr)
}
