package ezjson

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type JsonPart struct {
	key  string
	part interface{}
	err error
}

func NewJsonPart(key string, part interface{}) *JsonPart {
	return &JsonPart{key: key, part: part}
}

func (it *JsonPart) Error() string{
	return it.err.Error()
}

func (it *JsonPart) GetRawMap() map[string]interface{} {
	aMap,_:=it.getMap()
	// GetPart has made sure that JsonPart must have a map[string]interface{} part
	// so ignore getMap() error
	return aMap
}

func (it *JsonPart) GetFloat64(key string) (float64, error) {
	m, ok := it.getMap()
	if !ok {
		return 0, NewNonMapError(it.key)
	}

	if keyExist := it.keyExist(key, m); !keyExist {
		return 0, NewKeyDoesNotExistError(it.key, key)
	}

	valueType := it.getType(m[key])
	if valueType == nil {
		return 0, nil
	}
	if valueType.Kind() != reflect.Float64 {
		return 0, NewValueTypeMismatchError(it.key, key, "float64", valueType.Kind().String())
	}
	value := m[key].(float64)
	return value, nil
}

func (it *JsonPart) GetFloat64Or(key string, defaultValue float64) float64 {
	v, err := it.GetFloat64(key)
	if err != nil {
		return defaultValue
	}
	return v
}
func (it *JsonPart) GetFloat64F(key string) float64{
	v,err:=it.GetFloat64(key)
	it.err = err
	return v
}

func (it *JsonPart) GetPart(key string) (*JsonPart, error) {
	m, ok := it.getMap()
	if !ok {
		return nil, NewNonMapError(it.key)
	}
	if keyExist := it.keyExist(key, m); !keyExist {
		return nil, NewKeyDoesNotExistError(it.key, key)
	}
	valueType := it.getType(m[key])
	if valueType == nil {
		return nil, nil
	}
	if valueType.Kind() != reflect.Map {
		return nil, NewValueTypeMismatchError(it.key, key, "a JSON object", valueType.Kind().String())
	}
	value := m[key].(map[string]interface{})
	return NewJsonPart(key, value), nil
}

func (it *JsonPart) GetPartF(key string) *JsonPart {
	v,err:=it.GetPart(key)
	it.err = err
	return v
}

func (it *JsonPart) GetString(key string) (string, error) {
	m, ok := it.getMap()
	if !ok {
		return "", NewNonMapError(it.key)
	}
	if keyExist := it.keyExist(key, m); !keyExist {
		return "", NewKeyDoesNotExistError(it.key, key)
	}
	valueType := it.getType(m[key])
	if valueType == nil {
		return "", nil
	}
	if valueType.Kind() != reflect.String {
		return "", NewValueTypeMismatchError(it.key, key, "string", valueType.Kind().String())
	}
	value := m[key].(string)
	return value, nil
}

func (it *JsonPart) GetStringF(key string) string{
	v,err:=it.GetString(key)
	it.err = err
	return v
}

func (it *JsonPart) GetBoolean(key string) (bool, error) {
	m, ok := it.getMap()
	if !ok {
		return false, NewNonMapError(it.key)
	}
	if keyExist := it.keyExist(key, m); !keyExist {
		return false, NewKeyDoesNotExistError(it.key, key)
	}
	valueType := it.getType(m[key])
	if valueType == nil {
		return false, nil
	}
	if valueType.Kind() != reflect.Bool {
		return false, NewValueTypeMismatchError(it.key, key, "boolean", valueType.Kind().String())
	}
	value := m[key].(bool)
	return value, nil
}

func (it *JsonPart) GetBooleanF(key string) bool{
	v,err:=it.GetBoolean(key)
	it.err=err
	return v
}

func (it *JsonPart) GetStringCasted(key string) (string, error) {
	// test all possibilities - bool, float64, string, part, array
	if vBool, err := it.GetBoolean(key); err == nil {
		return strconv.FormatBool(vBool), nil
	}
	if vFloat, err := it.GetFloat64(key); err == nil {
		if isFloatInt(vFloat) {
			return strconv.FormatInt(int64(vFloat), 10), nil
		}
		return fmt.Sprintf("%f", vFloat), nil
	}
	if vString, err := it.GetString(key); err == nil {
		return vString, nil
	}
	if vPart, err := it.GetPart(key); err == nil {
		jsonStrBytes, _ := json.Marshal(vPart.part)
		return string(jsonStrBytes), nil
	}
	m, _ := it.getMap()
	actualType := it.getType(m[key])
	return "", NewValueTypeMismatchError(it.key, key, "bool/float64/string/JsonPart/JsonArray", actualType.String())
}
func (it *JsonPart) GetStringCastedF(key string) string {
	v,err:=it.GetStringCasted(key)
	it.err = err
	return v
}

func isFloatInt(val float64) bool {
	return val == float64(int(val))
}

func (it *JsonPart) getType(value interface{}) reflect.Type {
	return reflect.TypeOf(value)
}

func (it *JsonPart) getMap() (map[string]interface{}, bool) {
	result, ok := it.part.(map[string]interface{})
	return result, ok
}

func (it *JsonPart) keyExist(key string, m map[string]interface{}) bool {
	_, ok := m[key]
	return ok
}
