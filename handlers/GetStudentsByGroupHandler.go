package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/woozie-10/students-rest-api/database"
	"github.com/woozie-10/students-rest-api/entities"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func GetStudentsByGroupHandler(c *gin.Context) {
	group := c.Param("group")

	var students []entities.Student

	filter := bson.M{"group": group}

	cursor, err := database.StudentsCollection.Find(context.TODO(), filter)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var student entities.Student
		if err := cursor.Decode(&student); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
			return
		}
		students = append(students, student)
	}

	if err := cursor.Err(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}

	c.JSON(http.StatusOK, students)
}
