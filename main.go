package main

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
)

type Task struct {
    ID        int    `json:"id"`
    Name      string `json:"name"`
    Completed bool   `json:"completed"`
}

var tasks []Task

func main() {
	gin.SetMode(gin.ReleaseMode)
    router := gin.Default()

    router.GET("/tasks", getTasks)
    router.POST("/tasks", addTask)
    router.PUT("/tasks/:id", updateTask)

    router.StaticFile("/", "./frontend/index.html")
    router.Static("/static", "./frontend/static")

    router.Run(":8080")
}

func getTasks(c *gin.Context) {
    c.JSON(http.StatusOK, tasks)
}

func addTask(c *gin.Context) {
    var task Task
    if err := c.BindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    task.ID = len(tasks) + 1
    tasks = append(tasks, task)
    c.JSON(http.StatusCreated, task)
}

func updateTask(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }
    for i, task := range tasks {
        if id == task.ID {
            tasks[i].Completed = !task.Completed
            c.Status(http.StatusOK)
            return
        }
    }
    c.Status(http.StatusNotFound)
}