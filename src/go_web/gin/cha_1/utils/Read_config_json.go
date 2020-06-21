package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"learn-go/src/go_web/gin/cha_1/model"
	"sync"
)

var (
	configMux    sync.RWMutex
	GlobalConfig *model.ProjectConf
)

func LoadConfig(filepath string) error {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("读取配置文件失败 !", err)
		panic(err)
	}
	var conf model.ProjectConf

	if err = json.Unmarshal(file, &conf); err != nil {
		fmt.Println("配置文件读取失败", err)
		panic(err)
	}

	configMux.Lock()
	GlobalConfig = &conf
	configMux.Unlock()
	return nil
}
