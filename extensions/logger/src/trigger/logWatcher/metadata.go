package logWatcher

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	
	
    Filepath string `md:"filepath,required"`
	
}
type HandlerSettings struct {

}

type Output struct {
	
	
    Outputfield1 string `md:"outputField1,required"`
	
}

func (o *Output) ToMap() map[string]interface{} {
	
	return map[string]interface{}{

		"outputField1": o.Outputfield1,		
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {
	var err error


	// outputField1 string 	
	
	o. Outputfield1, err = coerce.ToString(values["outputField1"])
	
	if err != nil {
		return err
	}

	return nil
}
