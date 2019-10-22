package migpod

import (
	"strconv"
	"strings"

	"github.com/fusor/openshift-migration-plugin/velero-plugins/migcommon"
	"github.com/sirupsen/logrus"
	corev1API "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func buildStagePod(pod *corev1API.Pod, log logrus.FieldLogger) *corev1API.Pod {
	log.Infof("[pod-restore] creating stage pod for %s", pod.Name)

	// Map of Restic volumes.
	resticVolumes := map[string]bool{}
	for _, name := range strings.Split(pod.Annotations[migcommon.ResticBackupAnnotation], ",") {
		resticVolumes[name] = true
	}
	// Base pod.
	newPod := corev1API.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       pod.Kind,
			APIVersion: pod.APIVersion,
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: pod.Namespace,
			Name:      pod.Name + "-" + "stage",
			Annotations: map[string]string{
				migcommon.ResticBackupAnnotation: pod.Annotations[migcommon.ResticBackupAnnotation],
			},
			Labels: pod.Labels,
		},
		Spec: corev1API.PodSpec{
			Containers:                   []corev1API.Container{},
			Volumes:                      []corev1API.Volume{},
			SecurityContext:              pod.Spec.SecurityContext,
			ServiceAccountName:           pod.Spec.ServiceAccountName,
			AutomountServiceAccountToken: pod.Spec.AutomountServiceAccountToken,
		},
	}

	// Add volumes.
	for _, volume := range pod.Spec.Volumes {
		if _, found := resticVolumes[volume.Name]; found {
			newPod.Spec.Volumes = append(newPod.Spec.Volumes, volume)
		}
	}

	// Add containers.
	for i, container := range pod.Spec.Containers {
		stageContainer := corev1API.Container{
			Name:         "sleep-" + strconv.Itoa(i),
			Image:        "registry.access.redhat.com/rhel7",
			Command:      []string{"sleep"},
			Args:         []string{"infinity"},
			VolumeMounts: []corev1API.VolumeMount{},
		}
		for _, mount := range container.VolumeMounts {
			if _, found := resticVolumes[mount.Name]; found {
				stageContainer.VolumeMounts = append(stageContainer.VolumeMounts, mount)
			}
		}
		newPod.Spec.Containers = append(newPod.Spec.Containers, stageContainer)
	}

	return &newPod
}
