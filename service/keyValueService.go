package service

import (
	"encoding/json"
	"errors"
	"keyValueProject/core/env"
	"keyValueProject/models"
	"os"
	"time"
)

var KeyValueObject map[string]string

var filePath string

func Init() {
	filePath = os.TempDir() + string(os.PathSeparator) + env.AppSetting.FileName
	if KeyValueObject == nil {
		KeyValueObject = make(map[string]string)
	}
}

func StartInterval() {
	go interval()
}

func FetchValuesFromFile() {
	if file, err := os.ReadFile(filePath); err == nil {
		json.Unmarshal(file, &KeyValueObject)
	}
}

func AddKeyValue(obj models.KeyValueRequest) {
	KeyValueObject[obj.Key] = obj.Value
}

func GetKeyValue(key string) (string, error) {
	str := KeyValueObject[key]
	var err error = nil
	if str == "" {
		err = errors.New("key not found")
	}
	return str, err
}

func interval() {
	intervalValue := time.Duration(env.AppSetting.Interval) * time.Second
	for range time.Tick(intervalValue) {
		if json, err := json.Marshal(KeyValueObject); err == nil {
			os.WriteFile(filePath, []byte(json), 0644)
		}
	}
}
