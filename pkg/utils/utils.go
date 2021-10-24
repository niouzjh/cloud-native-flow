package utils

import (
	"errors"
	"reflect"
)

func CheckEntryFunc(entryFunc interface{}) error {

	funcType := reflect.TypeOf(entryFunc)
	if funcType.Kind() != reflect.Func {
		return errors.New("entryFunc must be a function")
	}

	// inbound parameters
	numIn := funcType.NumIn()
	// outbound parameters
	numOut := funcType.NumOut()
	if numIn != 2 || numOut != 2 {
		errors.New("entry function must have 2 parameters, and 2 return value")
	}

	inParam_0 := funcType.In(0)
	if inParam_0.Name() != "Context" {
		return errors.New("entry function's first parameter must be Context")
	}

	inParam_1 := funcType.In(1)
	if inParam_1.Kind() != reflect.Struct {
		return errors.New("entry function's second parameter must be struct")
	}

	outParam_0 := funcType.Out(0)
	if outParam_0.Kind() != reflect.Struct {
		return errors.New("entry function's first outpute must be struct")
	}

	outParam_1 := funcType.Out(1)
	if outParam_1.Name() != "error" {
		return errors.New("entry function's second output must be error")
	}

	return nil
}
