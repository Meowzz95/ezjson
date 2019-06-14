package ezjson

import "fmt"

const NON_MAP_ERROR_TEMPLATE = "Can't operate on non-map part %s"
type NonMapError struct {
	parent string
}

func NewNonMapError(parent string) *NonMapError {
	return &NonMapError{parent: parent}
}

func (it *NonMapError) Error() string{
	return fmt.Sprintf(NON_MAP_ERROR_TEMPLATE,it.parent)
}

const KEY_DOES_NOT_EXIST_TEMPLATE = "'%s.%s' does not exist."
type KeyDoesNotExistError struct {
	parent string
	key string
}

func NewKeyDoesNotExistError(parent string, key string) *KeyDoesNotExistError {
	return &KeyDoesNotExistError{parent: parent, key: key}
}

func (it *KeyDoesNotExistError) Error() string {
	return fmt.Sprintf(KEY_DOES_NOT_EXIST_TEMPLATE, it.parent,it.key)
}

const VALUE_TYPE_MISMATCH_ERR_TEMPLATE = "'%s.%s' has a type of '%s', but expected '%s'"
type ValueTypeMismatchError struct{
	parent string
	key string
	expectedType string
	actualType string
}

func NewValueTypeMismatchError(parent string, key string, expectedType string, actualType string) *ValueTypeMismatchError {
	return &ValueTypeMismatchError{parent: parent, key: key, expectedType: expectedType, actualType: actualType}
}

func (it *ValueTypeMismatchError) Error() string {
	return fmt.Sprintf(VALUE_TYPE_MISMATCH_ERR_TEMPLATE, it.parent,it.key,it.actualType,it.expectedType)
}





