//Package test
/*
@Title: base.go
@Description
@Author: kkw 2023/1/3 17:19
*/
package test

import (
	"github.com/magiconair/properties/assert"
	"go.kkw.top/gamesTimeLine/cmd/api/setupRouter"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Base(t *testing.T, method, url string, body io.Reader) {
	router := setupRouter.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/user", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
