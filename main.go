package main

import (
	"GoCrudApi/database"
	"GoCrudApi/types"

	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

//var users []types.User
//var courses []types.Course

func getUsers(c echo.Context) error {
	var user []types.User
	if err := database.Db.Preload("Courses").Find(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}
func deleteUsers(c echo.Context) error {
	id := c.Param("id")
	if err := database.Db.Delete(&types.User{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, id)
}

func getUsersById(c echo.Context) error {
	id := c.Param("id")
	var user types.User
	if err := database.Db.Preload("Courses").First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}
	return c.JSON(http.StatusOK, user)
}

func createUser(c echo.Context) error {

	var user types.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := database.Db.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, user)

}

func updateUsers(c echo.Context) error {
	id := c.Param("id")
	var user types.User
	if err := database.Db.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := database.Db.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func getCourses(c echo.Context) error {
	var courses []types.Course
	if err := database.Db.Find(&courses).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error fetching courses"})
	}
	return c.JSON(http.StatusOK, courses)
}

func getCourseById(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid course ID"})
	}

	var course types.Course
	if err := database.Db.First(&course, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "course not found"})
	}
	return c.JSON(http.StatusOK, course)
}

func createCourse(c echo.Context) error {
	var course types.Course
	if err := c.Bind(&course); err != nil {
		return err
	}
	if err := database.Db.Create(&course).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to create course"})
	}
	return c.JSON(http.StatusCreated, course)
}

func updateCourse(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid course ID"})
	}

	var course types.Course
	if err := database.Db.First(&course, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "course not found"})
	}

	var updatedCourse types.Course
	if err := c.Bind(&updatedCourse); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid input"})
	}

	course.Title = updatedCourse.Title

	if err := database.Db.Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to update course"})
	}

	return c.JSON(http.StatusOK, course)
}

func deleteCourse(c echo.Context) error {
	id := c.Param("id")
	var course types.Course
	if err := database.Db.First(&course, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Course not found"})
	}
	if err := database.Db.Delete(&course).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error deleting course"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Course deleted"})
}

func main() {
	fmt.Println("Hello World")
	e := echo.New()
	database.Connect()
	//users = append(users, types.User{
	//	Id:   "12",
	//	Name: "Adib",
	//	Age:  22,
	//	Courses: []types.Course{
	//		{
	//			ID:    "1",
	//			Title: "Bangla",
	//		},
	//	},
	//})
	//
	//courses = append(courses, types.Course{
	//	ID:    "2",
	//	Title: "English",
	//})
	e.GET("/users", getUsers)
	e.GET("/users/{id}", getUsersById)
	e.POST("/users", createUser)
	e.PUT("/users/{id}", updateUsers)
	e.DELETE("/users/{id}", deleteUsers)
	e.GET("/courses", getCourses)
	e.GET("/courses/{id}", getCourseById)
	e.POST("/courses", createCourse)
	e.PUT("/courses/{id}", updateCourse)
	e.DELETE("/courses/{id}", deleteCourse)

	fmt.Println("Server started at port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
