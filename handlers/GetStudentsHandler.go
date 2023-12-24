package handlers

import (
	"context"
	"github.com/woozie-10/students-rest-api/database"
	"github.com/woozie-10/students-rest-api/entities"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

// GetStudentsHandler godoc
// @Summary Get students
// @Description Retrieve a list of students from the database
// @Tags students
// @Accept json
// @Produce json
// @Success 200 {array} entities.Student "List of students"
// @Failure 400 {object} entities.ErrorResponse "Bad request"
// @Failure 500 {object} entities.ErrorResponse "Internal server error"
// @Router /students [get]
func GetStudentsHandler(c *gin.Context) {
	var students []entities.Student
	filter := bson.M{}
	cursor, err := database.StudentsCollection.Find(context.TODO(), filter)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, entities.ErrorResponse{err.Error()})
		return
	}
	if err := cursor.All(context.TODO(), &students); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, entities.ErrorResponse{err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, students)
}
