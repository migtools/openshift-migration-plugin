package main

import (
	"github.com/fusor/openshift-migration-plugin/velero-plugins/migdeployment"
	"github.com/fusor/openshift-migration-plugin/velero-plugins/migdeploymentconfig"
	"github.com/fusor/openshift-migration-plugin/velero-plugins/migimagestream"
	"github.com/fusor/openshift-migration-plugin/velero-plugins/migimagestreamtag"
	"github.com/fusor/openshift-migration-plugin/velero-plugins/mignamespace"
	"github.com/fusor/openshift-migration-plugin/velero-plugins/migpod"
	"github.com/fusor/openshift-migration-plugin/velero-plugins/migpv"
	"github.com/fusor/openshift-migration-plugin/velero-plugins/migpvc"
	"github.com/fusor/openshift-velero-plugin/velero-plugins/build"
	"github.com/fusor/openshift-velero-plugin/velero-plugins/common"
	"github.com/fusor/openshift-velero-plugin/velero-plugins/daemonset"
	"github.com/fusor/openshift-velero-plugin/velero-plugins/deployment"
	"github.com/fusor/openshift-velero-plugin/velero-plugins/deploymentconfig"
	"github.com/fusor/openshift-velero-plugin/velero-plugins/job"
	"github.com/fusor/openshift-velero-plugin/velero-plugins/replicaset"
	"github.com/fusor/openshift-velero-plugin/velero-plugins/replicationcontroller"
	"github.com/fusor/openshift-velero-plugin/velero-plugins/route"
	"github.com/fusor/openshift-velero-plugin/velero-plugins/serviceaccount"
	"github.com/fusor/openshift-velero-plugin/velero-plugins/statefulset"
	veleroplugin "github.com/heptio/velero/pkg/plugin/framework"
	"github.com/sirupsen/logrus"
)

func main() {
	veleroplugin.NewServer().
		RegisterBackupItemAction("openshift.io/01-common-backup-plugin", newCommonBackupPlugin).
		RegisterRestoreItemAction("openshift.io/01-common-restore-plugin", newCommonRestorePlugin).
		RegisterRestoreItemAction("openshift.io/01-namespace-restore-plugin", newNamespaceRestorePlugin).
		RegisterRestoreItemAction("openshift.io/02-serviceaccount-restore-plugin", newServiceAccountRestorePlugin).
		RegisterBackupItemAction("openshift.io/02-pv-backup-plugin", newPVBackupPlugin).
		RegisterRestoreItemAction("openshift.io/03-pv-restore-plugin", newPVRestorePlugin).
		RegisterRestoreItemAction("openshift.io/03-pvc-restore-plugin", newPVCRestorePlugin).
		RegisterBackupItemAction("openshift.io/03-is-backup-plugin", newImageStreamBackupPlugin).
		RegisterRestoreItemAction("openshift.io/03-is-restore-plugin", newImageStreamRestorePlugin).
		RegisterRestoreItemAction("openshift.io/04-imagestreamtag-restore-plugin", newImageStreamTagRestorePlugin).
		RegisterRestoreItemAction("openshift.io/05-route-restore-plugin", newRouteRestorePlugin).
		RegisterRestoreItemAction("openshift.io/06-build-restore-plugin", newBuildRestorePlugin).
		RegisterRestoreItemAction("openshift.io/07-pod-restore-plugin", newPodRestorePlugin).
		RegisterBackupItemAction("openshift.io/08-deploymentconfig-backup-plugin", newDeploymentConfigBackupPlugin).
		RegisterRestoreItemAction("openshift.io/08-deploymentconfig-restore-plugin", newDeploymentConfigRestorePlugin).
		RegisterRestoreItemAction("openshift.io/09-replicationcontroller-restore-plugin", newReplicationControllerRestorePlugin).
		RegisterRestoreItemAction("openshift.io/10-job-restore-plugin", newJobRestorePlugin).
		RegisterRestoreItemAction("openshift.io/11-daemonset-restore-plugin", newDaemonSetRestorePlugin).
		RegisterRestoreItemAction("openshift.io/12-replicaset-restore-plugin", newReplicaSetRestorePlugin).
		RegisterBackupItemAction("openshift.io/13-deployment-backup-plugin", newDeploymentBackupPlugin).
		RegisterRestoreItemAction("openshift.io/13-deployment-restore-plugin", newDeploymentRestorePlugin).
		RegisterRestoreItemAction("openshift.io/14-statefulset-restore-plugin", newStatefulSetRestorePlugin).
		Serve()
}

func newCommonBackupPlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &common.BackupPlugin{Log: logger}, nil
}

func newCommonRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &common.RestorePlugin{Log: logger}, nil
}

func newNamespaceRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &mignamespace.RestorePlugin{Log: logger}, nil
}

func newBuildRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &build.RestorePlugin{Log: logger}, nil
}

func newDaemonSetRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &daemonset.RestorePlugin{Log: logger}, nil
}

func newDeploymentRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &deployment.RestorePlugin{Log: logger}, nil
}

func newDeploymentConfigRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &deploymentconfig.RestorePlugin{Log: logger}, nil
}

func newJobRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &job.RestorePlugin{Log: logger}, nil
}

func newReplicaSetRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &replicaset.RestorePlugin{Log: logger}, nil
}

func newReplicationControllerRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &replicationcontroller.RestorePlugin{Log: logger}, nil
}

func newRouteRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &route.RestorePlugin{Log: logger}, nil
}

func newServiceAccountRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &serviceaccount.RestorePlugin{Log: logger}, nil
}

func newStatefulSetRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &statefulset.RestorePlugin{Log: logger}, nil
}

func newDeploymentBackupPlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &migdeployment.BackupPlugin{Log: logger}, nil
}

func newDeploymentConfigBackupPlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &migdeploymentconfig.BackupPlugin{Log: logger}, nil
}

func newPodRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &migpod.RestorePlugin{Log: logger}, nil
}

func newPVBackupPlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &migpv.BackupPlugin{Log: logger}, nil
}

func newPVRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &migpv.RestorePlugin{Log: logger}, nil
}

func newPVCRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &migpvc.RestorePlugin{Log: logger}, nil
}

func newImageStreamBackupPlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &migimagestream.BackupPlugin{Log: logger}, nil
}

func newImageStreamRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &migimagestream.RestorePlugin{Log: logger}, nil
}

func newImageStreamTagRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &migimagestreamtag.RestorePlugin{Log: logger}, nil
}
