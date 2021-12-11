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
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var testRouter *gin.Engine

func TestMain(m *testing.M) {
	service.Init()
	testRouter = routers.SetupRouter()
	os.Exit(m.Run())
}

func TestSuccessAddKeyValueService(t *testing.T) {
	obj := models.KeyValueRequest{Key: "Sean", Value: "Paul"}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(obj)
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
		t.Fatalf(`adding key %s, value %s service test failed`, obj.Key, obj.Value)
	}
}

func TestUnsuccessAddKeyValueService(t *testing.T) {
	obj := models.KeyValueRequest{Value: "Paul"}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(obj)
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
		t.Fatalf(`adding empty key, value %s service test failed`, obj.Value)
	}
}

func TestSuccessGetKeyValueService(t *testing.T) {
	obj := models.KeyValueRequest{Key: "Paul", Value: "Scholes"}
	service.AddKeyValue(obj)
	req, err := http.NewRequest(http.MethodGet, "/keyValues/"+obj.Key, nil)
	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	if resp.Code != 200 {
		t.Fatalf(`get success key %s, value %s service test failed`, obj.Key, obj.Value)
	}
}

func TestUnsuccessGetKeyValueService(t *testing.T) {
	key := "Herald"
	req, err := http.NewRequest(http.MethodGet, "/keyValues/"+key, nil)
	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	if resp.Code != 404 {
		t.Fatalf(`get unsucess key %s service test failed`, key)
	}
}
