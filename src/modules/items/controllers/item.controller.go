package controllers

import (
	"log"
	"template/src/modules/items/dto"
	"template/src/modules/items/services"
)

type ItemController struct {
	ItemService *services.ItemService
}

func NewItemController(service *services.ItemService) *ItemController {
	return &ItemController{ItemService: service}
}

func (ctrl *ItemController) CreateItem(data any) any {
	createDTO, ok := data.(dto.CreateItemDTO)
	if !ok {
		log.Println("Invalid data format for CreateItem")
		return map[string]string{"error": "Invalid data format"}
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
	id, ok := data.(float64) // Assuming ID is a number; adjust if needed
	if !ok {
		log.Println("Invalid data format for GetItem")
		return map[string]string{"error": "Invalid data format"}
	}

	item, err := ctrl.ItemService.FindByID(int(id))
	if err != nil {
		log.Println("Error retrieving item:", err)
		return map[string]string{"error": err.Error()}
	}
	return item
}

func (ctrl *ItemController) UpdateItem(data any) any {
	// Assuming data contains both ID and DTO in a map
	dataMap, ok := data.(map[string]any)
	if !ok {
		log.Println("Invalid data format for UpdateItem")
		return map[string]string{"error": "Invalid data format"}
	}

	id, idOk := dataMap["id"].(float64)               // Adjust as necessary
	dtoData, dtoOk := dataMap["dto"].(map[string]any) // Assuming DTO is a map

	if !idOk || !dtoOk {
		log.Println("Invalid data format for UpdateItem")
		return map[string]string{"error": "Invalid data format"}
	}

	updateDTO := dto.UpdateItemDTO{
		Name:    dtoData["name"].(string),    // Adjust field extraction
		Comment: dtoData["comment"].(string), // Adjust field extraction
	}

	item, err := ctrl.ItemService.Update(int(id), updateDTO)
	if err != nil {
		log.Println("Error updating item:", err)
		return map[string]string{"error": err.Error()}
	}
	return item
}
