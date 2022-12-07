package utils

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"math/rand"
	"os"
	"strings"
	"time"
)

// ParamsFilter 过滤指定数组中的key
func ParamsFilter(isFilterStr string, params map[string]any) map[string]any {
	var data = make(map[string]any)
	for key, value := range params {
		if value != "" {
			find := strings.Contains(isFilterStr, key)
			if !find {
				data[key] = value
			}
		}
	}
	return data
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { //文件或者目录存在
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func UUID() string {
	return uuid.Must(uuid.NewV4()).String()
}

// AnyToMap interface 转 map
func AnyToMap(v any) (map[string]any, error) {
	dataJson, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	var MapData map[string]any
	err = json.Unmarshal(dataJson, &MapData)
	if err != nil {
		return nil, err
	}
	return MapData, nil
}

// RandomString 生成指定长度的随机字符
func RandomString(n int) string {
	var letters = []byte("qwertyuioplkjhgfdsazxcvbnmMNBVCXZASDFGHJKLPOIUYTREWQ")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
