#!/usr/bin/bash
rm -f vendor/github.com/docker/docker/integration-cli/fixtures/https/client-cert.pem
rm -f vendor/github.com/docker/docker/integration-cli/fixtures/https/server-cert.pem
rm -f vendor/github.com/docker/docker/integration-cli/fixtures/https/ca.pem
rm -f vendor/github.com/docker/docker/integration-cli/fixtures/https/client-key.pem
rm -f vendor/github.com/docker/docker/integration-cli/fixtures/https/server-key.pem
[ -e vendor/github.com/docker/docker/integration-cli/fixtures/https ] && rmdir --ignore-fail-on-non-empty vendor/github.com/docker/docker/integration-cli/fixtures/https
[ -e vendor/github.com/docker/docker/integration-cli/fixtures ] && rmdir --ignore-fail-on-non-empty vendor/github.com/docker/docker/integration-cli/fixtures
[ -e vendor/github.com/docker/docker/integration-cli ] && rmdir --ignore-fail-on-non-empty vendor/github.com/docker/docker/integration-cli
rm -f vendor/github.com/docker/docker/project/CONTRIBUTING.md
[ -e vendor/github.com/docker/docker/project ] && rmdir --ignore-fail-on-non-empty vendor/github.com/docker/docker/project
rm -f vendor/github.com/containers/image/signature/fixtures/dir-img-manifest-digest-error/manifest.json
rm -f vendor/github.com/containers/image/signature/fixtures/dir-img-manifest-digest-error/signature-1
rm -f vendor/github.com/containers/image/signature/fixtures/dir-img-mixed/signature-2
rm -f vendor/github.com/containers/image/signature/fixtures/dir-img-mixed/manifest.json
rm -f vendor/github.com/containers/image/signature/fixtures/dir-img-mixed/signature-1
rm -f vendor/github.com/containers/image/signature/fixtures/dir-img-valid-2/manifest.json
rm -f vendor/github.com/containers/image/signature/fixtures/dir-img-valid-2/signature-1
rm -f vendor/github.com/containers/image/signature/fixtures/dir-img-valid/manifest.json
rm -f vendor/github.com/containers/image/signature/fixtures/dir-img-no-manifest/signature-1
rm -f vendor/github.com/containers/image/signature/fixtures/dir-img-unsigned/manifest.json
rm -f vendor/github.com/containers/image/signature/fixtures/dir-img-modified-manifest/signature-1
[ -e vendor/github.com/containers/image/signature/fixtures/dir-img-manifest-digest-error ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/signature/fixtures/dir-img-manifest-digest-error
[ -e vendor/github.com/containers/image/signature/fixtures/dir-img-mixed ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/signature/fixtures/dir-img-mixed
[ -e vendor/github.com/containers/image/signature/fixtures/dir-img-valid ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/signature/fixtures/dir-img-valid
[ -e vendor/github.com/containers/image/signature/fixtures/dir-img-valid-2 ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/signature/fixtures/dir-img-valid-2
[ -e vendor/github.com/containers/image/signature/fixtures/dir-img-no-manifest ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/signature/fixtures/dir-img-no-manifest
[ -e vendor/github.com/containers/image/signature/fixtures/dir-img-unsigned ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/signature/fixtures/dir-img-unsigned
[ -e vendor/github.com/containers/image/signature/fixtures/dir-image-modified-manifest ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/signature/fixtures/dir-img-modified-manifest
[ -e vendor/github.com/containers/image/signature/fixtures ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/signature/fixtures
rm -f vendor/github.com/containers/image/copy/fixtures/Hello.bz2
rm -f vendor/github.com/containers/image/copy/fixtures/Hello.gz
rm -f vendor/github.com/containers/image/copy/fixtures/Hello.xz
rm -f vendor/github.com/containers/image/copy/fixtures/Hello.uncompressed
[ -e vendor/github.com/containers/image/copy/fixtures ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/copy/fixtures
rm -f vendor/github.com/containers/image/manifest/fixtures/schema2-to-schema1-by-docker.json
[ -e vendor/github.com/containers/image/manifest/fixtures ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/manifest/fixtures/
rm -f vendor/github.com/containers/image/pkg/tlsclientconfig/testdata/missing-cert/client-cert-1.key
rm -f vendor/github.com/containers/image/pkg/tlsclientconfig/testdata/unreadable-ca/unreadable.crt
rm -f vendor/github.com/containers/image/pkg/tlsclientconfig/testdata/missing-key/client-cert-1.cert
rm -f vendor/github.com/containers/image/pkg/tlsclientconfig/testdata/unreadable-key/client-cert-1.cert
rm -f vendor/github.com/containers/image/pkg/tlsclientconfig/testdata/unreadable-key/client-cert-1.key
rm -f vendor/github.com/containers/image/pkg/tlsclientconfig/testdata/unreadable-cert/client-cert-1.cert
rm -f vendor/github.com/containers/image/pkg/tlsclientconfig/testdata/unreadable-cert/client-cert-1.key
[ -e vendor/github.com/containers/image/pkg/tlsclientconfig/testdata/missing-cert ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/pkg/tlsclientconfig/testdata/missing-cert
[ -e vendor/github.com/containers/image/pkg/tlsclientconfig/testdata/unreadable-ca ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/pkg/tlsclientconfig/testdata/unreadable-ca
[ -e vendor/github.com/containers/image/pkg/tlsclientconfig/testdata/missing-key ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/pkg/tlsclientconfig/testdata/missing-key
[ -e vendor/github.com/containers/image/pkg/tlsclientconfig/testdata/unreadable-key ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/pkg/tlsclientconfig/testdata/unreadable-key
[ -e vendor/github.com/containers/image/pkg/tlsclientconfig/testdata/unreadable-cert ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/pkg/tlsclientconfig/testdata/unreadable-cert
[ -e vendor/github.com/containers/image/pkg/tlsclientconfig/testdata ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/pkg/tlsclientconfig/testdata
