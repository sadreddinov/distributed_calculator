package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sadreddinov/distributed_calculator/orchestrator/pkg/models"
	"github.com/sirupsen/logrus"
)

// GetExpressions godoc
// @Summary      Get expressions
// @Security ApiKeyAuth
// @Description  Get expressions
// @Tags         expressions
// @Accept       json
// @Produce      json
// @Param        page   query      int  true  "Page num"
// @Param        recordPerPage   query      int  true  "Num of record per page"
// @Success      200  {array}  models.ExpressionToRead
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /expressions/ [get]
func (h *Handler) GetExpressions(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
	if err != nil || recordPerPage < 1 {
		recordPerPage = 10
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}

	startIndex := (page - 1) * recordPerPage
	startIndex, err = strconv.Atoi(c.Query("startIndex"))

	expressions, err := h.services.Expression.GetExpressions(startIndex, recordPerPage, userId)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, expressions)
}

// GetExpression godoc
// @Summary      Get expression
// @Security ApiKeyAuth
// @Description  Get expression by id
// @Tags         expressions
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Expression ID"
// @Success      200  {object}  models.ExpressionToRead
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /expressions/{id} [get]
func (h *Handler) GetExpression() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := getUserId(c)
		if err != nil {
			return
		}

		id := c.Param("id")
		uuid, err := uuid.Parse(id)
		if err != nil {
			logrus.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		expression, err := h.services.Expression.GetExpression(uuid, userId)
		if err != nil {
			logrus.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the expression item"})
			return
		}
		c.JSON(http.StatusOK, expression)
	}
}

// AddExpression godoc
// @Summary      Add expression
// @Security ApiKeyAuth
// @Description  Add new expression
// @Tags         expressions
// @Accept       json
// @Produce      json
// @Param        expression   body      models.ExpressionFromUser  true  "Expression info"
// @Success      200  {object}  uuid.UUID
// @Failure      400  {object}  map[string]string
// @Failure      404  {object} map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /expressions/ [post]
func (h *Handler) CreateExpression() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := getUserId(c)
		if err != nil {
			return
		}

		var expression models.ExpressionFromUser
		if err := c.BindJSON(&expression); err != nil {
			logrus.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id, err := h.services.Expression.CreateExpression(expression, userId)
		if err != nil {
			logrus.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, id)
	}
}
