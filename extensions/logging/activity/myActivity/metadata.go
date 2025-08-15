package myActivity

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	ASettingString string `md:"aSettingString,required"`
}

type Input struct { 
	AnInputString string `md:"anInputString,required"`
}

type Output struct {
	AnOutputString           string                 `md:"anOutputString,required"`
}


func (i *Input) FromMap(values map[string]interface{}) error {
	var err error 
	i.AnInputString, err = coerce.ToString(values["anInputString"])
	if err != nil {
		return err
	}

	return nil
}

func (i *Input) ToMap() map[string]interface{} {

	return map[string]interface{}{ 
		"AnInputString":           i.AnInputString,
	}


}


func (o *Output) FromMap(values map[string]interface{}) error {
	var err error
	o.AnOutputString, err = coerce.ToString(values["anOutputString"])
	if err != nil {
		return err
	}

	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"anOutputString":           o.AnOutputString,
	}
}
