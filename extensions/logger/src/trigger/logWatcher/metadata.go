package  logWatcher

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	
	
    Filepath string `md:"filepath,required"`
	
}
type HandlerSettings struct {
	
	
    Handlerfield1 string `md:"handlerField1,required"`
	
}

type Output struct {
	
	
    Changes string `md:"changes,required"`
	
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"output": o.Output,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {
	var err error
	o.Output, err = coerce.ToObject(values["output"])
	if err != nil {
		return err
	}
	return nil
}
