package input

import (
	"errors"

	"github.com/shirakiyo/ConveniGacha/internal/usecase"
)

type GetProduct struct {
	Category string
}

func (i GetProduct) Validate() error {
	switch usecase.ProductCategory(i.Category) {
	case "":
		return nil
	case usecase.FoodsPrefix:
		return nil
	case usecase.SweetsPrefix:
		return nil
	case usecase.SnacksPrefix:
		return nil
	}
	return errors.New("カテゴリ名が正しくありません")
}
