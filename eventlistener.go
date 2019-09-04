package ext

import (
	core "github.com/project-flogo/core/engine/event"
	"github.com/project-flogo/core/support/log"
	flow "github.com/project-flogo/flow/support/event"
)

func init() {
	core.RegisterListener("extTestListener", &SmapleListener{name: "extTestListener"}, []string{flow.FlowEventType, flow.TaskEventType})
}

type SmapleListener struct {
	name string
}

func (ls *SmapleListener) Name() string {
	return ls.name
}

func (ls *SmapleListener) HandleEvent(evt *core.Context) error {
	// Handle flow and task events
	switch t := evt.GetEvent().(type) {
	case flow.FlowEvent:
		log.RootLogger().Infof("Name: %s, ID: %s, Status: %s ", t.FlowName(), t.FlowID(), t.FlowStatus())
	case flow.TaskEvent:
		log.RootLogger().Infof("Name: %s, FID: %s, Status: %s", t.TaskName(), t.FlowID(), t.TaskStatus())
	}
	return nil
}
