package service

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var KeyValueObject map[string]string

const directory string = "C:\\Users\\Administrator\\Desktop\\project\\file.txt"

func InitKeyValueObject() {
	if KeyValueObject == nil {
		KeyValueObject = make(map[string]string)
		if _, err := os.Stat(directory); err == nil {
			if bytes, err := os.ReadFile(directory); err == nil && bytes != nil {
				myString := string(bytes)
				splitted := strings.Split(myString, "#")
				for _, item := range splitted {
					splittedItem := strings.Split(item, ",")
					KeyValueObject[splittedItem[0]] = splittedItem[1]
				}
			}
		}
	}
	go heartBeat()
}

func AddKeyValue(key string, value string) {
	KeyValueObject[key] = value
	str := ""
	for key, value := range KeyValueObject {
		str += key + "," + value + "#"
	}
	if last := len(str) - 1; last >= 0 && str[last] == '#' {
		str = str[:last]
	}
	os.WriteFile(directory, []byte(str), 0644)
}

func heartBeat() {
	for range time.Tick(time.Second * 1) {
		fmt.Println("Foo")
	}
}
