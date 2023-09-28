package services

import (
	"item/domain/items"
	"item/utils/errors"
	"net/http"
)

var (
	ItemsService itemsServiceInterfce = &itemService{}
)

type itemsServiceInterfce interface {
	Create(item items.Item) (*items.Item, errors.RestErr)
	Get(string) (*items.Item, errors.RestErr)
}
type itemService struct {
}

func (s *itemService) Create(item items.Item) (*items.Item, errors.RestErr) {
	return nil, errors.NewRestError("implement me", http.StatusNotImplemented, "not implemented", nil)
}
func (s *itemService) Get(string) (*items.Item, errors.RestErr) {
	return nil, errors.NewRestError("implement me", http.StatusNotImplemented, "not implemented", nil)
}
