package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"template/src/modules/items/dto"
	"template/src/modules/items/services"
)

type ItemController struct {
	ItemService *services.ItemService
}

func NewItemController(service *services.ItemService) *ItemController {
	return &ItemController{ItemService: service}
}

func (ctrl *ItemController) CreateItem(c *gin.Context) {
	var createDTO dto.CreateItemDTO
	if err := c.ShouldBindJSON(&createDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item, err := ctrl.ItemService.Create(createDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (ctrl *ItemController) GetItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	item, err := ctrl.ItemService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (ctrl *ItemController) UpdateItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updateDTO dto.UpdateItemDTO
	if err := c.ShouldBindJSON(&updateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item, err := ctrl.ItemService.Update(id, updateDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (ctrl *ItemController) DeleteItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.ItemService.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
}
