package handlers

import (
	"context"
	"github.com/woozie-10/students-rest-api/database"

	"go.mongodb.org/mongo-driver/bson"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteStudentByUsernameHandler(c *gin.Context) {
	username := c.Param("username")

	filter := bson.M{"username": username}
	_, err := database.StudentsCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"result": "Student has been deleted"})

}
