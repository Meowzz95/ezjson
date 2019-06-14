[![CircleCI](https://circleci.com/gh/Meowzz95/ezjson.svg?style=shield)](https://circleci.com/gh/Meowzz95/ezjson)
[![Coverage Status](https://coveralls.io/repos/github/Meowzz95/ezjson/badge.svg?branch=master)](https://coveralls.io/github/Meowzz95/ezjson?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/Meowzz95/ezjson)](https://goreportcard.com/report/github.com/Meowzz95/ezjson)
# Easy JSON

Miss the way you easily get data from JSON objects? EasyJSON gives you a easier interface to read data from JSON objects. No more messy `map[string]interface{}` from `json.Unmarshal()`

# Install

```
go get -u github.com/Meowzz95/ezjson
```

# API

## Init
```go
jsonMapper:=NewJsonMapper(jsonStr)
jsonPart,err:=jsonMapper.GetJsonPart()
```

`JsonPart` will be your good friend to read data.

## Read data

There are 5 types of data you can get out of a `JsonPart`, namely `bool, float64, string, JsonArray, JsonPart`.

Currently `JsonArray` is under development.

Use this JSON as an example:

```json
{
    "Status": "SUCCESS",
    "Err": null,
    "Info": "",
    "Price": 25.1,
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
```

### Get Number

All numbers are stored in `float64`

```go
price, err:=jsonPart.GetFloat64("Price")
```
You get 25.1 saved in `price`

A special version of `GetFloat64()` is `GetFloat64Or(key string, defaultValue float64)`

If key does not exist or the key does not have a number type, `defaultValue` is returned. No error is returned.


### Get Boolean

Similar to get number, just use
```go
jsonPart.GetBoolean("IsHigh")
```

### Get String

Similar to get number, just use
```go
GetString(key string)
```

A special version of `GetString()` is `GetStringCasted(key string)`, this method will return the data in `string` no matter what the underlying data type is.


### Get Part

Get a nested object

```go
innerPart, _:=jsonPart.GetPart("Payload")
v,_:=innerPart.GetFloat64("ID")
```

And you can perform `Getxxx` on the nested object too.


# Other

PRs are super welcome, please help a newbie :)

## TODO
- JsonArray support
- Chaining API

## License

MIT


