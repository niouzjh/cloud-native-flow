package cnf

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

// steps
type Step struct {
	// 步骤名称
	Name string
	// 注释，步骤详细描述
	Comment string
	// 步骤对应的函数
	Func interface{}
	// 子任务流
	SubFlow Flow

	// runtime field
	// 输出结果
	OutputParam map[string]interface{}
}

// flows
type Flow struct {
	// 任务流名称
	Name string
	// 任务流功能描述
	Comment string
	// 任务流包含的步骤
	Steps []Step

	// runtime filed
	// 任务流初始化参数
	InitialParam map[string]interface{}
}

// flow center
type CloudNativeFlow struct {
	// 持久化存储
	Storage string
	// 已注册的任务流
	Flows []Flow
}

// The global cloud native flow
var CNF CloudNativeFlow

func (f *CloudNativeFlow) getValueByNameFromPreviousSteps(stepIndex int, paramName string) (value interface{}, err error) {
	stepLength := len(f.Flows)
	if stepIndex >= stepLength {
		return nil, errors.New("error")
	}
	for i := stepIndex; i >= 0; i-- {
		value, ok := f.steps.OutputParam[paramName]
		if ok {
			return value, nil
		}
	}
	// find from flow initial parameters

	return nil, errors.New("parameter not found")
}

// 运行一个任务流
func (cnf *CloudNativeFlow) Run(flow Flow) error {

	for index, step := range flow.Steps {
		// todo:check if step finished
		funcType := reflect.TypeOf(step.Func)
		funcValue := reflect.ValueOf(step.Func)

		// constuct input parameter
		in := make([]reflect.Value, funcType.NumIn())

		ctx := context.TODO()
		in[0] = reflect.ValueOf(ctx)

		paramType := funcType.In(1)
		in[1] = reflect.New(paramType).Elem()
		for paramIndex := 0; paramIndex < in[1].NumField(); paramIndex++ {
			fieldType := paramType.Field(paramIndex)
			fieldValue := in[1].Field(paramIndex)

			value, err := cnf.getValueByNameFromPreviousSteps(stepIndex, fieldType.Name)
			if err != nil {
				fmt.Println("error")
			}

			// if fieldType is pointer of value's Type
			// set values's pointer to fieldValue
			if fieldType.Type.Kind() == reflect.Ptr &&
				(fieldType.Type.Elem().Kind() == reflect.Struct || fieldType.Type.Elem().Kind() == reflect.Map) {
				fieldv := reflect.New(fieldType.Type.Elem())
				str, err := json.Marshal(value)
				if err != nil {
					return errors.New("marshal error")
				}
				//fmt.Println(reflect.TypeOf(fieldv.Addr()))
				if err = json.Unmarshal([]byte(str), fieldv.Interface()); err != nil {
					return errors.New("unmarshal error")
				}
				fieldValue.Set(fieldv)
			} else if fieldType.Type.Kind() == reflect.Struct || fieldType.Type.Kind() == reflect.Map {
				fieldv := reflect.New(fieldType.Type)
				str, err := json.Marshal(value)
				if err != nil {
					return errors.New("marshal error")
				}
				if err = json.Unmarshal([]byte(str), fieldv.Interface()); err != nil {
					return errors.New("unmarshal error")
				}
				fieldValue.Set(fieldv.Elem())
			} else if fieldType.Type == reflect.PtrTo(reflect.TypeOf(value)) {
				if reflect.ValueOf(value).CanAddr() {
					fieldValue.Set(reflect.ValueOf(value).Addr())
				} else {
					// FIXME
				}
			} else {
				fieldValue.Set(reflect.ValueOf(value))
			}

			// call step's func
			result := funcValue.Call(in)

			if result[1].Type().Name() != "error" {
				return errors.New("second result of step function must be error")
			}
			result_error := result[1].Interface()
			if result_error != nil {
				// TODO set step status to error
				// TODO set flow status to error
				result_error = result[1].Interface().(error)
				fmt.Println("step error:", result_error)
				return result[1].Interface().(error)
			}
			// TODO update step status

			// deal with result
			for resultIndex := 0; resultIndex < result[0].NumMethod(); resultIndex++ {
				fieldType := result[0].Type().Field(resultIndex)
				fieldValue := result[0].Field(resultIndex)

				if !fieldValue.CanInterface() {
					fmt.Println("unexported output field ignored")
					continue
				}
				//steps.OutputParam[fieldType.Name]=fieldValue.Interface()
			}
			// TODO update step status
		}

		// TODO update flow status
	}
	return nil
}

// 开始任务监听
func (cnf *CloudNativeFlow) Start() error {

	return nil
}
