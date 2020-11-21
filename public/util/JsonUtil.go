package util

import "encoding/json"

/*
  JSON (map转json)
*/
func MapToJson(data map[string]interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

/*
  将对象转为json
*/
func ToJsonString(data interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), err
}

/*
	JSON (json转map)
*/
func ParseJsonString(data string) map[string]interface{} {
	var jsonData map[string]interface{}
	json.Unmarshal([]byte(data), &jsonData)
	return jsonData
}

/*
	JSONstring (json转IntList)
*/
func ParseToIntList(data string) []int {
	var tmp = make([]int, 0)
	json.Unmarshal([]byte(data), &tmp)
	return tmp
}
