package  utility

import (
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fn{})
}

type fn struct {
}

// Name returns the name of the function
func (fn) Name() string {
	return "ean13"
}

func (s *fn) GetCategory() string {
	return "utility"
}

// Sig returns the function signature
func (fn) Sig() (paramTypes []data.Type, isVariadic bool) {

    // TODO: Update the return the correctly configured slice 
    // data.TypeString
	// data.TypeInt
	// data.TypeFloat
	// data.TypeBool
	// data.TypeDateTime
	// data.TypeObject


	return []data.Type{data.TypeObject, data.TypeString}, false
}

// Eval executes the function
func (fn) Eval(params ...interface{}) (interface{}, error) {
	obj, err := coerce.ToObject(params[0])
	if err != nil {
		return nil, fmt.Errorf("Unable to coerce [%+v] to object: %s", params[0], err.Error())
	}
	key, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("Unable to coerce the key to string: %s", err.Error())
	}
	val, ok := obj[key]
	if ok {
		return val, nil
	}
	return nil, fmt.Errorf("Unable to get value for key [%s]", key)
}
