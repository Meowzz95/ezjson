package ezjson

import "fmt"

type NonMapError struct {
	parent string
}

func NewNonMapError(parent string) *NonMapError {
	return &NonMapError{parent: parent}
}

func (it *NonMapError) Error() string{
	return fmt.Sprintf("Can't operate on non-map part %s",it.parent)
}

type KeyDoesNotExistError struct {
	parent string
	key string
}

func NewKeyDoesNotExistError(parent string, key string) *KeyDoesNotExistError {
	return &KeyDoesNotExistError{parent: parent, key: key}
}

func (it *KeyDoesNotExistError) Error() string {
	return fmt.Sprintf("'%s.%s' does not exist.", it.parent,it.key)
}

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
	return fmt.Sprintf("'%s.%s' has a type of '%s', but expected '%s'", it.parent,it.key,it.actualType,it.expectedType)
}





