package category

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type CategoryServiceMock struct {
	mock.Mock
}

func (m *CategoryServiceMock) create(ctx context.Context, c createRequest) (str string, err error) {
	args := m.Called(ctx, c)
	return str, args.Error(0)
}

func (m *CategoryServiceMock) list(ctx context.Context) (response listResponse, err error) {
	args := m.Called(ctx)
	response, _ = args.Get(0).(listResponse)
	return response, args.Error(1)
}

func (m *CategoryServiceMock) update(ctx context.Context, c updateRequest) (err error) {
	args := m.Called(ctx, c)
	return args.Error(0)
}

func (m *CategoryServiceMock) findByID(ctx context.Context, id string) (response findByIDResponse, err error) {
	args := m.Called(ctx)
	response, _ = args.Get(0).(findByIDResponse)
	return response, args.Error(1)
}

func (m *CategoryServiceMock) deleteByID(ctx context.Context, id string) (err error) {
	args := m.Called(ctx, id)
	return args.Error(0)
}
