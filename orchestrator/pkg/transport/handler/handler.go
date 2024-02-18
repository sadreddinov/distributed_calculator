package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/sadreddinov/distributed_calculator/orchestrator/docs"
	"github.com/sadreddinov/distributed_calculator/orchestrator/pkg/service"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	expressions := router.Group("/expressions", h.OperationTime)
	{
		expressions.GET("/", h.GetExpressions)
		expressions.GET("/:id", h.GetExpression())
		expressions.POST("/", h.CreateExpression())
	}

	operations := router.Group("/operations", h.OperationTime)
	{
		operations.GET("/", h.GetOperations())
		operations.PATCH("/", h.UpdateOperations())
	}
	computing_resources := router.Group("/computing_resources", h.OperationTime)
	{
		computing_resources.GET("/", h.GetComputingResources())
		computing_resources.POST("/", h.CreateComputingResources())
		computing_resources.PATCH("/", h.UpdateComputingResource())
		computing_resources.PATCH("/shutdown", h.ShutdownComputingResource())
	}
	task := router.Group("/task", h.OperationTime)
	{
		task.GET("/:id", h.GetTask())
		task.POST("/", h.PostResult())
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
