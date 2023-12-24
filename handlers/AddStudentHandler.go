package handlers

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/woozie-10/students-rest-api/database"
	_ "github.com/woozie-10/students-rest-api/docs"
	"github.com/woozie-10/students-rest-api/entities"
	"net/http"
)

// AddStudentHandler godoc
// @Summary Add a new student
// @Description Add a new student to the database
// @Tags students
// @Accept json
// @Produce json
// @Param student body entities.Student true "Student object to be added"
// @Success 200 {object} entities.Response"Student has been added"
// @Failure 400 {object} entities.ErrorResponse "Bad request"
// @Failure 500 {object} entities.ErrorResponse "Internal server error"
// @Router /students [post]
func AddStudentHandler(c *gin.Context) {
	var student entities.Student
	if err := json.NewDecoder(c.Request.Body).Decode(&student); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, entities.ErrorResponse{err.Error()})
		return
	}
	_, err := database.StudentsCollection.InsertOne(context.TODO(), student)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, entities.ErrorResponse{err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, entities.Response{"Student has been added"})

}
