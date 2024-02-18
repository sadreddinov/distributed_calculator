package service

import (
	"github.com/google/uuid"
	"github.com/sadreddinov/distributed_calculator/orchestrator/pkg/models"
	"github.com/sadreddinov/distributed_calculator/orchestrator/pkg/storage/psql"
)

type ComputingResourceService struct {
	repo psql.ComputingResource
}

func NewComputingResourceService(repo psql.ComputingResource) *ComputingResourceService {
	return &ComputingResourceService{repo: repo}
}

func (s *ComputingResourceService) CreateComputingResource(agent models.ComputingResource) (uuid.UUID, error) {
	return s.repo.CreateComputingResource(agent)
}

func (s *ComputingResourceService) GetComputingResources() ([]models.ComputingResource, error) {
	return s.repo.GetComputingResources()
}

func (s *ComputingResourceService) UpdateComputingResource(agent models.ComputingResource) (error) {
	return s.repo.UpdateComputingResource(agent)
}

func (s *ComputingResourceService) ShutdownComputingResource(agent models.ComputingResource) (error) {
	return s.repo.ShutdownComputingResource(agent)
}

