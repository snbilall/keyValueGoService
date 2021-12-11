package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"keyValueProject/models"
	"keyValueProject/routers"
	"keyValueProject/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSuccessAddKeyValueService(t *testing.T) {
	service.Init()
	testRouter := routers.SetupRouter()
	key := "Sean"
	value := "Paul"
	body := models.KeyValueRequest{Key: key, Value: value}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		fmt.Println(err)
	}
	req, err := http.NewRequest(http.MethodPost, "/createKeyValue", &buf)
	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	if resp.Code != 201 {
		t.Fatalf(`adding key %s, value %s service test failed`, key, value)
	}
}

func TestUnsuccessAddKeyValueService(t *testing.T) {
	service.Init()
	testRouter := routers.SetupRouter()
	value := "Paul"
	body := models.KeyValueRequest{Value: value}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		fmt.Println(err)
	}
	req, err := http.NewRequest(http.MethodPost, "/createKeyValue", &buf)
	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	if resp.Code != 400 {
		t.Fatalf(`adding empty key, value %s service test failed`, value)
	}
}
