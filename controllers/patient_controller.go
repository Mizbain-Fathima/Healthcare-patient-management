package controllers

import (
	"context"
	"health-api/database"
	"health-api/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Get collection dynamically AFTER DB is connected
func getCollection() *mongo.Collection {
	return database.GetPatientCollection()
}

// CREATE
func CreatePatient(c *gin.Context) {
	var patient models.Patient
	c.BindJSON(&patient)

	patient.ID = primitive.NewObjectID().Hex()

	collection := getCollection()

	_, err := collection.InsertOne(context.Background(), patient)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, patient)
}

// READ ALL
func GetPatients(c *gin.Context) {
	collection := getCollection()

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	patients := []models.Patient{}
	cursor.All(context.Background(), &patients)

	c.Header("Content-Type", "application/json")
	c.JSON(200, patients)
}

// READ ONE
func GetPatient(c *gin.Context) {
	id := c.Param("id")

	collection := getCollection()

	var patient models.Patient
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&patient)

	if err != nil {
		c.JSON(404, gin.H{"error": "Patient not found"})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, patient)
}

// UPDATE
func UpdatePatient(c *gin.Context) {
	id := c.Param("id")
	var body models.Patient
	c.BindJSON(&body)

	collection := getCollection()

	update := bson.M{
		"$set": bson.M{
			"name":         body.Name,
			"age":          body.Age,
			"diagnosis":    body.Diagnosis,
			"phone_number": body.PhoneNumber,
		},
	}

	_, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, update)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, gin.H{"message": "Patient updated"})
}

// DELETE
func DeletePatient(c *gin.Context) {
	id := c.Param("id")

	collection := getCollection()

	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, gin.H{"message": "Patient deleted"})
}
