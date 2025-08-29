package myActivity

import (
	"fmt"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
)

func init() {
	
	//err := activity.Register(&Activity{}) 
	//err := activity.Register(&Activity{}, New) to create instances using factory method 'New'
	
	err := activity.Register(&Activity{}, New)
	if err != nil {
		log.RootLogger().Error(err)
	}
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})
var activityLog = log.ChildLogger(log.RootLogger(), "logger-myActivity")

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {

	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}
	ctx.Logger().Debugf("Setting: %v", settings.Fieldname1)
	ctx.Logger().Debugf("Setting: %v", settings.Fieldname2)

	act := &Activity{logger: log.ChildLogger(ctx.Logger(), "logger-myActivity"), activityName: "myActivity"}

	return act, nil
}

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
	logger       log.Logger
	activityName string
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Cleanup method
func (a *Activity) Cleanup() error {

	return nil
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	activityLog.Debugf("Executing Activity [%s] ", ctx.Name())

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return false, fmt.Errorf("Error while getting input object: %s", err.Error())
	}
	ctx.Logger().Debugf("Input: %v", input.Inputfield1)
	ctx.Logger().Debugf("Input: %v", input.Inputfield2)

    // TODO: Implement your activity logic here

	output := &Output{}
	ctx.Logger().Debugf("Output: %v", output.Outputfield1)
	ctx.Logger().Debugf("Output: %v", output.Outputfield2)

	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	activityLog.Debugf("Execution of Activity [%s] " + context.Name() + " completed")
	return true, nil
}