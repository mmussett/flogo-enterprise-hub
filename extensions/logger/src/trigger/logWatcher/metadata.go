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
		
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {
	var err error



	return nil
}
