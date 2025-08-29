package  myActivity

import "github.com/project-flogo/core/data/coerce"

type Settings struct {	
	
    Fieldname1 string `md:"fieldName1,required"`
		
	
    Fieldname2 float64 `md:"fieldName2,required"`
	
}

type Input struct {	
	
    Inputfield1 string `md:"inputField1,required"`
		
	
    Inputfield2 float64 `md:"inputField2,required"`
	
}

type Output struct {	
	
    Outputfield1 string `md:"outputField1,required"`
		
	
    Outputfield2 bool `md:"outputField2,required"`
	
}


func (i *Input) FromMap(values map[string]interface{}) error {
	var err error

	// inputField1 string 
	
	i. Inputfield1, err = coerce.ToString(values["inputField1"])
	
	if err != nil {
		return err
	}

	// inputField2 number 
	
	i.Inputfield2, err = coerce.ToFloat64(values["inputField2"])
	
	if err != nil {
		return err
	}

	return nil
}

func (i *Input) ToMap() map[string]interface{} {

	return map[string]interface{}{
		"inputField1": i.Inputfield1,
		"inputField2": i.Inputfield2,
	}

}


func (o *Output) FromMap(values map[string]interface{}) error {
	var err error
	// outputField1 string 	
	
	o. Outputfield1, err = coerce.ToString(values["outputField1"])
	
	if err != nil {
		return err
	}
	// outputField2 boolean 	
	
	o.Outputfield2, err = coerce.ToBool(values["outputField2"])
	
	if err != nil {
		return err
	}

	return nil
}

func (o *Output) ToMap() map[string]interface{} {

	return map[string]interface{}{
		"outputField1": o.Outputfield1,
		"outputField2": o.Outputfield2,		
	}

}
