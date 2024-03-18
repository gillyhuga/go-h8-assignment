package item_repository

import (
	"assignment-2/entity"
	"assignment-2/pkg/errs"
)

type Repository interface {
	GetItemsByCodes(itemCodes []string) ([]entity.Item, errs.Error)
}
