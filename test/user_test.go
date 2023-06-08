//Package user
/*
@Title: user_test.go
@Description
@Author: kkw 2023/1/3 16:51
*/
package test

import (
	"github.com/magiconair/properties/assert"
	"go.kkw.top/gamesTimeLine/cmd/api/setupRouter"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var router = setupRouter.SetupRouter()

func TestUser_GetAll(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/user/all", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUser_Add(t *testing.T) {

	w := httptest.NewRecorder()

	payload := strings.NewReader(`{"user_name": "看看我"}`)
	req, _ := http.NewRequest("POST", "/v1/user", payload)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUser_Add_BindErr(t *testing.T) {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/v1/user", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
