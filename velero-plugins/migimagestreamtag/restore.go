package migimagestreamtag

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/fusor/openshift-velero-plugin/velero-plugins/clients"
	"github.com/fusor/openshift-velero-plugin/velero-plugins/common"
	"github.com/heptio/velero/pkg/plugin/velero"
	imagev1API "github.com/openshift/api/image/v1"
	"github.com/sirupsen/logrus"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// RestorePlugin is a restore item action plugin for Velero
type RestorePlugin struct {
	Log logrus.FieldLogger
}

// AppliesTo returns a velero.ResourceSelector that applies to everything
func (p *RestorePlugin) AppliesTo() (velero.ResourceSelector, error) {
	return velero.ResourceSelector{
		IncludedResources: []string{"imagestreamtags"},
	}, nil
}

func (p *RestorePlugin) Execute(input *velero.RestoreItemActionExecuteInput) (*velero.RestoreItemActionExecuteOutput, error) {
	p.Log.Info("[istag-restore] Entering ImageStreamTag restore plugin")

	imageStreamTag := imagev1API.ImageStreamTag{}
	itemMarshal, _ := json.Marshal(input.Item)
	json.Unmarshal(itemMarshal, &imageStreamTag)
	annotations := imageStreamTag.Annotations
	if annotations == nil {
		annotations = make(map[string]string)
	}

	p.Log.Info(fmt.Sprintf("[istag-restore] Restoring imagestreamtag %s", imageStreamTag.Name))

	internalRegistry := annotations[common.RestoreRegistryHostname]
	if len(internalRegistry) == 0 {
		return nil, errors.New("restore cluster registry not found for annotation \"openshift.io/restore-registry-hostname\"")
	}
	backupInternalRegistry := annotations[common.BackupRegistryHostname]
	if len(backupInternalRegistry) == 0 {
		return nil, errors.New("backup cluster registry not found for annotation \"openshift.io/backup-registry-hostname\"")
	}
	p.Log.Info(fmt.Sprintf("[istag-restore] backup internal registry: %#v", backupInternalRegistry))
	dockerImageReference := imageStreamTag.Image.DockerImageReference
	localImage := common.HasImageRefPrefix(dockerImageReference, backupInternalRegistry)
	if localImage {
		p.Log.Info(fmt.Sprintf("[istag-restore] Local image: %v", dockerImageReference))
	}
	referenceTag := imageStreamTag.Tag != nil && imageStreamTag.Tag.From != nil
	var additionalItems []velero.ResourceIdentifier
	if referenceTag {
		p.Log.Info(fmt.Sprintf("[istag-restore] Reference tag: %v, tag: %v", imageStreamTag.Tag.From.Kind, imageStreamTag.Tag.From.Name))

		// Removing annotations from the tag, to prevent mismatch
		imageStreamTag.Tag.Annotations = nil
		if imageStreamTag.Tag.From.Kind == "ImageStreamTag" {
			p.Log.Info("[istag-restore] ImageStreamTag reference")
			refNamespace := imageStreamTag.Namespace
			if imageStreamTag.Tag.From.Namespace != "" {
				refNamespace = imageStreamTag.Tag.From.Namespace
			}
			p.Log.Info(fmt.Sprintf("[istag-restore] Looking up reference tag: %s/%s", refNamespace, imageStreamTag.Tag.From.Name))
			client, err := clients.ImageClient()
			_, err = client.ImageStreamTags(refNamespace).Get(imageStreamTag.Tag.From.Name, metav1.GetOptions{})
			if err == nil {
				p.Log.Info("[istag-restore] reference tag found in cluster")
			} else {
				p.Log.Info("[istag-restore] reference tag not found in cluster, adding to restore")
				addImageStream := false
				nameSplit := strings.SplitN(imageStreamTag.Tag.From.Name, ":", 2)
				if len(nameSplit) == 2 {
					_, err = client.ImageStreams(refNamespace).Get(nameSplit[0], metav1.GetOptions{})
					if err == nil {
						p.Log.Info("[istag-restore] reference imagestream found in cluster")
					} else {
						p.Log.Info("[istag-restore] reference imagestream not found in cluster, adding to restore")
						addImageStream = true
					}
				}

				additionalItems = []velero.ResourceIdentifier{
					{
						GroupResource: schema.GroupResource{
							Group:    "image.openshift.io",
							Resource: "imagestreamtags",
						},
						Namespace: refNamespace,
						Name:      imageStreamTag.Tag.From.Name,
					},
				}
				if addImageStream {
					additionalItems = append([]velero.ResourceIdentifier{
						{
							GroupResource: schema.GroupResource{
								Group:    "image.openshift.io",
								Resource: "imagestreams",
							},
							Namespace: refNamespace,
							Name:      nameSplit[0],
						},
					}, additionalItems...)
				}
			}
		}
	}

	// Restore the tag if this is a reference tag *or* an external image. Otherwise,
	// image import will create the imagestreamtag automatically.
	if referenceTag || !localImage {
		var out map[string]interface{}
		objrec, _ := json.Marshal(imageStreamTag)
		json.Unmarshal(objrec, &out)
		input.Item.SetUnstructuredContent(out)

		p.Log.Info("[istag-restore] Restoring reference or remote imagestreamtag")
		return &velero.RestoreItemActionExecuteOutput{
			UpdatedItem:     input.Item,
			AdditionalItems: additionalItems,
		}, nil
	}
	p.Log.Info("[istag-restore] Not restoring local imagestreamtag")
	return velero.NewRestoreItemActionExecuteOutput(input.Item).WithoutRestore(), nil
}
