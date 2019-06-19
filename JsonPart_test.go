package ezjson

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const SAMPLE_JSON_1 = `
{
    "Status": "SUCCESS",
    "Err": null,
    "Info": "",
	"Price": 25.1,
	"NumOfPeople":999,
    "IsHigh": true,
    "Payload": {
        "ID": 14,
        "CreatedAt": "2019-06-13T10:28:42.396549+08:00",
        "UpdatedAt": "2019-06-13T10:28:42.396549+08:00",
        "DeletedAt": null,
        "Key": "test_ke11111y",
        "Value": "test_value22222"
    }
}
`

const ARRAY_JSON_1 = `
	[{"A":"Value A"},{"B":"Value B"}]
`

const CORRUPTED_PART_JSON_1 = `
{
    "Status": "SUCCESS",
    "Err": null,
    "Info": "",
	"Price": 25.1,
	"NumOfPeople":999,
    "IsHigh": true,
    "Payload": {
        "ID":'::
    }
}
`

func getJsonPart(jsonStr string) *JsonPart {
	jsonMapper := NewJsonMapper(jsonStr)
	jsonPart, err := jsonMapper.GetJsonPart()
	if err != nil {
		panic("Error getting json part")
	}
	return jsonPart
}

func TestNewJsonMapper_fail(t *testing.T) {
	jsonMapper := NewJsonMapper("{rubbish json :-=-={{{}")
	jsonPart, err := jsonMapper.GetJsonPart()
	assert.NotNil(t, err)
	assert.Nil(t, jsonPart)
}
func TestJsonPart_GetFloat64(t *testing.T) {
	jsonPart := getJsonPart(SAMPLE_JSON_1)

	price, err := jsonPart.GetFloat64("Price")
	assert.Nil(t, err)
	assert.Equal(t, price, 25.1)
}
func TestJsonPart_GetFloat64_TypeMismatch(t *testing.T) {
	jsonPart := getJsonPart(SAMPLE_JSON_1)

	price, err := jsonPart.GetFloat64("Status")
	assert.NotNil(t, err)
	assert.Equal(t, price, float64(0))
	assert.Equal(t, err.Error(), fmt.Sprintf(VALUE_TYPE_MISMATCH_ERR_TEMPLATE, jsonPart.key, "Status", "string", "float64"))
}

func TestJsonPart_GetFloat64_fail(t *testing.T) {
	jsonPart := getJsonPart(SAMPLE_JSON_1)
	v, err := jsonPart.GetFloat64("Info")
	assert.NotNil(t, err)
	assert.Equal(t, v, float64(0))
}

func TestJsonPart_WrongKey(t *testing.T) {
	jsonPart := getJsonPart(SAMPLE_JSON_1)
	const nonExistField = "NON_EXIST_FIELD"
	v, err := jsonPart.GetFloat64(nonExistField)
	assert.NotNil(t, err)
	assert.Equal(t, v, float64(0))

	vBool, err := jsonPart.GetBoolean(nonExistField)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), fmt.Sprintf(KEY_DOES_NOT_EXIST_TEMPLATE, jsonPart.key, nonExistField))
	assert.Equal(t, vBool, false)

	vStr, err := jsonPart.GetString(nonExistField)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), fmt.Sprintf(KEY_DOES_NOT_EXIST_TEMPLATE, jsonPart.key, nonExistField))
	assert.Equal(t, vStr, "")

	vPart, err := jsonPart.GetPart(nonExistField)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), fmt.Sprintf(KEY_DOES_NOT_EXIST_TEMPLATE, jsonPart.key, nonExistField))
	assert.Nil(t, vPart)
}

func TestJsonPart_GetFloat64Or(t *testing.T) {
	jsonPart := getJsonPart(SAMPLE_JSON_1)
	v := jsonPart.GetFloat64Or("Price", -1)
	assert.Equal(t, v, 25.1)

	defaultValue := jsonPart.GetFloat64Or("P", -1)
	assert.Equal(t, defaultValue, float64(-1))
}

func TestJsonPart_GetPart(t *testing.T) {
	jsonPart := getJsonPart(SAMPLE_JSON_1)
	innerPart, err := jsonPart.GetPart("Payload")
	assert.Nil(t, err)
	assert.Equal(t, innerPart.key, "Payload")
}

func TestJsonPart_GetPart_NonPart_Fail(t *testing.T) {
	jsonPart := getJsonPart(SAMPLE_JSON_1)
	innerPart, err := jsonPart.GetPart("Price")
	assert.NotNil(t, err)
	assert.Nil(t, innerPart)

}

func TestJsonPart_GetPart_SubTest(t *testing.T) {
	jsonPart := getJsonPart(SAMPLE_JSON_1)
	innerPart, err := jsonPart.GetPart("Payload")
	assert.Nil(t, err)
	v, err := innerPart.GetFloat64("ID")
	assert.Nil(t, err)
	assert.Equal(t, v, float64(14))
}

func TestJsonPart_GetString(t *testing.T) {
	jsonPart := getJsonPart(SAMPLE_JSON_1)
	v, err := jsonPart.GetString("Status")
	assert.Nil(t, err)
	assert.Equal(t, v, "SUCCESS")
}

