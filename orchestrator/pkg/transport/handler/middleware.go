package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sadreddinov/distributed_calculator/orchestrator/pkg/models"
	"github.com/sirupsen/logrus"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "user_id"
)

func (h *Handler) userIndentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		logrus.Error("empty auth header")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "empty auth header"})
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		logrus.Error("invalid auth header")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid auth header"})
		return
	}
	userId, err := h.services.ParseToken(headerParts[1])
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	fmt.Println(id)
	if !ok {
		logrus.Error("user id not found")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user id not found"})
		return 0, errors.New("user id not found")
	}
	idInt, ok := id.(int)
	if !ok {
		logrus.Error("user id is of invalid type")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user id is of invalid type"})
		return 0, errors.New("user id not found")
	}
	return idInt, nil
}

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
