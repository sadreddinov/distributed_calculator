package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sadreddinov/distributed_calculator/orchestrator/pkg/models"
)

func (h *Handler) OperationTime(c *gin.Context) {
	if _, ok := c.Get("plus"); !ok {
		c.Set("plus", models.Operations.Plus)
	}
	if _, ok := c.Get("minus"); !ok {
		c.Set("minus", models.Operations.Minus)
	}
	if _, ok := c.Get("multiply"); !ok {
		c.Set("multiply", models.Operations.Multiply)
	}
	if _, ok := c.Get("divide"); !ok {
		c.Set("divide", models.Operations.Divide)
	}
}
