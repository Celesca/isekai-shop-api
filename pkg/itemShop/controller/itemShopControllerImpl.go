package controller

import (
	"net/http"

	_itemShopService "github.com/Celesca/isekai-shop-api/pkg/itemShop/service"
	"github.com/labstack/echo/v4"
)

type itemShopControllerImpl struct {
	itemShopService _itemShopService.ItemShopService
}

func NewItemShopControllerImpl(itemShopService _itemShopService.ItemShopService) ItemShopController {
	return &itemShopControllerImpl{itemShopService}
}

func (c *itemShopControllerImpl) Listing(pctx echo.Context) error {
	itemModelList, err := c.itemShopService.Listing()
	if err != nil {
		return pctx.String(http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusOK, itemModelList)
}
