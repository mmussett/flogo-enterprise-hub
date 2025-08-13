package anotherActivity

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	ASettingString string `md:"aSettingString,required"`
}

type Input struct {
}

func (i *Input) FromMap(values map[string]interface{}) error {
	var err error

	return nil
}

func (i *Input) ToMap() map[string]interface{} {

	return map[string]interface{}{

}

type Output struct {
}

func (o *Output) FromMap(values map[string]interface{}) error {
	var err error

	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
	}
}
