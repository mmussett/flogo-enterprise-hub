package anotherActivity

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	ASettingString string `md:"aSettingString,required"`
}

type Input struct { 
	AnInputString string `md:"anInputString,required"`
	AnInputBoolean bool                   `md:"anInputBoolean,required"`
	AnInputInteger          int64                  `md:"anInputInteger,required"`
	AnInputDecimal          float64                `md:"anInputDecimal,required"`
	AnInputObject           map[string]interface{} `md:"AnInputObject,required"`
	AnInputObjectWithSchema map[string]interface{} `md:"anInputObjectWithSchema,required"`
	AnInputParameters       map[string]interface{} `md:"anInputParameters,required"`
}

type Output struct {
	AnOutputString           string                 `md:"anOutputString,required"`
	AnOutputBoolean          bool                   `md:"anOutputBoolean,required"`
	AnOutputInteger          int64                  `md:"anOutputInteger,required"`
	AnOutputDecimal          float64                `md:"anOutputDecimal,required"`
	AnOutputObject           map[string]interface{} `md:"anOutputObject,required"`
	AnOutputObjectWithSchema map[string]interface{} `md:"anOutputObjectWithSchema,required"`
}


func (i *Input) FromMap(values map[string]interface{}) error {
	var err error 
	i.AnInputString, err = coerce.ToString(values["anInputString"])
	if err != nil {
		return err
	}
	i.AnInputBoolean, err = coerce.ToBool(values["anInputBoolean"])
	if err != nil {
		return err
	}
	i.AnInputInteger, err = coerce.ToInt64(values["anInputInteger"])
	if err != nil {
		return err
	}
	i.AnInputDecimal, err = coerce.ToFloat64(values["anInputDecimal"])
	if err != nil {
		return err
	}
	i.AnInputObject, err = coerce.ToObject(values["anInputObject"])
	if err != nil {
		return err
	}
	i.AnInputObjectWithSchema, err = coerce.ToObject(values["anInputObject"])
	if err != nil {
		return err
	}
	i.AnInputParameters, err = coerce.ToObject(values["anOutputObjectWithSchema"])
	if err != nil {
		return err
	}

	return nil
}

func (i *Input) ToMap() map[string]interface{} {

	return map[string]interface{}{ 
		"AnInputString":           i.AnInputString,
		"AnInputBoolean":          i.AnInputBoolean,
		"AnInputInteger":          i.AnInputInteger,
		"AnInputDecimal":          i.AnInputDecimal,
		"AnInputObject":           i.AnInputObject,
		"AnInputObjectWithSchema": i.AnInputObjectWithSchema,
		"AnInputParameters":       i.AnInputParameters,
	}


}


func (o *Output) FromMap(values map[string]interface{}) error {
	var err error
	o.AnOutputString, err = coerce.ToString(values["anOutputString"])
	if err != nil {
		return err
	}
	o.AnOutputBoolean, err = coerce.ToBool(values["anOutputBoolean"])
	if err != nil {
		return err
	}
	o.AnOutputInteger, err = coerce.ToInt64(values["anOutputInteger"])
	if err != nil {
		return err
	}
	o.AnOutputDecimal, err = coerce.ToFloat64(values["anOutputDecimal"])
	if err != nil {
		return err
	}
	o.AnOutputObject, err = coerce.ToObject(values["anOutputObject"])
	if err != nil {
		return err
	}
	o.AnOutputObjectWithSchema, err = coerce.ToObject(values["anOutputObjectWithSchema"])
	if err != nil {
		return err
	}

	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"anOutputString":           o.AnOutputString,
		"anOutputBoolean":          o.AnOutputBoolean,
		"anOutputInteger":          o.AnOutputInteger,
		"anOutputDecimal":          o.AnOutputDecimal,
		"anOutputObject":           o.AnOutputObject,
		"anOutputObjectWithSchema": o.AnOutputObjectWithSchema,
	}
}
