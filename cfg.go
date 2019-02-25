package cfg

import (
	"github.com/BurntSushi/toml"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

type cfgHandler struct{
	confInfo interface{}
}

var cfgHandlerVal = cfgHandler{}

func InitConfFile(filePath string){
	filePath, err := filepath.Abs(filePath)
	if err != nil{
		panic(err)
	}
	if _, err := toml.DecodeFile(filePath, &cfgHandlerVal.confInfo); err != nil {
		if err != nil{
			panic(err)
		}
	}
}

func GetString(key string) string{
	ret := ParseNode(key)
	showRet,ok:=ret.(string)
	if ok != true {
		return ""
	}
	return showRet
}

func  GetInt(key string) int{
	ret := ParseNode(key)
	showRet,ok:=ret.(int64)
	if ok != true {
		return 0
	}
	return int(showRet)
}

func GetInt64(key string) int64{
	ret := ParseNode(key)
	showRet,ok:=ret.(int64)
	if ok != true {
		return 0
	}
	return showRet
}


func GetBool(key string) bool{
	ret := ParseNode(key)
	showRet,ok:=ret.(bool)
	if ok != true {
		return false
	}
	return showRet
}


func GetDuration(key string) time.Duration {
	ret := GetInt64(key)
	showRet := time.Duration(ret)
	return showRet
}


func ParseNode(parse string)(ret interface{}){
	nodeParse := strings.Split(parse,".")

	ret = cfgHandlerVal.confInfo
	for _,item:=range nodeParse{
		ret = readNode(item, ret)
	}

	return
}


// 解析配置节点
func readNode(key string, node interface{})(ret interface{}){
	if node == nil {
		return
	}
	nodeType := reflect.TypeOf(node).String()
	if nodeType == "map[string]interface {}"{
		nodeInfo := node.(map[string]interface{})
		return nodeInfo[key]
	}

	return node
}