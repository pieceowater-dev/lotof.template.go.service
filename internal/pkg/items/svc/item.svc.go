package svc

import (
	"application/internal/pkg/items/dto"
	"application/internal/pkg/items/ent"
	g "github.com/pieceowater-dev/lotof.lib.gossiper"
	t "github.com/pieceowater-dev/lotof.lib.gossiper/types"
	"gorm.io/gorm"
	"log"
)

type ItemService struct {
	DB *gorm.DB
}

func NewItemService(db *gorm.DB) *ItemService {
	return &ItemService{DB: db}
}

// Create a new item
func (s *ItemService) Create(dto dto.CreateItemDTO) (*ent.Item, error) {
	item := ent.Item{
		Name:    dto.Name,
		Comment: *dto.Comment,
	}

	if err := s.DB.Create(&item).Error; err != nil {
		return nil, err
	}

	return &item, nil
}

// FindAll retrieves items based on the filter with pagination and search
func (s *ItemService) FindAll(filter t.DefaultFilter[dto.UpdateItemDTO]) g.PaginatedEntity[ent.Item] {
	var items []ent.Item
	var count int64

	log.Println(filter.Pagination)

	query := s.DB.Model(&ent.Item{})

	// Apply search filter
	if filter.Search != "" {
		query = query.Where("name ILIKE ?", "%"+filter.Search+"%")
	}

	// Count total items (without pagination)
	err := query.Count(&count).Error
	if err != nil {
		return g.PaginatedEntity[ent.Item]{}
	}

	// Apply pagination
	offset := (filter.Pagination.Page - 1) * int(filter.Pagination.Length)
	query = query.Offset(offset).Limit(int(filter.Pagination.Length))

	// Get the filtered items
	err = query.Find(&items).Error
	if err != nil {
		return g.PaginatedEntity[ent.Item]{}
	}

	return g.ToPaginated(items, int(count))
}

// FindByID retrieves a single item by ID
func (s *ItemService) FindByID(id int) (*ent.Item, error) {
	var item ent.Item
	if err := s.DB.First(&item, id).Error; err != nil {
		return nil, err
	}

	return &item, nil
}

// Update an existing item
func (s *ItemService) Update(id int, dto dto.UpdateItemDTO) (*ent.Item, error) {
	var item ent.Item
	if err := s.DB.First(&item, id).Error; err != nil {
		return nil, err
	}

	item.Name = dto.Name
	if dto.Comment != nil && *dto.Comment != "" {
		item.Comment = *dto.Comment
	}

	if err := s.DB.Save(&item).Error; err != nil {
		return nil, err
	}

	return &item, nil
}
