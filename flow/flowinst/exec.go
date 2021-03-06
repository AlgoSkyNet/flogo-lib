package flowinst

import (
	"github.com/TIBCOSoftware/flogo-lib/flow/support"
)

// ExecOptions are optional Patch & Interceptor to be used during instance execution
type ExecOptions struct {
	Patch       *support.Patch
	Interceptor *support.Interceptor
}

// IDGenerator generates IDs for flow instances
type IDGenerator interface {

	//NewFlowInstanceID generate a new instance ID
	NewFlowInstanceID() string
}

// applyExecOptions applies any execution options to the flow instance
func applyExecOptions(instance *Instance, execOptions *ExecOptions) {

	if execOptions != nil {

		if execOptions.Patch != nil {
			log.Infof("Instance [%s] has patch", instance.ID())
			instance.Patch = execOptions.Patch
			instance.Patch.Init()
		}

		if execOptions.Interceptor != nil {
			log.Infof("Instance [%s] has interceptor", instance.ID)
			instance.Interceptor = execOptions.Interceptor
			instance.Interceptor.Init()
		}
	}
}
