package handlers

import (
	"context"
	"github.com/woozie-10/students-rest-api/database"
	"github.com/woozie-10/students-rest-api/entities"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func GetStudentsHandler(c *gin.Context) {
	var students []entities.Student
	filter := bson.M{}
	cursor, err := database.StudentsCollection.Find(context.TODO(), filter)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}
	if err := cursor.All(context.TODO(), &students); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, students)
}
