package service

import (
	"navigate/dto"
	"navigate/entity"
	"navigate/repository"
)

type ResourceService interface {
	// 增删改查
	Create(resource dto.ResourceCreateDto) entity.Resource
	Update(id int, resource dto.ResourceUpdateDto) entity.Resource
	Delete(id int) entity.Resource
	Get(id int) entity.Resource
	GetAll(page int, limit int) ([]*entity.Resource, int64)
	GetByType(typeId int, page int, limit int) ([]*entity.Resource, int64)
}

type resourceService struct {
	repo repository.ResourceRepository
}

func NewResourceService(repo repository.ResourceRepository) ResourceService {
	return &resourceService{
		repo: repo,
	}
}

func (r *resourceService) Create(resource dto.ResourceCreateDto) entity.Resource {
	return r.repo.Create(&entity.Resource{
		Title:     resource.Title,
		SubTitle:  resource.SubTitle,
		Recommend: resource.Recommend,
		Image:     resource.Image,
		Rid:       resource.Rid,
	})
}

func (r *resourceService) Update(id int, resource dto.ResourceUpdateDto) entity.Resource {
	return r.repo.Update(id, &entity.Resource{
		Title:     resource.Title,
		SubTitle:  resource.SubTitle,
		Recommend: resource.Recommend,
		Image:     resource.Image,
	})
}

func (r *resourceService) Get(id int) entity.Resource {
	return r.repo.Get(id)
}

func (r *resourceService) GetAll(page int, limit int) ([]*entity.Resource, int64) {

	return r.repo.GetAll(page, limit)
}

func (r *resourceService) Delete(id int) entity.Resource {
	return r.repo.Delete(id)
}

func (r *resourceService) GetByType(typeId int, page int, limit int) ([]*entity.Resource, int64) {

	return r.repo.GetByType(typeId, page, limit)
}
