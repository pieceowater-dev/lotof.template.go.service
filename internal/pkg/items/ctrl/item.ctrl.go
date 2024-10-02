package ctrl

import (
	. "application/internal/core/utils/common/dto"
	"application/internal/pkg/items/dto"
	"application/internal/pkg/items/svc"
	g "github.com/pieceowater-dev/lotof.lib.gossiper"
	"log"
	"net/http"
)

type ItemController struct {
	ItemService *svc.ItemService
}

func NewItemController(service *svc.ItemService) *ItemController {
	return &ItemController{ItemService: service}
}

// CreateItem handles the creation of an item
func (ctrl *ItemController) CreateItem(data any) any {
	var createDTO dto.CreateItemDTO
	err := g.Satisfies(data, &createDTO)
	if err != nil {
		log.Println("Error validating input for CreateItem:", err)
		return g.NewServiceError("Invalid input", http.StatusBadRequest).GetError()
	}

	item, err := ctrl.ItemService.Create(createDTO)
	if err != nil {
		log.Println("Error creating item:", err)
		return g.NewServiceError(err.Error(), http.StatusInternalServerError).GetError()
	}

	return item
}

// GetItems handles retrieving items with pagination and filtering
func (ctrl *ItemController) GetItems(data any) any {
	filter := g.NewFilter[dto.UpdateItemDTO]()

	err := g.Satisfies(data, &filter)
	if err != nil {
		log.Println("Error validating input for GetItems:", err)
		return g.NewServiceError("Invalid input").GetError()
	}

	// Pass filter to FindAll
	return ctrl.ItemService.FindAll(filter)
}

// GetItem retrieves a single item by ID
func (ctrl *ItemController) GetItem(data any) any {
	var idDTO ID
	err := g.Satisfies(data, &idDTO)
	if err != nil {
		log.Println("Error validating input for GetItem:", err)
		return g.NewServiceError("Invalid input").GetError()
	}

	item, err := ctrl.ItemService.FindByID(idDTO.ID)
	if err != nil {
		log.Println("Error retrieving item:", err)
		return g.NewServiceError(err.Error()).GetError()
	}

	return item
}

// UpdateItem handles updating an item
func (ctrl *ItemController) UpdateItem(data any) any {
	updateDTO := dto.UpdateItemDTO{}
	err := g.Satisfies(data, &updateDTO)
	if err != nil {
		log.Println("Error validating input for UpdateItem:", err)
		return g.NewServiceError("Invalid input").GetError()
	}

	item, err := ctrl.ItemService.Update(updateDTO.ID, updateDTO)
	if err != nil {
		log.Println("Error updating item:", err)
		return g.NewServiceError(err.Error()).GetError()
	}
	return item
}
