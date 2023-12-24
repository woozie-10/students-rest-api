package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/woozie-10/students-rest-api/database"
	"github.com/woozie-10/students-rest-api/entities"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

// GetStudentsByGroupHandler godoc
// @Summary Get students by group
// @Description Retrieve a list of students from the database by group
// @Tags students
// @Accept json
// @Produce json
// @Param group path string true "Group of the students to be retrieved"
// @Success 200 {array} entities.Student "List of students"
// @Failure 400 {object} entities.ErrorResponse "Bad request"
// @Failure 404 {object} entities.ErrorResponse "Group not found"
// @Failure 500 {object} entities.ErrorResponse "Internal server error"
// @Router /students/group/{group} [get]
func GetStudentsByGroupHandler(c *gin.Context) {
	group := c.Param("group")

	var students []entities.Student

	filter := bson.M{"group": group}

	cursor, err := database.StudentsCollection.Find(context.TODO(), filter)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, entities.ErrorResponse{err.Error()})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var student entities.Student
		if err := cursor.Decode(&student); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, entities.ErrorResponse{err.Error()})
			return
		}
		students = append(students, student)
	}

	if err := cursor.Err(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, entities.ErrorResponse{err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, students)
}
