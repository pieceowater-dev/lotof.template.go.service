package services

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
	"template/src/modules/items/dto"
	"template/src/modules/items/entities"
	"template/src/utils/logs"
)

type ItemService struct {
	DB       *gorm.DB
	RabbitMQ *amqp.Connection
}

func NewItemService(db *gorm.DB, rabbitMQ *amqp.Connection) *ItemService {
	return &ItemService{DB: db, RabbitMQ: rabbitMQ}
}

func (s *ItemService) Create(dto dto.CreateItemDTO) (*entities.Item, error) {
	item := entities.Item{
		Name:    dto.Name,
		Comment: dto.Comment,
	}

	if err := s.DB.Create(&item).Error; err != nil {
		return nil, err
	}

	if err := s.publishMessage("createItem", item); err != nil {
		log.Error(fmt.Errorf("failed to publish create item message: %w", err), nil)
	}

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

	if err := s.publishMessage("updateItem", item); err != nil {
		log.Error(fmt.Errorf("failed to publish update item message: %w", err), nil)
	}

	return &item, nil
}

func (s *ItemService) Delete(id int) error {
	if err := s.DB.Delete(&entities.Item{}, id).Error; err != nil {
		return err
	}
	if err := s.publishMessage("deleteItem", id); err != nil {
		log.Error(fmt.Errorf("failed to publish delete item message: %w", err), nil)
	}
	return nil
}

func (s *ItemService) publishMessage(routingKey string, message interface{}) error {
	ch, err := s.RabbitMQ.Channel()
	if err != nil {
		return err
	}
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {

		}
	}(ch)

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",         // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
