package cfg

import (
	"github.com/BurntSushi/toml"
	"github.com/fsnotify/fsnotify"
	"log"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

// 配置对象
type cfgHandler struct {
	confInfo interface{}
}

// New get a cfg handler
func New(filePath string) (*cfgHandler, error) {
	filePath, err := filepath.Abs(filePath)
	if err != nil {
		panic(err)
	}

	handler := cfgHandler{}
	if _, err := toml.DecodeFile(filePath, &handler.confInfo); err != nil {
		if err != nil {
			return nil, err
		}
	}
	go watchChange(filePath, &handler)
	return &handler, nil
}

// GetString get config string type
func (c *cfgHandler) GetString(key string) string {
	ret := c.ParseNode(key)
	showRet, ok := ret.(string)
	if ok != true {
		return ""
	}
	return showRet
}

// GetInt get config int type
func (c *cfgHandler) GetInt(key string) int {
	ret := c.ParseNode(key)
	showRet, ok := ret.(int64)
	if ok != true {
		return 0
	}
	return int(showRet)
}

// GetInt64 get config int64 type
func (c *cfgHandler) GetInt64(key string) int64 {
	ret := c.ParseNode(key)
	showRet, ok := ret.(int64)
	if ok != true {
		return 0
	}
	return showRet
}

// GetBool get config bool type
func (c *cfgHandler) GetBool(key string) bool {
	ret := c.ParseNode(key)
	showRet, ok := ret.(bool)
	if ok != true {
		return false
	}
	return showRet
}

// GetDuration get config time.Duration type
func (c *cfgHandler) GetDuration(key string) time.Duration {
	ret := c.GetInt64(key)
	showRet := time.Duration(ret)
	return showRet
}

// GetSliceInt get slice int type
func (c *cfgHandler) GetSliceInt(key string) []int {
	slice := c.getSlice(key)
	showRet := make([]int, 0)
	if len(slice) == 0 {
		return showRet
	}
	for _, item := range slice {
		value, ok := item.(int64)
		if ok != true {
			break
		}
		showRet = append(showRet, int(value))
	}

	return showRet
}

// GetSliceInt64 get slice int64 type
func (c *cfgHandler) GetSliceInt64(key string) []int64 {
	slice := c.getSlice(key)
	showRet := make([]int64, 0)
	if len(slice) == 0 {
		return showRet
	}
	for _, item := range slice {
		value, ok := item.(int64)
		if ok != true {
			break
		}
		showRet = append(showRet, value)
	}

	return showRet
}

// GetSliceString get slice string type
func (c *cfgHandler) GetSliceString(key string) []string {
	slice := c.getSlice(key)
	showRet := make([]string, 0)
	if len(slice) == 0 {
		return showRet
	}
	for _, item := range slice {
		value, ok := item.(string)
		if ok != true {
			break
		}
		showRet = append(showRet, value)
	}
	return showRet
}

// ParseNode parse config node
func (c *cfgHandler) ParseNode(parse string) (ret interface{}) {
	nodeParse := strings.Split(parse, ".")

	ret = c.confInfo
	for _, item := range nodeParse {
		ret = readNode(item, ret)
	}

	return
}

// getSlice get config slice []interface{} type
func (c *cfgHandler) getSlice(key string) []interface{} {

	ret := c.ParseNode(key)
	slice, ok := ret.([]interface{})
	if ok != true {
		return []interface{}{}
	}
	return slice
}

// readNode ...
func readNode(key string, node interface{}) (ret interface{}) {
	if node == nil {
		return
	}
	nodeType := reflect.TypeOf(node).String()
	if nodeType == "map[string]interface {}" {
		nodeInfo := node.(map[string]interface{})
		return nodeInfo[key]
	}

	return node
}

func watchChange(path string, handler *cfgHandler) {
	watcher, err := fsnotify.NewWatcher()
	defer watcher.Close()
	if err != nil {
		panic(err)
	}

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Printf("reload config " + path)
				if event.Op&fsnotify.Write == fsnotify.Write {
					var confInfo interface{}
					toml.DecodeFile(path, &confInfo)
					handler.confInfo = confInfo
				}
			}
		}
	}()

	watcher.Add(path)
	<-done

}
