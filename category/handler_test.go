package category

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func makeHTTPCall(handler http.HandlerFunc, method, path, body string) (rr *httptest.ResponseRecorder) {
	request := []byte(body)
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(request))
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return
}

// Create:
func TestSuccessfullCreate(t *testing.T) {
	cs := &CategoryServiceMock{}
	cs.On("create", mock.Anything, mock.Anything).Return("", nil)

	rr := makeHTTPCall(Create(cs), http.MethodPost, "/categories", `{"name":"Sports"}`)

	checkResponseCode(t, http.StatusCreated, rr.Code)
	cs.AssertExpectations(t)
}

func TestCreateWhenInvalidRequestBody(t *testing.T) {
	cs := &CategoryServiceMock{}

	rr := makeHTTPCall(Create(cs), http.MethodPost, "/categories", `{"name":"",}`)

	checkResponseCode(t, http.StatusBadRequest, rr.Code)
	cs.AssertExpectations(t)
}

func TestCreateWhenEmptyName(t *testing.T) {
	cs := &CategoryServiceMock{}
	cs.On("create", mock.Anything, mock.Anything).Return("", errEmptyName)

	rr := makeHTTPCall(Create(cs), http.MethodPost, "/categories", `{"name":""}`)

	checkResponseCode(t, http.StatusBadRequest, rr.Code)
	cs.AssertExpectations(t)
}

func TestCreateWhenInternalError(t *testing.T) {
	cs := &CategoryServiceMock{}
	cs.On("create", mock.Anything, mock.Anything).Return("", errors.New("Internal Error"))

	rr := makeHTTPCall(Create(cs), http.MethodPost, "/categories", `{"name":"Sports"}`)

	checkResponseCode(t, http.StatusInternalServerError, rr.Code)
	cs.AssertExpectations(t)
}

// List :
func TestSuccessfullList(t *testing.T) {
	cs := &CategoryServiceMock{}
	cs.On("list", mock.Anything).Return(mock.Anything, nil)

	rr := makeHTTPCall(List(cs), http.MethodGet, "/categories", "")

	checkResponseCode(t, http.StatusOK, rr.Code)
	cs.AssertExpectations(t)
}

func TestListWhenNoCategories(t *testing.T) {
	cs := &CategoryServiceMock{}
	cs.On("list", mock.Anything).Return(mock.Anything, errNoCategories)

	rr := makeHTTPCall(List(cs), http.MethodGet, "/categories", "")

	checkResponseCode(t, http.StatusNotFound, rr.Code)
	cs.AssertExpectations(t)
}

func TestListInternalError(t *testing.T) {
	cs := &CategoryServiceMock{}
	cs.On("list", mock.Anything).Return(mock.Anything, errors.New("Internal Error"))

	rr := makeHTTPCall(List(cs), http.MethodGet, "/categories", "")

	checkResponseCode(t, http.StatusInternalServerError, rr.Code)
	cs.AssertExpectations(t)
}

//FindById
//not bad reqe
//not find err
func TestSuccessfullFindByID(t *testing.T) {
	cs := &CategoryServiceMock{}
	cs.On("findByID", mock.Anything, mock.Anything).Return(mock.Anything, nil)

	rr := makeHTTPCall(FindByID(cs), http.MethodGet, "/categories/1", "")

	checkResponseCode(t, http.StatusOK, rr.Code)
	cs.AssertExpectations(t)
}

func TestFindByIDWhenIDNotExist(t *testing.T) {
	cs := &CategoryServiceMock{}
	cs.On("findByID", mock.Anything, mock.Anything).Return(mock.Anything, errNoCategoryId)

	rr := makeHTTPCall(FindByID(cs), http.MethodGet, "/categories/1", "")

	checkResponseCode(t, http.StatusNotFound, rr.Code)
	cs.AssertExpectations(t)
}

func TestFindByIdWhenInternalError(t *testing.T) {
	cs := &CategoryServiceMock{}
	cs.On("findByID", mock.Anything, mock.Anything).Return(mock.Anything, errors.New("Internal Error"))

	rr := makeHTTPCall(FindByID(cs), http.MethodGet, "/categories/1", "")

	checkResponseCode(t, http.StatusInternalServerError, rr.Code)
	cs.AssertExpectations(t)
}

//DeleteByID
func TestSuccessfullDeleteByID(t *testing.T) {
	cs := &CategoryServiceMock{}

	cs.On("deleteByID", mock.Anything, mock.Anything).Return(nil)

	rr := makeHTTPCall(DeleteByID(cs), http.MethodDelete, "/categories/1", "")

	checkResponseCode(t, http.StatusOK, rr.Code)
	cs.AssertExpectations(t)
}

func TestDeleteByIDWhenIDNotExist(t *testing.T) {
	cs := &CategoryServiceMock{}
	cs.On("deleteByID", mock.Anything, mock.Anything).Return(errNoCategoryId)

	rr := makeHTTPCall(DeleteByID(cs), http.MethodDelete, "/categories/1", "")

	checkResponseCode(t, http.StatusNotFound, rr.Code)
	cs.AssertExpectations(t)
}

func TestDeleteByIDWhenInternalError(t *testing.T) {
	cs := &CategoryServiceMock{}
	cs.On("deleteByID", mock.Anything, mock.Anything).Return(errors.New("Internal Error"))

	rr := makeHTTPCall(DeleteByID(cs), http.MethodDelete, "/categories/1", "")

	checkResponseCode(t, http.StatusInternalServerError, rr.Code)
	cs.AssertExpectations(t)
}

func TestSuccessfullUpdate(t *testing.T) {
	cs := &CategoryServiceMock{}
	cs.On("update", mock.Anything, mock.Anything).Return(nil)

	rr := makeHTTPCall(Update(cs), http.MethodPut, "/categories", `{"id":"1", "name":"sports"}`)

	checkResponseCode(t, http.StatusOK, rr.Code)
	cs.AssertExpectations(t)
}

func TestUpdateWhenInvalidRequestBody(t *testing.T) {
	cs := &CategoryServiceMock{}

	rr := makeHTTPCall(Update(cs), http.MethodPut, "/categories", `{"id":"1", "name":"sports",}`)

	checkResponseCode(t, http.StatusBadRequest, rr.Code)
	cs.AssertExpectations(t)
}

func TestUpdateWhenEmptyID(t *testing.T) {
	cs := &CategoryServiceMock{}
	cs.On("update", mock.Anything, mock.Anything).Return(errEmptyID)

	rr := makeHTTPCall(Update(cs), http.MethodPut, "/categories", `{"name":"Sports"}`)

	checkResponseCode(t, http.StatusBadRequest, rr.Code)
	cs.AssertExpectations(t)
}

func TestUpdateWhenEmptyName(t *testing.T) {
	cs := &CategoryServiceMock{}
	cs.On("update", mock.Anything, mock.Anything).Return(errEmptyName)

	rr := makeHTTPCall(Update(cs), http.MethodPut, "/categories", `{"id":"1"}`)

	checkResponseCode(t, http.StatusBadRequest, rr.Code)
	cs.AssertExpectations(t)
}

func TestUpdateWhenInternalError(t *testing.T) {
	cs := &CategoryServiceMock{}
	cs.On("update", mock.Anything, mock.Anything).Return(errors.New("Internal Error"))

	rr := makeHTTPCall(Update(cs), http.MethodPut, "/categories", `{"id":"1"}`)

	checkResponseCode(t, http.StatusInternalServerError, rr.Code)
	cs.AssertExpectations(t)
}
