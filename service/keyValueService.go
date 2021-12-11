package service

import (
	"bufio"
	"errors"
	"keyValueProject/core/env"
	"keyValueProject/models"
	"log"
	"os"
	"strings"
	"time"
)

var KeyValueObject map[string]string

var directory string

func Init() {
	if KeyValueObject == nil {
		KeyValueObject = make(map[string]string)
	}
}

func StartInterval() {
	if directory != "" {
		go interval()
	}
}

func FetchValuesFromFile() {
	directory = env.AppSetting.FilePath + env.AppSetting.FileName
	file, err := os.Open(directory)

	if err != nil {
		log.Printf("failed to open")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		splitted := strings.Split(scanner.Text(), ",")
		KeyValueObject[splitted[0]] = splitted[1]
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
		str := ""
		for key, value := range KeyValueObject {
			str += key + "," + value + "\n"
		}
		if last := len(str) - 1; last >= 0 && str[last] == '#' {
			str = str[:last]
		}
		os.WriteFile(directory, []byte(str), 0644)
	}
}
