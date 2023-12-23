package handlers

import (
	"context"
	"encoding/json"
	"github.com/woozie-10/students-rest-api/database"
	"github.com/woozie-10/students-rest-api/entities"

	"github.com/gin-gonic/gin"
	"net/http"
)

func AddStudentHandler(c *gin.Context) {
	var student entities.Student
	if err := json.NewDecoder(c.Request.Body).Decode(&student); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}
	_, err := database.StudentsCollection.InsertOne(context.TODO(), student)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"result": "Student has been added"})

}
