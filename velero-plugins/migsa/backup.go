package migsa

import (
	"encoding/json"
	"strings"

	"github.com/fusor/openshift-migration-plugin/velero-plugins/migcommon"
	v1 "github.com/heptio/velero/pkg/apis/velero/v1"
	"github.com/heptio/velero/pkg/plugin/velero"
	apisecurity "github.com/openshift/api/security/v1"
	security "github.com/openshift/client-go/security/clientset/versioned/typed/security/v1"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
)

// BackupPlugin is a backup item action plugin for Heptio Velero.
type BackupPlugin struct {
	Log              logrus.FieldLogger
	SCCMap           map[string]map[string][]apisecurity.SecurityContextConstraints
	UpdatedForBackup map[string]bool
}

// AppliesTo returns a velero.ResourceSelector that applies to everything.
func (p *BackupPlugin) AppliesTo() (velero.ResourceSelector, error) {
	return velero.ResourceSelector{
		IncludedResources: []string{"serviceaccounts"},
	}, nil
}

// This should be moved to clients package in future
var securityClient *security.SecurityV1Client
var securityClientError error

// Execute copies local registry images into migration registry
func (p *BackupPlugin) Execute(item runtime.Unstructured, backup *v1.Backup) (runtime.Unstructured, []velero.ResourceIdentifier, error) {
	migrating := migcommon.IsMigBackup(backup)
	if !migrating {
		return item, nil, nil
	}
	p.Log.Info("[serviceaccount-backup] Entering ServiceAccount backup plugin")

	if !p.UpdatedForBackup[backup.Name] {
		err := p.UpdateSCCMap()
		if err != nil {
			return nil, nil, err
		}

		p.UpdatedForBackup[backup.Name] = true
	}

	serviceAccount := corev1.ServiceAccount{}
	itemMarshal, _ := json.Marshal(item)
	json.Unmarshal(itemMarshal, &serviceAccount)

	var additionalItems []velero.ResourceIdentifier

	if p.SCCMap[serviceAccount.Namespace] == nil {
		return item, additionalItems, nil
	}

	for _, scc := range p.SCCMap[serviceAccount.Namespace][serviceAccount.Name] {
		p.Log.Infof("Adding security context constraint - %s as additional item for service account - %s in namespace - %s", scc.Name,
			serviceAccount.Name, serviceAccount.Namespace)
		additionalItems = append(additionalItems, velero.ResourceIdentifier{
			Name:          scc.Name,
			GroupResource: schema.GroupResource{Group: "security.openshift.io", Resource: "securitycontextconstraints"},
		})
	}

	return item, additionalItems, nil
}

// UpdateSCCMap fill scc map with service account as key and SCCs slice as value
func (p *BackupPlugin) UpdateSCCMap() error {
	client, err := SecurityClient()
	if err != nil {
		return err
	}

	sccs, err := client.SecurityContextConstraints().List(metav1.ListOptions{})

	for _, scc := range sccs.Items {
		for _, user := range scc.Users {
			// Service account username format role:serviceaccount:namespace:serviceaccountname
			splitUsername := strings.Split(user, ":")
			if len(splitUsername) <= 2 { // safety check
				continue
			}

			// if second element is serviceaccount then last element is serviceaccountname
			if splitUsername[1] == "serviceaccount" {
				namespace := splitUsername[2]
				saName := splitUsername[3]

				if p.SCCMap[namespace] == nil {
					p.SCCMap[namespace] = make(map[string][]apisecurity.SecurityContextConstraints)
				}

				if p.SCCMap[namespace][saName] == nil {
					p.SCCMap[namespace][saName] = make([]apisecurity.SecurityContextConstraints, 0)
				}

				p.SCCMap[namespace][saName] = append(p.SCCMap[namespace][saName], scc)
			}
		}
	}

	return nil
}

// This should be moved to clients package in future

// SecurityClient returns an openshift AppsV1Client
func SecurityClient() (*security.SecurityV1Client, error) {
	if securityClient == nil && securityClientError == nil {
		securityClient, securityClientError = newSecurityClient()
	}
	return securityClient, securityClientError
}

// This should be moved to clients package in future
func newSecurityClient() (*security.SecurityV1Client, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	client, err := security.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return client, nil
}
