package  barcode

import (
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
)

func init() {
	_ = function.Register(&ean13Func{})
}

type ean13Func struct {
}

// Name returns the name of the function
func (ean13Func) Name() string {
	return "ean13"
}


// GetCategory returns the function category
func (s *ean13Func) GetCategory() string {
	return "barcode"
}

// Sig returns the function signature
func (ean13Func) Sig() (paramTypes []data.Type, isVariadic bool) {

	return []data.Type{data.TypeString}, false
}

// Eval executes the function
func (fn) Eval(params ...interface{}) (interface{}, error) {

	log.RootLogger().Debug("Start of function ean13")

    // Validate parameter count
    if len(params) != 1 {
        return nil, fmt.Errorf("Expected 1 parameters, got %d", len(params))
    }

    // Parameter coercion and validation
    // Coerce parameter 0 (barcode) to string
    barcode, err := coerce.ToString(params[0])
    if err != nil {
        return nil, fmt.Errorf("Unable to coerce parameter 0 (barcode) to string: %s", err.Error())
    }

    // TODO: Implement your function logic here using the coerced parameters
    // Example implementation - replace with your actual logic
    result := "function result" // Replace with actual implementation
    return result, nil
}
