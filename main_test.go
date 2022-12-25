package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := setupGin()

	respByID := httptest.NewRecorder()
	respAll := httptest.NewRecorder()
	respPost := httptest.NewRecorder()
	reqID, _ := http.NewRequest("GET", "/albums/1", nil)
	router.ServeHTTP(respByID, reqID)

	req, _ := http.NewRequest("GET", "/albums", nil)
	router.ServeHTTP(respAll, req)
	json:=[]byte(`{"id": "6","title": "Excelsior","artist": "Lmorphine","price": 36}`)
	reqPost, _ := http.NewRequest("POST", "/albums",bytes.NewBuffer(json))
	router.ServeHTTP(respPost, reqPost)

	assert.Equal(t, 200, respByID.Code)
	assert.Equal(t, 200, respAll.Code)
	assert.Equal(t, 201, respPost.Code)
}
