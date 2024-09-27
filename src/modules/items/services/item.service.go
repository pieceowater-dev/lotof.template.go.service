package services

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"template/src/modules/items/dto"
	"template/src/modules/items/entities"
)

type ItemService struct {
	DB *gorm.DB
}

func NewItemService(db *gorm.DB) *ItemService {
	return &ItemService{DB: db}
}

func (s *ItemService) Create(dto dto.CreateItemDTO) (*entities.Item, error) {
	item := entities.Item{
		Name:    dto.Name,
		Comment: dto.Comment,
	}

	if err := s.DB.Create(&item).Error; err != nil {
		return nil, err
	}

	s.logAction("createItem", item)

	return &item, nil
}

func (s *ItemService) FindByID(id int) (*entities.Item, error) {
	var item entities.Item
	if err := s.DB.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *ItemService) Update(id int, dto dto.UpdateItemDTO) (*entities.Item, error) {
	var item entities.Item
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
	if err := s.DB.Delete(&entities.Item{}, id).Error; err != nil {
		return err
	}

	s.logAction("deleteItem", id)

	return nil
}

func (s *ItemService) logAction(action string, message any) {
	body, err := json.Marshal(message)
	if err != nil {
		//log.Error(fmt.Errorf("failed to log action: %w", err))
		return
	}

	log.Info().Str("action", action).Bytes("data", body).Msg("Action logged")
}
