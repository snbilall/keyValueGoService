package env

import (
	"encoding/json"
	"fmt"
	"os"
)

type AppSettings struct {
	Interval int
	FilePath string
	FileName string
}

var AppSetting = &AppSettings{}

func Init() {
	if bytes, err := os.ReadFile("conf/sharedConf.json"); err == nil && bytes != nil {
		jsonErr := json.Unmarshal(bytes, AppSetting)
		if jsonErr != nil {
			fmt.Println("configuration parse error: ", jsonErr)
		}
	} else {
		fmt.Println("configuration read error: ", err)
	}
}
