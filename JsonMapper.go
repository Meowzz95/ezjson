package ezjson

import "encoding/json"

type JsonMapper struct {
	jsonStr string
}

func NewJsonMapper(jsonStr string) *JsonMapper {
	return &JsonMapper{jsonStr: jsonStr}
}

func (it *JsonMapper) GetJsonPart() (*JsonPart, error) {
	jsonBytes := []byte(it.jsonStr)
	var part interface{}
	err := json.Unmarshal(jsonBytes, &part)
	if err != nil {
		return nil, err
	}
	return NewJsonPart("root", part), nil
}
