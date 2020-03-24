module github.com/konveyor/openshift-migration-plugin

go 1.14

require (
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/containerd/cgroups v0.0.0-20191217135530-f26d9e0c03a9 // indirect
	github.com/containerd/continuity v0.0.0-20191214063359-1097c8bae83b // indirect
	github.com/containers/image/v5 v5.1.1-0.20200108174622-6dbed6e4847c
	github.com/containers/ocicrypt v0.0.0-20191112201931-142388cb70de // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v1.4.2-0.20180522102801-da99009bbb11 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-metrics v0.0.1 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/groupcache v0.0.0-20191027212112-611e8accdfc9 // indirect
	github.com/golang/protobuf v1.3.3-0.20191022195553-ed6926b37a63 // indirect
	github.com/googleapis/gnostic v0.3.1 // indirect
	github.com/gorilla/mux v1.7.3 // indirect
	github.com/hashicorp/go-hclog v0.10.0 // indirect
	github.com/hashicorp/yamux v0.0.0-20190923154419-df201c70410d // indirect
	github.com/konveyor/openshift-velero-plugin v0.0.0-20200303232324-d6e74edcc40f
	github.com/mattn/go-isatty v0.0.11 // indirect
	github.com/openshift/api v0.0.0-20190927182313-d4a64ec2cbd8
	github.com/openshift/client-go v0.0.0-20190813201236-5a5508328169
	github.com/openshift/library-go v0.0.0-20191212223530-61f035b92647 // indirect
	github.com/pquerna/ffjson v0.0.0-20190930134022-aa0246cd15f7 // indirect
	github.com/prometheus/client_golang v1.2.1 // indirect
	github.com/prometheus/client_model v0.1.0 // indirect
	github.com/prometheus/procfs v0.0.8 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/vbauerster/mpb/v4 v4.12.1 // indirect
	github.com/vmware-tanzu/velero v0.0.0-eafaef2a4fb13029d7533e501dc57f9073f4e06b
	go.opencensus.io v0.22.2 // indirect
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553 // indirect
	golang.org/x/oauth2 v0.0.0-20191202225959-858c2ad4c8b6 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	google.golang.org/appengine v1.6.5 // indirect
	google.golang.org/genproto v0.0.0-20191216205247-b31c10ee225f // indirect
	gopkg.in/square/go-jose.v2 v2.4.1 // indirect
	gopkg.in/yaml.v2 v2.2.7 // indirect
	k8s.io/api v0.0.0-20190819141258-3544db3b9e44
	k8s.io/apimachinery v0.0.0-20190817020851-f2f3a405f61d
	k8s.io/client-go v0.0.0-20190819141724-e14f31a72a77
	k8s.io/klog v1.0.0 // indirect
	k8s.io/utils v0.0.0-20191217112158-dcd0c905194b // indirect
)

replace github.com/vmware-tanzu/velero v0.0.0-eafaef2a4fb13029d7533e501dc57f9073f4e06b => github.com/konveyor/velero v0.10.2-0.20200122145352-eafaef2a4fb1
