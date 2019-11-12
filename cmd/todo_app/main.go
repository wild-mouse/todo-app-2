package main

import (
	"net/http"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/tasks", getTasks)
	e.GET("/tasks/:id", getTask)
	e.POST("/tasks", saveTask)
	e.PUT("/tasks/:id", updateTask)
	e.DELETE("/tasks/:id", deleteTask)

	e.Logger.Fatal(e.Start(":1323"))
}

type Task struct {
	Name string `json:"name" form:"name" query:"name"`
	Owner string `json:"owner" form:"owner" query:"owner"`
}

func getTasks(c echo.Context) error {
	return c.String(http.StatusOK, "This endpoint should returns tasks.")
}

func getTask(c echo.Context) error {
	return c.String(http.StatusOK, "This endpoint should returns single task.")
}

func saveTask(c echo.Context) error {
	t := new(Task)
	if err := c.Bind(t); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, t)
}

func updateTask(c echo.Context) error {
	return c.String(http.StatusOK, "This endpoint should update single task.")
}

func deleteTask(c echo.Context) error {
	return c.String(http.StatusOK, "This endpoint should delete single task.")
}
