package 

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/carlescere/scheduler"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/trigger"
	"github.com/robfig/cron/v3"
)

var triggerMd = trigger.NewMetadata(&HandlerSettings{})

func init() {
	_ = trigger.Register(&Trigger{}, &TriggerFactory{})
}

type Trigger struct {
}

// Trigger factory
type TriggerFactory struct {
}

// Metadata implements trigger.Trigger.Metadata
func (t *TriggerFactory) Metadata() *trigger.Metadata {
	return triggerMd
}

//New Creates a new trigger instance for a given id
func (t *TriggerFactory) New(config *trigger.Config) (trigger.Trigger, error) {
	return &Trigger{name: config.Id, config: config}, nil
}

// Init implements ext.Trigger.Init
func (t *Trigger) Initialize(ctx trigger.InitContext) error {
	t.logger = ctx.Logger()

	t.logger.Debugf("Initializing %s", t.config.Id)

	for _, handler := range ctx.GetHandlers() {

		handlerSettings := &HandlerSettings{}

		err := metadata.MapToStruct(handler.Settings(), handlerSettings, true)

		if err != nil {
			return fmt.Errorf("error - %s", err.Error())
		}


	}
	return nil
}

func (t *Trigger) Start() error {

	t.logger.Infof("Starting %s...", t.name)

	return nil
}

// Stop implements ext.Trigger.Stop
func (t *Trigger) Stop() error {

	t.logger.Debugf("Stopping %s", t.name)

	return nil
}

//commenting pause resume since its not implemented completely for FlowLimit
// func (t *Trigger) Pause() error {
// 	return nil
// }

// func (t *Trigger) Resume() error {
// 	return nil
// }

