package main

import (
	"github.com/konveyor/openshift-migration-plugin/velero-plugins/migdeployment"
	"github.com/konveyor/openshift-migration-plugin/velero-plugins/migdeploymentconfig"
	"github.com/konveyor/openshift-migration-plugin/velero-plugins/migimagestream"
	"github.com/konveyor/openshift-migration-plugin/velero-plugins/migimagestreamtag"
	"github.com/konveyor/openshift-migration-plugin/velero-plugins/mignamespace"
	"github.com/konveyor/openshift-migration-plugin/velero-plugins/migpod"
	"github.com/konveyor/openshift-migration-plugin/velero-plugins/migpv"
	"github.com/konveyor/openshift-migration-plugin/velero-plugins/migpvc"
	"github.com/konveyor/openshift-migration-plugin/velero-plugins/migsa"
	"github.com/konveyor/openshift-migration-plugin/velero-plugins/migscc"
	"github.com/konveyor/openshift-migration-plugin/velero-plugins/migrolebindings"
	"github.com/konveyor/openshift-migration-plugin/velero-plugins/migclusterrolebindings"
	"github.com/konveyor/openshift-velero-plugin/velero-plugins/build"
	"github.com/konveyor/openshift-velero-plugin/velero-plugins/buildconfig"
	"github.com/konveyor/openshift-velero-plugin/velero-plugins/common"
	"github.com/konveyor/openshift-velero-plugin/velero-plugins/cronjob"
	"github.com/konveyor/openshift-velero-plugin/velero-plugins/daemonset"
	"github.com/konveyor/openshift-velero-plugin/velero-plugins/deployment"
	"github.com/konveyor/openshift-velero-plugin/velero-plugins/deploymentconfig"
	"github.com/konveyor/openshift-velero-plugin/velero-plugins/job"
	"github.com/konveyor/openshift-velero-plugin/velero-plugins/replicaset"
	"github.com/konveyor/openshift-velero-plugin/velero-plugins/replicationcontroller"
	"github.com/konveyor/openshift-velero-plugin/velero-plugins/route"
	"github.com/konveyor/openshift-velero-plugin/velero-plugins/secret"
	"github.com/konveyor/openshift-velero-plugin/velero-plugins/service"
	"github.com/konveyor/openshift-velero-plugin/velero-plugins/serviceaccount"
	"github.com/konveyor/openshift-velero-plugin/velero-plugins/statefulset"
	apisecurity "github.com/openshift/api/security/v1"
	"github.com/sirupsen/logrus"
	veleroplugin "github.com/vmware-tanzu/velero/pkg/plugin/framework"
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
		RegisterRestoreItemAction("openshift.io/15-service-restore-plugin", newServiceRestorePlugin).
		RegisterRestoreItemAction("openshift.io/16-cronjob-restore-plugin", newCronJobRestorePlugin).
		RegisterRestoreItemAction("openshift.io/17-buildconfig-restore-plugin", newBuildConfigRestorePlugin).
		RegisterBackupItemAction("openshift.io/18-serviceaccount-backup-plugin", newServiceAccountBackupPlugin).
		RegisterRestoreItemAction("openshift.io/19-secret-restore-plugin", newSecretRestorePlugin).
		RegisterRestoreItemAction("openshift.io/20-scc-restore-plugin", newSCCRestorePlugin).
		RegisterRestoreItemAction("openshift.io/21-role-bindings-restore-plugin", newRoleBindingRestorePlugin).
		RegisterRestoreItemAction("openshift.io/22-cluster-role-bindings-restore-plugin", newClusterRoleBindingRestorePlugin).
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

func newBuildConfigRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &buildconfig.RestorePlugin{Log: logger}, nil
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

func newCronJobRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &cronjob.RestorePlugin{Log: logger}, nil
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

func newServiceRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &service.RestorePlugin{Log: logger}, nil
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

func newServiceAccountBackupPlugin(logger logrus.FieldLogger) (interface{}, error) {
	saBackupPlugin := &migsa.BackupPlugin{Log: logger}
	saBackupPlugin.UpdatedForBackup = make(map[string]bool)
	// we need to create a dependency between scc and service accounts. Service accounts are listed in SCC's users list.
	saBackupPlugin.SCCMap = make(map[string]map[string][]apisecurity.SecurityContextConstraints)

	return saBackupPlugin, nil
}

func newSecretRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &secret.RestorePlugin{Log: logger}, nil
}

func newSCCRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &migscc.RestorePlugin{Log: logger}, nil
}

func newRoleBindingRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &migrolebindings.RestorePlugin{Log: logger}, nil
}

func newClusterRoleBindingRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &migclusterrolebindings.RestorePlugin{Log: logger}, nil
}
