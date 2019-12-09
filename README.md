# OpenShift Migration Plugin [![Build Status](https://travis-ci.com/fusor/openshift-migration-plugin.svg?branch=master)](https://travis-ci.com/fusor/openshift-migration-plugin) [![Maintainability](https://api.codeclimate.com/v1/badges/95d3aaf8af1cfdd529c4/maintainability)](https://codeclimate.com/github/fusor/ocp-velero-plugin/maintainability)

## Kinds of Plugins

Velero currently supports the following kinds of plugins:

- **Backup Item Action** - performs arbitrary logic on individual items prior to storing them in the backup file.
- **Restore Item Action** - performs arbitrary logic on individual items prior to restoring them in the Kubernetes cluster.

## Building the plugins

To build the plugins, run

```bash
$ make
```

To build the image, run

```bash
$ make container
```

This builds an image tagged as `docker.io/fusor/openshift-migration-plugin`. If you want to specify a
different name, run

```bash
$ make container IMAGE=your-repo/your-name:here
```

## Deploying the plugins

To deploy your plugin image to an Velero server:

1. Make sure your image is pushed to a registry that is accessible to your cluster's nodes.
2. Run `velero plugin add <image>`, e.g. `velero plugin add quay.io/ocpmigrate/migration-plugin`

## Modified version of `dep` needed

At this point, the latest `dep` release does not support dependencies which use go
modules. Since this plugin has dependencies which use go modules, a custom build of
dep based on an open PR which adds go modules support is required. To build dep with
go modules support, run
```
go get -d -u github.com/golang/dep
cd $(go env GOPATH)/src/github.com/golang/dep
git fetch origin +refs/pull/1963/merge
git checkout FETCH_HEAD
go install ./cmd/dep
```
