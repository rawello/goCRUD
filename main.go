package main

import (
	"awesomeProject/models"
	"log"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/students", getStudents)
	router.GET("/students/:id", getStudent)
	router.POST("/students/add", addStudent)
	router.DELETE("/students/del/:id", deleteStudent)
	router.POST("/login/students/", autorizeStudent)
	router.Run("localhost:8080")
}

func getStudents(c *gin.Context) {

	students := models.GetStudents()

	if students == nil || len(students) == 0 {

		c.AbortWithStatus(http.StatusNotFound)

	} else {

		c.IndentedJSON(http.StatusOK, students)

	}
}

func getStudent(c *gin.Context) {

	id := c.Param("id")

	student := models.GetStudent(id)

	if student == nil {
		c.AbortWithStatus(http.StatusNotFound)

	} else {

		c.IndentedJSON(http.StatusOK, student)

	}
}

func addStudent(c *gin.Context) {
	var stud models.Student
	c.Bind(&stud)

	log.Println(stud)

	if stud.Name != "0" && stud.Age != 0 && stud.Country != "0" && stud.City != "0" {
		models.AddStudent(stud.Name, stud.Age, stud.Country, stud.City)
	} else {
		c.JSON(400, gin.H{"error": " field empty"})
	}
}

func deleteStudent(c *gin.Context) {
	id := c.Param("id")
	x, _ := strconv.Atoi(id)
	if x != 0 {
		models.DeleteStudent(x)
		c.JSON(200, gin.H{"success": "Id " + id + " deleted"})
	} else {
		c.JSON(404, gin.H{"error": " User not found"})
	}
}

func autorizeStudent(c *gin.Context) {
	name := c.Param("name")
	if name != "0" {
		models.AutorizeStudent(name)
		c.JSON(200, gin.H{"ok": "ok"})
	} else {
		c.JSON(404, gin.H{"error": " User not found"})
	}
}
