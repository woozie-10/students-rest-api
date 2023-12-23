package handlers

import (
	"context"
	"encoding/json"
	"github.com/woozie-10/students-rest-api/database"
	"github.com/woozie-10/students-rest-api/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateStudentHandler(c *gin.Context) {
	var newStudent entities.Student
	if err := json.NewDecoder(c.Request.Body).Decode(&newStudent); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}
	username := c.Param("username")
	filter := bson.M{"username": username}
	update := bson.M{"$set": newStudent}
	opts := options.Update().SetUpsert(true)
	_, err := database.StudentsCollection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Student has been updated"})

}
