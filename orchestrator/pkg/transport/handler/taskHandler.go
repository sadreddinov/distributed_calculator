package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sadreddinov/distributed_calculator/orchestrator/pkg/models"
	"github.com/sirupsen/logrus"
)

func (h *Handler) GetTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		expression, err := h.services.Task.GetTask(id)
		if err != nil  {
			logrus.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the expression item"})
			return
		}
		c.Header("plus", c.GetString("plus"))
		c.Header("minus", c.GetString("minus"))
		c.Header("multiply", c.GetString("multiply"))
		c.Header("divide", c.GetString("divide"))
		c.JSON(http.StatusOK,expression)
	}
}

func (h *Handler) PostResult() gin.HandlerFunc {
	return func(c *gin.Context) {
		var expression models.Expression
		if err := c.BindJSON(&expression); err != nil {
			logrus.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := h.services.Task.PostResult(expression)
		if err != nil {
			logrus.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
}