func TestJsonPart_GetNull(t *testing.T) {
	jsonPart := getJsonPart(SAMPLE_JSON_1)
	v, err := jsonPart.GetString("Err")
	assert.Nil(t, err)
	assert.Equal(t, v, "")

	vFloat, err := jsonPart.GetFloat64("Err")
	assert.Nil(t, err)
	assert.Equal(t, vFloat, float64(0))

	vBool, err := jsonPart.GetBoolean("Err")
	assert.Nil(t, err)
	assert.Equal(t, vBool, false)

	vPart, err := jsonPart.GetPart("Err")
	assert.Nil(t, err)
	assert.Nil(t, vPart)
}

func TestJsonPart_GetStringCasted(t *testing.T) {
	jsonPart := getJsonPart(SAMPLE_JSON_1)
	//float
	v, err := jsonPart.GetStringCasted("Price")
	assert.Nil(t, err)
	assert.Equal(t, v, fmt.Sprintf("%f", float64(25.1)))

	//bool
	v, err = jsonPart.GetStringCasted("IsHigh")
	assert.Nil(t, err)
	assert.Equal(t, v, "true")

	//part
	v, err = jsonPart.GetStringCasted("Payload")
	assert.Nil(t, err)

	var sampleJson1Obj map[string]interface{}
	json.Unmarshal([]byte(SAMPLE_JSON_1), &sampleJson1Obj)

	expectedStrBytes, _ := json.Marshal(sampleJson1Obj["Payload"])
	assert.Equal(t, v, string(expectedStrBytes))

	//string
	v, err = jsonPart.GetStringCasted("Status")
	assert.Nil(t, err)
	assert.Equal(t, v, "SUCCESS")
}

func TestJsonPart_GetStringCasted_Int(t *testing.T) {
	jsonPart := getJsonPart(SAMPLE_JSON_1)
	v, err := jsonPart.GetStringCasted("NumOfPeople")
	assert.Nil(t, err)
	assert.Equal(t, v, "999")
}

func TestJsonPart_WrongParent(t *testing.T) {
	jsonPart := getJsonPart(ARRAY_JSON_1)
	v, err := jsonPart.GetString("A")
	assert.NotNil(t, err)
	assert.Equal(t, v, "")
	assert.Equal(t, err.Error(), fmt.Sprintf(NON_MAP_ERROR_TEMPLATE, jsonPart.key))

	vFloat, err := jsonPart.GetFloat64("A")
	assert.NotNil(t, err)
	assert.Equal(t, vFloat, float64(0))

	vBool, err := jsonPart.GetBoolean("A")
	assert.NotNil(t, err)
	assert.Equal(t, vBool, false)

	vStr, err := jsonPart.GetString("A")
	assert.NotNil(t, err)
	assert.Equal(t, vStr, "")

	vPart, err := jsonPart.GetPart("A")
	assert.NotNil(t, err)
	assert.Nil(t, vPart)

}

func TestJsonPart_GetRawMap(t *testing.T) {
	jsonPart := getJsonPart(SAMPLE_JSON_1)
	rawMap:=jsonPart.GetRawMap()
	status:=rawMap["Status"].(string)
	assert.Equal(t,"SUCCESS", status)
}

func TestJsonPart_GetBooleanF(t *testing.T) {
	jsonPart := getJsonPart(SAMPLE_JSON_1)
	isHigh := jsonPart.GetBooleanF("IsHigh")
	assert.Equal(t,true, isHigh)

	f:=jsonPart.GetBooleanF("Status")
	assert.Equal(t,false,f)
}

func TestJsonPart_GetFloat64F(t *testing.T) {
	jsonPart := getJsonPart(SAMPLE_JSON_1)
	price:=jsonPart.GetFloat64F("Price")
	assert.Equal(t,float64(25.1), price)

	f:=jsonPart.GetFloat64F("Status")
	assert.Equal(t,float64(0),f)
	err:=jsonPart.Error()
	assert.Equal(t, err, fmt.Sprintf(VALUE_TYPE_MISMATCH_ERR_TEMPLATE, jsonPart.key, "Status", "string", "float64"))

}

func TestJsonPart_GetPartF(t *testing.T) {
	jsonPart := getJsonPart(SAMPLE_JSON_1)
	part:=jsonPart.GetPartF("Payload")
	assert.NotNil(t,part)

	f:=jsonPart.GetPartF("Status")
	assert.Nil(t,f)
}

func TestJsonPart_GetStringCastedF(t *testing.T) {
	jsonPart := getJsonPart(SAMPLE_JSON_1)
	str:=jsonPart.GetStringCastedF("Status")
	assert.Equal(t,"SUCCESS", str)

	numOfPplString:=jsonPart.GetStringCastedF("NumOfPeople")
	assert.Equal(t,"999",numOfPplString)
}

func TestJsonPart_GetStringF(t *testing.T) {
	jsonPart := getJsonPart(SAMPLE_JSON_1)
	str:=jsonPart.GetStringF("Status")
	assert.Equal(t,"SUCCESS", str)

	f:=jsonPart.GetStringF("Payload")
	assert.Equal(t,"",f)
}
