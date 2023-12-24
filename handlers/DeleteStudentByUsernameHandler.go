package handlers

import (
	"context"
	"github.com/woozie-10/students-rest-api/database"
	"github.com/woozie-10/students-rest-api/entities"

	"go.mongodb.org/mongo-driver/bson"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteStudentByUsernameHandler godoc
// @Summary Delete a student by username
// @Description Delete a student from the database by username
// @Tags students
// @Accept json
// @Produce json
// @Param username path string true "Username of the student to be deleted"
// @Success 200 {object} entities.Response "Student has been deleted"
// @Failure 400 {object} entities.ErrorResponse "Bad request"
// @Failure 500 {object} entities.ErrorResponse "Internal server error"
// @Router /students/{username} [delete]
func DeleteStudentByUsernameHandler(c *gin.Context) {
	username := c.Param("username")

	filter := bson.M{"username": username}
	_, err := database.StudentsCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, entities.ErrorResponse{err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, entities.Response{"Student has been deleted"})

}
