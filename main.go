package main

import (
	"GoCrudApi/types"

	"fmt"
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
	"strconv"
)

var users []types.User
var courses []types.Course

func getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}
func deleteUsers(c echo.Context) error {
	id := c.Param("id")
	for index, item := range users {
		if item.Id == id {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	return c.JSON(http.StatusOK, users)
}

func getUsersById(c echo.Context) error {
	id := c.Param("id")

	for _, item := range users {
		if item.Id == id {

			return c.JSON(http.StatusOK, item)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "user not found"})
}

func createUser(c echo.Context) error {

	var user types.User
	if err := c.Bind(&user); err != nil {
		return err
	}
	user.Id = strconv.Itoa(rand.Intn(10000000))
	users = append(users, user)
	return c.JSON(http.StatusOK, user)

}

func updateUsers(c echo.Context) error {
	id := c.Param("id")
	for index, item := range users {
		if item.Id == id {
			users = append(users[:index], users[index+1:]...)
			var updatedUser types.User
			if err := c.Bind(&updatedUser); err != nil {
				return err
			}
			updatedUser.Id = id
			users = append(users, updatedUser)
			return c.JSON(http.StatusOK, updatedUser)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "user not found"})
}

func getCourses(c echo.Context) error {
	return c.JSON(http.StatusOK, courses)
}
func deleteCourse(c echo.Context) error {
	id := c.Param("id")
	for i, course := range courses {
		if course.ID == id {
			courses = append(courses[:i], courses[i+1:]...)
			break
		}
	}
	return c.JSON(http.StatusOK, courses)
}

func getCourseById(c echo.Context) error {
	id := c.Param("id")
	for _, course := range courses {
		if course.ID == id {
			return c.JSON(http.StatusOK, course)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "course not found"})
}

func createCourse(c echo.Context) error {
	var course types.Course
	if err := c.Bind(&course); err != nil {
		return err
	}
	course.ID = strconv.Itoa(rand.Intn(10000000))
	courses = append(courses, course)
	return c.JSON(http.StatusCreated, course)
}

func updateCourse(c echo.Context) error {
	id := c.Param("id")
	for i, course := range courses {
		if course.ID == id {
			courses = append(courses[:i], courses[i+1:]...)
			var updatedCourse types.Course
			if err := c.Bind(&updatedCourse); err != nil {
				return err
			}
			updatedCourse.ID = id
			courses = append(courses, updatedCourse)
			return c.JSON(http.StatusOK, updatedCourse)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "course not found"})
}

func main() {
	fmt.Println("Hello World")
	e := echo.New()
	users = append(users, types.User{
		Id:   "12",
		Name: "Adib",
		Age:  22,
		Courses: []types.Course{
			{
				ID:    "1",
				Title: "Bangla",
			},
		},
	})

	courses = append(courses, types.Course{
		ID:    "2",
		Title: "English",
	})
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
