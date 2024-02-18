package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sadreddinov/distributed_calculator/orchestrator/pkg/models"
	"github.com/sirupsen/logrus"
)


//  GetComputingResources godoc
// @Summary      Get computing resources
// @Description  Get computing resources
// @Tags         computing resources
// @Accept       json
// @Produce      json
// @Success      200  {array}  models.ComputingResource
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /computing_resources/ [get]
func (h *Handler) GetComputingResources() gin.HandlerFunc {
	return func(c *gin.Context) {
		agents, err := h.services.ComputingResource.GetComputingResources()
		if err != nil {
			logrus.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			return
		}
		c.JSON(http.StatusOK, agents)
	}
}

func (h *Handler) CreateComputingResources() gin.HandlerFunc {
	return func(c *gin.Context) {
		var agent models.ComputingResource
		if err := c.BindJSON(&agent); err != nil {
			logrus.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		uuid, err := h.services.ComputingResource.CreateComputingResource(agent)
		if err != nil {
			logrus.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, uuid)
	}
}

func (h *Handler) UpdateComputingResource() gin.HandlerFunc {
	return func(c *gin.Context) {
		var agent models.ComputingResource
		if err := c.BindJSON(&agent); err != nil {
			logrus.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := h.services.ComputingResource.UpdateComputingResource(agent); err != nil {
			logrus.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}

func (h *Handler) ShutdownComputingResource() gin.HandlerFunc {
	return func(c *gin.Context) {
		var agent models.ComputingResource
		if err := c.BindJSON(&agent); err != nil {
			logrus.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := h.services.ComputingResource.ShutdownComputingResource(agent); err != nil {
			logrus.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}