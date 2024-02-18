package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sadreddinov/distributed_calculator/orchestrator/pkg/models"
)

// @Summary      Get time of operations
// @Description  Get time of operations in seconds
// @Tags         operations
// @Accept       json
// @Produce      json
// @Success      200  {array}  models.Operation
// @Router       /operations/ [get]

func (h *Handler) GetOperations() gin.HandlerFunc {
	return func(c *gin.Context) {
		var operations models.Operation
		operations.Plus = c.GetString("plus")
		operations.Minus = c.GetString("minus")
		operations.Multiply = c.GetString("multiply")
		operations.Divide = c.GetString("divide")
		c.JSON(http.StatusOK, operations)
	}
}

// @Summary      Update operation time
// @Description  Update operation time
// @Tags         operations
// @Accept       json
// @Produce      json
// @Param        operations   body      models.Operation  true  "Time of operations"
// @Success      200  {object}  models.Operation
// @Failure      400  {object}  map{string}string
// @Failure      404  {object}  map{string}string
// @Failure      500  {object}  map{string}string
// @Router       /expression/ [patch]

func (h *Handler) UpdateOperations() gin.HandlerFunc {
	return func(c *gin.Context) {
		var operations models.Operation
		if err := c.BindJSON(&operations); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		fmt.Println(operations)
		if operations.Plus != "" && operations.Plus != c.GetString("plus") {
			models.Operations.Plus = operations.Plus
		}
		if operations.Minus != "" && operations.Minus != c.GetString("minus") {
			models.Operations.Minus = operations.Minus
		}
		if operations.Multiply != "" && operations.Multiply != c.GetString("multiply") {
			models.Operations.Multiply = operations.Multiply
		}
		if operations.Divide != "" && operations.Divide != c.GetString("divide") {
			models.Operations.Divide = operations.Divide
		}
		c.JSON(http.StatusOK, models.Operations)
	}
}
