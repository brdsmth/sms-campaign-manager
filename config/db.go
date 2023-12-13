package config

import (
	"context"
	"fmt"
	"log"
	"sms-campaign-manager/graph/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

// ConnectDB initializes and returns a new DB object with a MongoDB client.
func ConnectDB() *DB {
	// Retrieve MongoDB connection string from environment variables for security
	mongoDBURI := readEnv("MONGODB_URI")
	if mongoDBURI == "" {
		log.Fatal("MONGODB_URI environment variable not set")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoDBURI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ping MongoDB to confirm connection
	if err := client.Database("admin").RunCommand(context.Background(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	fmt.Println("Connected to MongoDB successfully.")

	return &DB{client: client}
}

func colHelper(db *DB, collectionName string) *mongo.Collection {
	return db.client.Database("hooksms").Collection(collectionName)
}

// Campaigns
func (db *DB) CreateCampaign(input *model.NewCampaign) (*model.Campaign, error) {
	collection := colHelper(db, "campaigns")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, input)

	if err != nil {
		return nil, err
	}

	owner := &model.Campaign{
		ID:        res.InsertedID.(primitive.ObjectID).Hex(),
		Name:      input.Name,
		StartTime: input.StartTime,
	}

	return owner, err
}

func (db *DB) GetCampaigns() ([]*model.Campaign, error) {
	collection := colHelper(db, "campaigns")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var campaigns []*model.Campaign
	defer cancel()

	res, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer res.Close(ctx)
	for res.Next(ctx) {
		var singleCampaign *model.Campaign
		if err = res.Decode(&singleCampaign); err != nil {
			log.Fatal(err)
		}
		campaigns = append(campaigns, singleCampaign)
	}

	return campaigns, err
}

// Segments
func (db *DB) CreateSegment(input *model.NewSegment) (*model.Segment, error) {
	collection := colHelper(db, "segments")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, input)

	if err != nil {
		return nil, err
	}

	segment := &model.Segment{
		ID:          res.InsertedID.(primitive.ObjectID).Hex(),
		Name:        input.Name,
		Description: input.Description,
	}

	return segment, err
}

func (db *DB) GetSegments() ([]*model.Segment, error) {
	collection := colHelper(db, "segments")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var segments []*model.Segment
	defer cancel()

	res, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer res.Close(ctx)
	for res.Next(ctx) {
		var singleSegment *model.Segment
		if err = res.Decode(&singleSegment); err != nil {
			log.Fatal(err)
		}
		segments = append(segments, singleSegment)
	}

	return segments, err
}

// Templates
func (db *DB) CreateTemplate(input *model.NewTemplate) (*model.Template, error) {
	collection := colHelper(db, "templates")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, input)

	if err != nil {
		return nil, err
	}

	template := &model.Template{
		ID:      res.InsertedID.(primitive.ObjectID).Hex(),
		Name:    input.Name,
		Content: input.Content,
	}

	return template, err
}

func (db *DB) GetTemplates() ([]*model.Template, error) {
	collection := colHelper(db, "templates")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var templates []*model.Template
	defer cancel()

	res, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer res.Close(ctx)
	for res.Next(ctx) {
		var singleTemplate *model.Template
		if err = res.Decode(&singleTemplate); err != nil {
			log.Fatal(err)
		}
		templates = append(templates, singleTemplate)
	}

	return templates, err
}
