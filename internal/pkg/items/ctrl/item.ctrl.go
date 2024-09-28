package ctrl

import (
	. "application/internal/core/utils/common/dto"
	"application/internal/pkg/items/dto"
	"application/internal/pkg/items/svc"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper"
	"log"
)

type ItemController struct {
	ItemService *svc.ItemService
}

func NewItemController(service *svc.ItemService) *ItemController {
	return &ItemController{ItemService: service}
}

func (ctrl *ItemController) CreateItem(data any) any {
	var createDTO dto.CreateItemDTO
	err := gossiper.Satisfies(data, &createDTO)
	if err != nil {
		log.Println("Error validating input for CreateItem:", err)
		return map[string]string{"error": "Invalid input"}
	}

	item, err := ctrl.ItemService.Create(createDTO)
	if err != nil {
		log.Println("Error creating item:", err)
		return map[string]string{"error": err.Error()}
	}

	return item
}

func (ctrl *ItemController) GetItems(_ any) any {
	items, err := ctrl.ItemService.FindAll()
	if err != nil {
		log.Println("Error retrieving items:", err)
		return map[string]string{"error": err.Error()}
	}
	return items
}

func (ctrl *ItemController) GetItem(data any) any {
	var idDTO ID
	err := gossiper.Satisfies(data, &idDTO)
	if err != nil {
		log.Println("Error validating input for GetItem:", err)
		return map[string]string{"error": "Invalid input"}
	}

	item, err := ctrl.ItemService.FindByID(idDTO.ID)
	if err != nil {
		log.Println("Error retrieving item:", err)
		return map[string]string{"error": err.Error()}
	}

	return item
}

func (ctrl *ItemController) UpdateItem(data any) any {
	updateDTO := dto.UpdateItemDTO{}
	err := gossiper.Satisfies(data, &updateDTO)
	if err != nil {
		log.Println("Error validating input for UpdateItem:", err)
		return map[string]string{"error": "Invalid input"}
	}

	item, err := ctrl.ItemService.Update(updateDTO.ID, updateDTO)
	if err != nil {
		log.Println("Error updating item:", err)
		return map[string]string{"error": err.Error()}
	}
	return item
}
