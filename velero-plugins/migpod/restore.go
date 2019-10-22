package migpod

import (
	"encoding/json"

	"github.com/fusor/openshift-migration-plugin/velero-plugins/migcommon"
	"github.com/heptio/velero/pkg/plugin/velero"
	"github.com/sirupsen/logrus"
	corev1API "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// RestorePlugin is a restore item action plugin for Velero
type RestorePlugin struct {
	Log logrus.FieldLogger
}

// AppliesTo returns a velero.ResourceSelector that applies to pods
func (p *RestorePlugin) AppliesTo() (velero.ResourceSelector, error) {
	return velero.ResourceSelector{
		IncludedResources: []string{"pods"},
	}, nil
}

// Execute action for the restore plugin for the pod resource
func (p *RestorePlugin) Execute(input *velero.RestoreItemActionExecuteInput) (*velero.RestoreItemActionExecuteOutput, error) {
	p.Log.Info("[pod-restore] Entering Pod restore plugin")

	pod := corev1API.Pod{}
	itemMarshal, _ := json.Marshal(input.Item)
	json.Unmarshal(itemMarshal, &pod)
	p.Log.Infof("[pod-restore] pod: %s", pod.Name)

	newPod := &corev1API.Pod{}
	if _, ok := pod.Labels[migcommon.IncludedInStageBackupLabel]; ok {
		newPod = buildStagePod(&pod, p.Log)
	} else {
		var withoutRestore bool
		var err error
		newPod, withoutRestore, err = migrateApplicationPod(&pod, input, p.Log)
		if err != nil {
			return nil, err
		}
		if withoutRestore {
			return velero.NewRestoreItemActionExecuteOutput(input.Item).WithoutRestore(), nil
		}
	}

	var out map[string]interface{}
	objrec, _ := json.Marshal(&newPod)
	json.Unmarshal(objrec, &out)

	return velero.NewRestoreItemActionExecuteOutput(&unstructured.Unstructured{Object: out}), nil
}
