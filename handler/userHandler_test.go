package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BohdanPatrash/indenticon/dto"

	"github.com/BohdanPatrash/indenticon/repo"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/user", UserGet())
	r.POST("/user", UserPost())
	return r
}

func TestSaveAndGetUsers(t *testing.T) {
	//Setup database
	repo.Init()
	repo.DatabaseSetup()

	r := SetupRouter()

	//testing if users table is empty
	w := performRequest(r, "GET", "/user", nil)
	assert.Equal(t, http.StatusOK, w.Code)
	var response []dto.User
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, make([]dto.User, 0), response, "There should be no users atm")

	//testing if post request passed
	requestBody := []byte(`{"email":"test@mail.com"}`)
	w = performRequest(r, "POST", "/user", bytes.NewBuffer(requestBody))
	assert.Equal(t, http.StatusOK, w.Code)

	//testing if we got user saved
	w = performRequest(r, "GET", "/user", nil)
	assert.Equal(t, http.StatusOK, w.Code)
	err = json.Unmarshal([]byte(w.Body.String()), &response)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1, len(response), "There should be 1 user")
}
