package service

import (
	"navigate/dto"
	"navigate/entity"
	"navigate/repository"
)

type ResourceTypeService interface {
	ResourceList(page int, limit int) ([]*entity.ResourceType, int64)
	CreateResourceType(resourceType dto.ResourceTypeCreateDto) entity.ResourceType
	UpdateResourceType(id int, resourceType dto.ResourceTypeUpdateDto) entity.ResourceType
	FindResourceType(id int) entity.ResourceType
	DeleteResourceType(id int) entity.ResourceType
}

type resourceTypeService struct {
	repo repository.ResourceTypeRepository
}

func NewResourceTypeRepository(repo repository.ResourceTypeRepository) ResourceTypeService {
	return &resourceTypeService{
		repo: repo,
	}
}

func (r *resourceTypeService) ResourceList(page int, limit int) ([]*entity.ResourceType, int64) {
	return r.repo.ResourceList(page, limit)
}

func (r *resourceTypeService) CreateResourceType(resourceType dto.ResourceTypeCreateDto) entity.ResourceType {
	return r.repo.CreateResourceType(&entity.ResourceType{

		Name: resourceType.Name,
		Icon: resourceType.Icon,
		Url:  resourceType.Url,
	})

}

func (r *resourceTypeService) UpdateResourceType(id int, resourceType dto.ResourceTypeUpdateDto) entity.ResourceType {
	return r.repo.UpdateResourceType(id, &entity.ResourceType{
		Name: resourceType.Name,
		Icon: resourceType.Icon,
		Url:  resourceType.Url,
	})
}

func (r *resourceTypeService) FindResourceType(id int) entity.ResourceType {
	return r.repo.FindResourceType(id)

}

func (r *resourceTypeService) DeleteResourceType(id int) entity.ResourceType {
	return r.repo.DeleteResourceType(id)
}
