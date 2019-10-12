package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

		group.GET("/:id", taskController.GetByID)
		group.POST("/", taskController.Save)
		group.PUT("/", taskController.Update)
		group.DELETE("/:id", taskController.Delete)
		/*
			group.GET("/period/dtInicio/dataFim", controller.FetchArticle)
			group.GET("/search/dtInicio/dataFim", controller.FetchArticle)
		*/
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

	id := c.Param("id")

	task, err := taskController.TaskService.GetByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, task)
}

func (taskController *TaskController) Update(c *gin.Context) {

	id := c.Param("id")

	task, err := taskController.TaskService.GetByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, task)
}

func (taskController *TaskController) Delete(c *gin.Context) {

	id := c.Param("id")

	task, err := taskController.TaskService.GetByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, task)
}
