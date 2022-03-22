package repository

import (
	"navigate/entity"

	"gorm.io/gorm"
)

type ResourceRepository interface {
	// 增删改查
	Create(resource *entity.Resource) entity.Resource
	Update(id int, resource *entity.Resource) entity.Resource
	Delete(id int) entity.Resource
	Get(id int) entity.Resource
	GetAll(page int, limit int) ([]*entity.Resource, int64)
}

type resourceRepository struct {
	db *gorm.DB
}

func NewResourceRepository(db *gorm.DB) ResourceRepository {
	return &resourceRepository{
		db: db,
	}
}

func (r *resourceRepository) Create(resource *entity.Resource) entity.Resource {
	r.db.Create(&resource)
	return *resource
}

func (r *resourceRepository) Update(id int, resource *entity.Resource) entity.Resource {
	r.db.Model(&entity.Resource{}).Where("id = ?", id).Updates(resource)
	return *resource
}

func (r *resourceRepository) Delete(id int) entity.Resource {
	var resource entity.Resource
	r.db.First(&resource, id)
	r.db.Delete(&resource)
	return resource
}

func (r *resourceRepository) Get(id int) entity.Resource {
	var resource entity.Resource
	r.db.First(&resource, id)
	return resource
}

func (r *resourceRepository) GetAll(page int, limit int) ([]*entity.Resource, int64) {

	var resources []*entity.Resource
	var count int64
	r.db.Find(&resources).Count(&count)
	r.db.Offset((page - 1) * limit).Limit(limit).Find(&resources)
	return resources, count
}
