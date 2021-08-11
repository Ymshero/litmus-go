package environment

import (
	"strconv"

	experimentTypes "github.com/litmuschaos/litmus-go/pkg/generic/node-io-stress/types"
	"github.com/litmuschaos/litmus-go/pkg/types"
	"github.com/litmuschaos/litmus-go/pkg/utils/common"
	clientTypes "k8s.io/apimachinery/pkg/types"
)

//GetENV fetches all the env variables from the runner pod
func GetENV(experimentDetails *experimentTypes.ExperimentDetails) {
	experimentDetails.ExperimentName = common.Getenv("EXPERIMENT_NAME", "node-io-stress")
	experimentDetails.ChaosNamespace = common.Getenv("CHAOS_NAMESPACE", "litmus")
	experimentDetails.EngineName = common.Getenv("CHAOSENGINE", "")
	experimentDetails.ChaosDuration, _ = strconv.Atoi(common.Getenv("TOTAL_CHAOS_DURATION", "120"))
	experimentDetails.RampTime, _ = strconv.Atoi(common.Getenv("RAMP_TIME", "0"))
	experimentDetails.ChaosLib = common.Getenv("LIB", "litmus")
	experimentDetails.AppNS = common.Getenv("APP_NAMESPACE", "")
	experimentDetails.AppLabel = common.Getenv("APP_LABEL", "")
	experimentDetails.AppKind = common.Getenv("APP_KIND", "")
	experimentDetails.ChaosUID = clientTypes.UID(common.Getenv("CHAOS_UID", ""))
	experimentDetails.InstanceID = common.Getenv("INSTANCE_ID", "")
	experimentDetails.ChaosPodName = common.Getenv("POD_NAME", "")
	experimentDetails.FilesystemUtilizationPercentage, _ = strconv.Atoi(common.Getenv("FILESYSTEM_UTILIZATION_PERCENTAGE", ""))
	experimentDetails.FilesystemUtilizationBytes, _ = strconv.Atoi(common.Getenv("FILESYSTEM_UTILIZATION_BYTES", ""))
	experimentDetails.CPU, _ = strconv.Atoi(common.Getenv("CPU", "1"))
	experimentDetails.LIBImage = common.Getenv("LIB_IMAGE", "litmuschaos/go-runner:latest")
	experimentDetails.LIBImagePullPolicy = common.Getenv("LIB_IMAGE_PULL_POLICY", "Always")
	experimentDetails.AuxiliaryAppInfo = common.Getenv("AUXILIARY_APPINFO", "")
	experimentDetails.Delay, _ = strconv.Atoi(common.Getenv("STATUS_CHECK_DELAY", "2"))
	experimentDetails.Timeout, _ = strconv.Atoi(common.Getenv("STATUS_CHECK_TIMEOUT", "180"))
	experimentDetails.TargetNodes = common.Getenv("TARGET_NODES", "")
	experimentDetails.NumberOfWorkers, _ = strconv.Atoi(common.Getenv("NUMBER_OF_WORKERS", "4"))
	experimentDetails.VMWorkers, _ = strconv.Atoi(common.Getenv("VM_WORKERS", "1"))
	experimentDetails.NodesAffectedPerc, _ = strconv.Atoi(common.Getenv("NODES_AFFECTED_PERC", "0"))
	experimentDetails.Sequence = common.Getenv("SEQUENCE", "parallel")
	experimentDetails.TargetContainer = common.Getenv("TARGET_CONTAINER", "")
	experimentDetails.NodeLabel = common.Getenv("NODE_LABEL", "")
	experimentDetails.HostNetwork = Getenvbool("HOST_NETWORK", "false")
	experimentDetails.VolMount = common.Getenv("VOL_MOUNT", "")
}

// Getenv fetch the env and set the default value for bool, if any
func Getenvbool(key string, defaultValue string) bool {
	value := os.Getenv(key)
	var val bool
	if value == "" || defaultValue == "false" {
		val = false
	} else {
		val = true
	}
	return val
}

//InitialiseChaosVariables initialise all the global variables
func InitialiseChaosVariables(chaosDetails *types.ChaosDetails, experimentDetails *experimentTypes.ExperimentDetails) {

	chaosDetails.ChaosNamespace = experimentDetails.ChaosNamespace
	chaosDetails.ChaosPodName = experimentDetails.ChaosPodName
	chaosDetails.ChaosUID = experimentDetails.ChaosUID
	chaosDetails.EngineName = experimentDetails.EngineName
	chaosDetails.ExperimentName = experimentDetails.ExperimentName
	chaosDetails.InstanceID = experimentDetails.InstanceID
	chaosDetails.Timeout = experimentDetails.Timeout
	chaosDetails.Delay = experimentDetails.Delay
	chaosDetails.JobCleanupPolicy = common.Getenv("JOB_CLEANUP_POLICY", "retain")
	chaosDetails.ProbeImagePullPolicy = experimentDetails.LIBImagePullPolicy
	chaosDetails.HostNetwork = experimentDetails.HostNetwork
	chaosDetails.VolMount = experimentDetails.VolMount
}
