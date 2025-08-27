package  barcode

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
	return "barcode"
}

// Sig returns the function signature
func (fn) Sig() (paramTypes []data.Type, isVariadic bool) {

	return []data.Type{data.TypeString		
	}, false
}

// Eval executes the function
func (fn) Eval(params ...interface{}) (interface{}, error) {
    // Validate parameter count
    if len(params) != {{ values.arguments | length }} {
        return nil, fmt.Errorf("Expected {{ values.arguments | length }} parameters, got %d", len(params))
    }

    // Parameter coercion and validation
    // Coerce parameter {{ loop.index0 }} ({{ arg.argName }}) to {{ arg.argType }}
    {{ arg.argName }}, err := coerce.ToString(params[{{ loop.index0 }}])
    if err != nil {
        return nil, fmt.Errorf("Unable to coerce parameter {{ loop.index0 }} ({{ arg.argName }}) to string: %s", err.Error())
    }

    // TODO: Implement your function logic here using the coerced parameters
    // Example implementation - replace with your actual logic
    result := "function result" // Replace with actual implementation
    return result, nil
}
