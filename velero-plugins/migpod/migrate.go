package migpod

import (
	"fmt"

	"github.com/fusor/openshift-migration-plugin/velero-plugins/migcommon"
	"github.com/fusor/openshift-velero-plugin/velero-plugins/common"

	"github.com/heptio/velero/pkg/plugin/velero"
	"github.com/sirupsen/logrus"
	corev1API "k8s.io/api/core/v1"
)

func migrateApplicationPod(pod *corev1API.Pod, input *velero.RestoreItemActionExecuteInput, log logrus.FieldLogger) (*corev1API.Pod, bool, error) {
	log.Infof("[pod-restore] migrating pod %s", pod.Name)

	// ISSUE-61 : removing the node selectors from pods
	// to avoid pod being `unschedulable` on destination
	pod.Spec.NodeSelector = nil

	registry := pod.Annotations[common.RestoreRegistryHostname]
	backupRegistry := pod.Annotations[common.BackupRegistryHostname]
	if registry == "" || backupRegistry == "" {
		return nil, false, fmt.Errorf("failed to find restore registry annotation")
	}
	common.SwapContainerImageRefs(pod.Spec.Containers, backupRegistry, registry, log, input.Restore.Spec.NamespaceMapping)
	common.SwapContainerImageRefs(pod.Spec.InitContainers, backupRegistry, registry, log, input.Restore.Spec.NamespaceMapping)

	ownerRefs, err := common.GetOwnerReferences(input.ItemFromBackup)
	if err != nil {
		return nil, false, err
	}

	// Check if pod has owner Refs and does not have restic backup associated with it
	if len(ownerRefs) > 0 && pod.Annotations[migcommon.ResticBackupAnnotation] == "" {
		log.Infof("[pod-restore] skipping restore of pod %s, has owner references and no restic backup", pod.Name)
		return nil, true, err
	}

	return pod, false, nil
}
