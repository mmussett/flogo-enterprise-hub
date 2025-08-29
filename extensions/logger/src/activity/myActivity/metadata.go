package  myActivity

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	
	
    Fieldname1 string `md:"fieldName1,required"`
	
}

type Input struct {
	
	
    Inputfield1 string `md:"inputField1,required"`
	
}

type Output struct {
	
	
    Outputfield1 string `md:"outputField1,required"`
	
}


func (i *Input) FromMap(values map[string]interface{}) error {
	var err error

	
	
	i. Inputfield1, err = coerce.ToString(values["inputField1"])
	
	if err != nil {
		return err
	}

	return nil
}

func (i *Input) ToMap() map[string]interface{} {

	return map[string]interface{}{

		"Inputfield1": i.inputField1,
	}

}


func (o *Output) FromMap(values map[string]interface{}) error {
	var err error

	
	
	o. Outputfield1, err = coerce.ToString(values["outputField1"])
	
	if err != nil {
		return err
	}

	return nil
}

func (o *Output) ToMap() map[string]interface{} {


	return map[string]interface{}{

		"outputField1": o.Outputfield1,		
	}

}
