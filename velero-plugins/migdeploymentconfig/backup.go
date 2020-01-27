package migdeploymentconfig

import (
	"encoding/json"

	"github.com/konveyor/openshift-migration-plugin/velero-plugins/migcommon"
	"github.com/konveyor/openshift-velero-plugin/velero-plugins/clients"
	appsv1 "github.com/openshift/api/apps/v1"
	"github.com/sirupsen/logrus"
	v1 "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	"github.com/vmware-tanzu/velero/pkg/plugin/velero"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// BackupPlugin is a backup item action plugin for Heptio Ark.
type BackupPlugin struct {
	Log logrus.FieldLogger
}

// AppliesTo returns a backup.ResourceSelector that applies to everything.
func (p *BackupPlugin) AppliesTo() (velero.ResourceSelector, error) {
	return velero.ResourceSelector{
		IncludedResources: []string{"deploymentconfigs"},
	}, nil
}

// Execute sets a custom annotation on the item being backed up.
func (p *BackupPlugin) Execute(item runtime.Unstructured, backup *v1.Backup) (runtime.Unstructured, []velero.ResourceIdentifier, error) {
	// If this is stage don't do anything
	if backup.Annotations[migcommon.MigrateCopyPhaseAnnotation] != "final" || backup.Annotations[migcommon.MigrateQuiesceAnnotation] != "true" {
		return item, nil, nil
	}

	p.Log.Info("[deploymentconfig-backup] Entering DeploymentConfig backup plugin")

	// Convert to DC
	backupDC := appsv1.DeploymentConfig{}
	itemMarshal, _ := json.Marshal(item)
	json.Unmarshal(itemMarshal, &backupDC)

	client, err := clients.OCPAppsClient()
	if err != nil {
		return nil, nil, err
	}
	// Get and update DC on the running cluster to have 0 replicas if this is
	// final migration
	dc, err := client.DeploymentConfigs(backupDC.Namespace).Get(backupDC.Name, metav1.GetOptions{})
	if err != nil {
		return nil, nil, err
	}
	// Scale down DC
	dc.Spec.Replicas = 0

	// Update DC
	dc, err = client.DeploymentConfigs(backupDC.Namespace).Update(dc)
	if err != nil {
		return nil, nil, err
	}

	return item, nil, nil
}
