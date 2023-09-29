package handlers

import (
	"go-logger/database"
	model "go-logger/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func HandleGetAllRequestLogs(c *fiber.Ctx) error {
	collection := database.GetCollection("request_logs")

	filter := bson.M{}
	options := options.Find().SetSkip(0).SetLimit(100)

	// find all
	results, err := collection.Find(c.Context(), filter, options)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"internal server error a": err.Error()})
	}

	// convert results to slice
	requestLogs := make([]model.RequestLog, 0)
	if err = results.All(c.Context(), &requestLogs); err != nil {
		return c.Status(500).JSON(fiber.Map{"internal server error b": err.Error()})
	}

	return c.Status(200).JSON(requestLogs)
}

type CreateRequestLog struct {
	Method    string    `json:"method" bson:"method"`
	URI       string    `json:"uri" bson:"uri"`
	CreatedAt time.Time `json:"createAt" bson:"created_at"`
}

func HandleCreateRequestLog(c *fiber.Ctx) error {
	data := new(CreateRequestLog)

	// validate request body
	if err := c.BodyParser(data); err != nil {
		return c.Status(400).JSON(fiber.Map{"Bad input": err.Error()})
	}

	// set date time now
	data.CreatedAt = time.Now()

	collection := database.GetCollection("request_logs")
	res, err := collection.InsertOne(c.Context(), data)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"Add request log success id:": res.InsertedID})
}
