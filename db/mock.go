package db

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type StorerMock struct {
	mock.Mock
}

func (m *StorerMock) CreateCategory(ctx context.Context, category *Category) (str string, err error) {
	args := m.Called(ctx, category)
	return args.String(0), args.Error(1)
}

func (m *StorerMock) ListCategories(ctx context.Context) (categories []Category, err error) {
	args := m.Called(ctx)
	categories, _ = args.Get(0).([]Category)
	return categories, args.Error(1)
}

func (m *StorerMock) FindCategoryByID(ctx context.Context, id string) (category Category, err error) {
	args := m.Called(ctx, id)
	category, _ = args.Get(0).(Category)
	return category, args.Error(1)
}

func (m *StorerMock) DeleteCategoryByID(ctx context.Context, id string) (err error) {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *StorerMock) UpdateCategory(ctx context.Context, category *Category) (err error) {
	args := m.Called(ctx, category)
	return args.Error(0)
}
