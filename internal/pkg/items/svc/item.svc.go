package svc

import (
	"application/internal/pkg/items/dto"
	"application/internal/pkg/items/ent"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type ItemService struct {
	DB *gorm.DB
}

func NewItemService(db *gorm.DB) *ItemService {
	return &ItemService{DB: db}
}

func (s *ItemService) Create(dto dto.CreateItemDTO) (*ent.Item, error) {
	item := ent.Item{
		Name:    dto.Name,
		Comment: dto.Comment,
	}

	if err := s.DB.Create(&item).Error; err != nil {
		return nil, err
	}

	s.logAction("createItem", item)
	return &item, nil
}

func (s *ItemService) FindByID(id int) (*ent.Item, error) {
	var item ent.Item
	if err := s.DB.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *ItemService) FindAll() ([]ent.Item, error) {
	var items []ent.Item
	if err := s.DB.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (s *ItemService) Update(id int, dto dto.UpdateItemDTO) (*ent.Item, error) {
	var item ent.Item
	if err := s.DB.First(&item, id).Error; err != nil {
		return nil, err
	}

	item.Name = dto.Name
	item.Comment = dto.Comment

	if err := s.DB.Save(&item).Error; err != nil {
		return nil, err
	}

	s.logAction("updateItem", item)
	return &item, nil
}

func (s *ItemService) Delete(id int) error {
	if err := s.DB.Delete(&ent.Item{}, id).Error; err != nil {
		return err
	}

	s.logAction("deleteItem", id)
	return nil
}

func (s *ItemService) logAction(action string, data any) {
	body, err := json.Marshal(data)
	if err != nil {
		log.Error().Err(err).Msg("Failed to marshal log data")
		return
	}

	log.Info().Str("action", action).Bytes("data", body).Msg("Action logged")
}
