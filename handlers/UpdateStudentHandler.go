package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/woozie-10/students-rest-api/database"
	"github.com/woozie-10/students-rest-api/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gin-gonic/gin"
)

// UpdateStudentHandler godoc
// @Summary Update a student
// @Description Update an existing student in the database
// @Tags students
// @Accept json
// @Produce json
// @Param username path string true "Username of the student to be updated"
// @Param student body entities.Student true "Updated student object"
// @Success 200 {object} entities.Response "Student has been updated"
// @Failure 400 {object} entities.ErrorResponse "Bad request"
// @Failure 500 {object} entities.ErrorResponse "Internal server error"
// @Router /students/{username} [patch]
func UpdateStudentHandler(c *gin.Context) {
	var newStudent entities.Student
	if err := json.NewDecoder(c.Request.Body).Decode(&newStudent); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, entities.ErrorResponse{err.Error()})
		return
	}
	username := c.Param("username")
	filter := bson.M{"username": username}
	update := bson.M{"$set": newStudent}
	opts := options.Update().SetUpsert(true)
	_, err := database.StudentsCollection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, entities.ErrorResponse{err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, entities.Response{"Student has been updated"})

}
