package repository

import "github.com/mhaatha/go-learn/go-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
