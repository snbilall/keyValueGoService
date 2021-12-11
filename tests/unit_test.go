package tests

import (
	"keyValueProject/models"
	"keyValueProject/service"
	"testing"
)

func TestAddKeyValue(t *testing.T) {
	service.Init()
	key := "Sean"
	value := "Paul"
	obj := models.KeyValueRequest{Key: key, Value: value}
	service.AddKeyValue(obj)
	if service.KeyValueObject[key] == "" {
		t.Fatalf(`adding key %s, value %s test failed`, key, value)
	}
}

func TestGetKeyValue(t *testing.T) {
	service.Init()
	key := "Sean"
	value := "Paul"
	obj := models.KeyValueRequest{Key: key, Value: value}
	service.AddKeyValue(obj)
	if _, err := service.GetKeyValue(key); err != nil {
		t.Fatalf(`getting key %s, value %s test failed`, key, value)
	}
}
