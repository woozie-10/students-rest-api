package handlers

import (
	"context"
	"github.com/woozie-10/students-rest-api/database"
	"github.com/woozie-10/students-rest-api/entities"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetStudentByUsernameHandler godoc
// @Summary Get a student by username
// @Description Retrieve a student from the database by username
// @Tags students
// @Accept json
// @Produce json
// @Param username path string true "Username of the student to be retrieved"
// @Success 200 {object} entities.Student "Student information"
// @Failure 400 {object} entities.ErrorResponse "Bad request"
// @Failure 404 {object} entities.ErrorResponse "Student not found"
// @Failure 500 {object} entities.ErrorResponse "Internal server error"
// @Router /students/{username} [get]
func GetStudentByUsernameHandler(c *gin.Context) {
	username := c.Param("username")
	var student entities.Student
	filter := bson.M{"username": username}
	err := database.StudentsCollection.FindOne(context.TODO(), filter).Decode(&student)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, entities.ErrorResponse{err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, student)
}
