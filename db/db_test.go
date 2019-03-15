package db

import (
	"context"
	"errors"
	"testing"

	"github.com/joshsoftware/golang-boilerplate/app"
	"github.com/joshsoftware/golang-boilerplate/config"
	"github.com/stretchr/testify/assert"
)

var (
	dbStore    Storer
	categoryID string
)

func Init() {
	config.Load()
	app.Init()
	appDB := app.GetDB()

	dbStore = NewStorer(appDB)
}

func createTestingCategory(ctx context.Context) {
	category := &Category{
		Name: "TestCategoryCreated",
	}
	categoryID, _ = dbStore.CreateCategory(ctx, category)
}

func deleteTestingCategory(ctx context.Context) {
	dbStore.DeleteCategoryByID(ctx, categoryID)
}

func TestSuccessfullCreateCategory(t *testing.T) {
	Init()
	ctx := context.Background()
	category := &Category{
		Name: "TestCategoryCreated",
	}

	id, err := dbStore.CreateCategory(ctx, category)
	assert.Equal(t, nil, err)
	defer dbStore.DeleteCategoryByID(ctx, id)
}

func TestCreateCategoryWhenInternalError(t *testing.T) {
	Init()
	app.GetDB().Close()
	ctx := context.Background()
	category := &Category{
		Name: "TestCategoryCreated",
	}

	id, err := dbStore.CreateCategory(ctx, category)
	assert.IsType(t, errors.New("Internal Error"), err)
	defer dbStore.DeleteCategoryByID(ctx, id)
}

func TestCreateCategoryWhenAlreadyExists(t *testing.T) {
	Init()
	ctx := context.Background()
	createTestingCategory(ctx)
	defer deleteTestingCategory(ctx)
	category := &Category{
		Name: "TestCategoryCreated",
	}

	_, err := dbStore.CreateCategory(ctx, category)
	assert.Equal(t, errCategoryDuplicateKeyValue, err)
}

func TestSuccessfullListCategory(t *testing.T) {
	Init()
	ctx := context.Background()
	var c []Category

	lc, err := dbStore.ListCategories(ctx)
	assert.IsType(t, c, lc)
	assert.Equal(t, nil, err)
}
func TestListCategoryWhenInternalError(t *testing.T) {
	Init()
	app.GetDB().Close()
	ctx := context.Background()
	var c []Category

	lc, err := dbStore.ListCategories(ctx)
	assert.IsType(t, c, lc)
	assert.IsType(t, errors.New("Internal Error"), err)
}
func TestSuccessfullFindCategoryByID(t *testing.T) {
	Init()
	ctx := context.Background()
	createTestingCategory(ctx)
	defer deleteTestingCategory(ctx)
	var c Category

	ls, err := dbStore.FindCategoryByID(ctx, categoryID)
	assert.IsType(t, c, ls)
	assert.Equal(t, nil, err)
}
func TestFindCategoryByIDWhenInternalError(t *testing.T) {
	Init()
	app.GetDB().Close()
	ctx := context.Background()
	var c Category

	ls, err := dbStore.FindCategoryByID(ctx, "fafe5ce0-e52c-400a-9a5e-2b46d26b329a")
	assert.IsType(t, c, ls)
	assert.IsType(t, errors.New("Internal Error"), err)
}

func TestFindCategoryByIDWhenCategoryNotexists(t *testing.T) {
	Init()
	ctx := context.Background()
	CID := "fafe5ce0-e52c-400a-9a5e-2b46d26b329a"
	var c Category
	ls, err := dbStore.FindCategoryByID(ctx, CID)
	assert.IsType(t, c, ls)
	assert.Equal(t, ErrCategoryNotExist, err)
}
func TestSuccessfullDeleteCategoryByID(t *testing.T) {
	Init()
	ctx := context.Background()
	createTestingCategory(ctx)
	defer deleteTestingCategory(ctx)

	err := dbStore.DeleteCategoryByID(ctx, categoryID)
	assert.Equal(t, nil, err)
}
func TestDeleteCategoryByIDWhenInternalError(t *testing.T) {
	Init()
	app.GetDB().Close()
	ctx := context.Background()

	err := dbStore.DeleteCategoryByID(ctx, "fafe5ce0-e52c-400a-9a5e-2b46d26b329a")
	assert.IsType(t, errors.New("Internal Error"), err)
}
func TestSuccessfullUpdateCategoryForCategory(t *testing.T) {
	Init()
	ctx := context.Background()
	createTestingCategory(ctx)
	defer deleteTestingCategory(ctx)
	var c = Category{
		ID:   categoryID,
		Name: "NewUpdateTestCategory",
	}

	err := dbStore.UpdateCategory(ctx, &c)
	assert.Equal(t, nil, err)
}
func TestUpdateCategoryWhenInternalErrorForCategory(t *testing.T) {
	Init()
	ctx := context.Background()
	app.GetDB().Close()
	var c = Category{
		ID:   "fafe5ce0-e52c-400a-9a5e-2b46d26b329a",
		Name: "FakeCategoryNameUpdate",
	}

	err := dbStore.UpdateCategory(ctx, &c)
	assert.IsType(t, errors.New("NewUpdateTestCategory"), err)
}
