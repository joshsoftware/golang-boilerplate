package category

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/joshsoftware/golang-boilerplate/app"
	"github.com/joshsoftware/golang-boilerplate/db"
	"github.com/stretchr/testify/assert"
)

func TestSuccessfullCreateService(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var tests = []struct {
		contx    context.Context
		req      createRequest
		expected error
	}{
		{ctx, createRequest{Name: "Sports"}, nil},
		{ctx, createRequest{Name: "Reading"}, nil},
	}
	for _, test := range tests {
		sm.On("CreateCategory", test.contx, mock.Anything).Return("", nil)
		_, err := cs.create(test.contx, test.req)
		assert.Equal(err, test.expected)
		sm.AssertExpectations(t)
	}
}

func TestCreateServiceWhenEmptyName(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		req      createRequest
		expected error
	}{
		ctx,
		createRequest{Name: ""},
		errEmptyName,
	}
	_, err := cs.create(test.contx, test.req)
	assert.Equal(err, test.expected)
	sm.AssertExpectations(t)
}

func TestCreateServiceWhenInternalError(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		req      createRequest
		expected error
	}{
		ctx,
		createRequest{Name: "Sports"},
		errors.New("Internal Error"),
	}

	sm.On("CreateCategory", test.contx, mock.Anything).Return("", errors.New("Internal Error"))
	_, err := cs.create(test.contx, test.req)
	assert.Equal(err, test.expected)
	sm.AssertExpectations(t)
}

func TestSuccessfullListService(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		expected error
	}{ctx, nil}
	sm.On("ListCategories", test.contx).Return(mock.Anything, nil)
	_, err := cs.list(test.contx)
	assert.Equal(err, test.expected)
	sm.AssertExpectations(t)
}

func TestListServiceWhenCategoriesNotExists(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		expected error
	}{ctx, errNoCategories}
	sm.On("ListCategories", test.contx).Return(mock.Anything, db.ErrCategoryNotExist)
	_, err := cs.list(test.contx)
	assert.Equal(err, test.expected)
	sm.AssertExpectations(t)
}

func TestListServiceWhenInternalError(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		expected error
	}{ctx, errors.New("Internal Error")}
	sm.On("ListCategories", test.contx).Return(mock.Anything, errors.New("Internal Error"))
	_, err := cs.list(test.contx)
	assert.Equal(err, test.expected)
	sm.AssertExpectations(t)
}

func TestSuccessfullUpdateService(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		req      updateRequest
		expected error
	}{ctx, updateRequest{ID: "1", Name: "Sports"}, nil}
	sm.On("UpdateCategory", test.contx, mock.Anything).Return(nil)
	assert.Equal(cs.update(test.contx, test.req), test.expected)
	sm.AssertExpectations(t)
}

func TestUpdateServiceWhenEmptyID(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		req      updateRequest
		expected error
	}{ctx, updateRequest{ID: "", Name: "Sports"}, errEmptyID}
	assert.Equal(cs.update(test.contx, test.req), test.expected)
	sm.AssertExpectations(t)
}

func TestUpdateServiceWhenEmptyName(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		req      updateRequest
		expected error
	}{ctx, updateRequest{ID: "1", Name: ""}, errEmptyName}
	assert.Equal(cs.update(test.contx, test.req), test.expected)
	sm.AssertExpectations(t)
}

func TestUpdateServiceWhenInternalError(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		req      updateRequest
		expected error
	}{ctx, updateRequest{ID: "1", Name: "Sports"}, errors.New("Internal Error")}
	sm.On("UpdateCategory", test.contx, mock.Anything).Return(errors.New("Internal Error"))
	assert.Equal(cs.update(test.contx, test.req), test.expected)
	sm.AssertExpectations(t)
}

func TestSuccessfullFindByIDService(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		id       string
		expected error
	}{ctx, "1", nil}
	sm.On("FindCategoryByID", test.contx, test.id).Return(mock.Anything, nil)
	_, err := cs.findByID(test.contx, test.id)
	assert.Equal(err, test.expected)
	sm.AssertExpectations(t)
}

func TestFindByIDServiceWhenCategoryNotExist(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		id       string
		expected error
	}{ctx, "1", errNoCategoryId}
	sm.On("FindCategoryByID", test.contx, mock.Anything).Return(mock.Anything, db.ErrCategoryNotExist)
	_, err := cs.findByID(test.contx, test.id)
	assert.Equal(err, test.expected)
	sm.AssertExpectations(t)
}

func TestFindByIDServiceWhenInternalError(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		id       string
		expected error
	}{ctx, "1", errors.New("Internal Error")}
	sm.On("FindCategoryByID", test.contx, mock.Anything).Return(mock.Anything, errors.New("Internal Error"))
	_, err := cs.findByID(test.contx, test.id)
	assert.Equal(err, test.expected)
	sm.AssertExpectations(t)
}

func TestSuccessfullDeleteByIDService(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		id       string
		expected error
	}{ctx, "1", nil}
	sm.On("DeleteCategoryByID", test.contx, test.id).Return(nil)
	assert.Equal(cs.deleteByID(test.contx, test.id), test.expected)
	sm.AssertExpectations(t)
}

func TestDeleteByIDServiceWhenCategoryNotExist(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		id       string
		expected error
	}{ctx, "1", errNoCategoryId}
	sm.On("DeleteCategoryByID", test.contx, test.id).Return(db.ErrCategoryNotExist)
	assert.Equal(cs.deleteByID(test.contx, test.id), test.expected)
	sm.AssertExpectations(t)
}

func TestDeleteByIDServiceWhenInternalError(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		id       string
		expected error
	}{ctx, "1", errors.New("Internal Error")}
	sm.On("DeleteCategoryByID", test.contx, test.id).Return(errors.New("Internal Error"))
	assert.Equal(cs.deleteByID(test.contx, test.id), test.expected)
	sm.AssertExpectations(t)
}
