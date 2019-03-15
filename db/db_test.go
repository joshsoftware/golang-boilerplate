package db

import (
	"context"
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
