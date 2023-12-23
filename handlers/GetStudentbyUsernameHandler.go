package handlers

import (
	"context"
	"github.com/woozie-10/students-rest-api/database"
	"github.com/woozie-10/students-rest-api/entities"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStudentbyUsernameHandler(c *gin.Context) {
	username := c.Param("username")
	var student entities.Student
	filter := bson.M{"username": username}
	err := database.StudentsCollection.FindOne(context.TODO(), filter).Decode(&student)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, student)
}
