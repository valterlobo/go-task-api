package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"lobo.tech/task/model"
	"lobo.tech/task/service"
)

type TaskController struct {
	TaskService *service.TaskService
}

func NewTaskController(router *gin.Engine) {

	taskController := &TaskController{}
	taskService := service.NewTaskService()
	taskController.TaskService = taskService

	group := router.Group("/task")
	{

		group.GET("/id/:id", taskController.GetByID)
		group.GET("/period/:start/:finish", taskController.GetByPeriod)
		group.GET("/day/:date", taskController.GetByDate)
		group.POST("/", taskController.Save)
		group.PUT("/", taskController.Update)
		group.DELETE("/:id", taskController.Delete)
	}

}

func (taskController *TaskController) GetByID(c *gin.Context) {

	id := c.Param("id")

	task, err := taskController.TaskService.GetByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, task)
}

func (taskController *TaskController) Save(c *gin.Context) {

	taskParam := model.Task{}

	c.ShouldBind(&taskParam)

	task, err := taskController.TaskService.Save(taskParam)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, task)
	}

}

func (taskController *TaskController) Update(c *gin.Context) {

	taskParam := model.Task{}
	c.ShouldBind(&taskParam)
	fmt.Println(taskParam)
	task, err := taskController.TaskService.Update(taskParam)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, task)
	}
}

func (taskController *TaskController) Delete(c *gin.Context) {

	id := c.Param("id")

	ok, err := taskController.TaskService.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, ok)
}

func (taskController *TaskController) GetByPeriod(c *gin.Context) {

	start := c.Param("start")
	finsih := c.Param("finish")

	tasks, err := taskController.TaskService.GetByPeriod(start, finsih)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, tasks)
}

func (taskController *TaskController) GetByDate(c *gin.Context) {

	date := c.Param("date")

	tasks, err := taskController.TaskService.GetByPeriod(date, date)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, tasks)
}
