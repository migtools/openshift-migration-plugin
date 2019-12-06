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
rm -f vendor/github.com/containers/image/v5/copy/fixtures/Hello.bz2
rm -f vendor/github.com/containers/image/v5/copy/fixtures/Hello.gz
rm -f vendor/github.com/containers/image/v5/copy/fixtures/Hello.xz
rm -f vendor/github.com/containers/image/v5/copy/fixtures/Hello.uncompressed
[ -e vendor/github.com/containers/image/v5/copy/fixtures ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/v5/copy/fixtures
rm -f vendor/github.com/containers/image/v5/manifest/fixtures/schema2-to-schema1-by-docker.json
[ -e vendor/github.com/containers/image/v5/manifest/fixtures ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/v5/manifest/fixtures/
rm -f vendor/github.com/containers/image/v5/pkg/tlsclientconfig/testdata/missing-cert/client-cert-1.key
rm -f vendor/github.com/containers/image/v5/pkg/tlsclientconfig/testdata/unreadable-ca/unreadable.crt
rm -f vendor/github.com/containers/image/v5/pkg/tlsclientconfig/testdata/missing-key/client-cert-1.cert
rm -f vendor/github.com/containers/image/v5/pkg/tlsclientconfig/testdata/unreadable-key/client-cert-1.cert
rm -f vendor/github.com/containers/image/v5/pkg/tlsclientconfig/testdata/unreadable-key/client-cert-1.key
rm -f vendor/github.com/containers/image/v5/pkg/tlsclientconfig/testdata/unreadable-cert/client-cert-1.cert
rm -f vendor/github.com/containers/image/v5/pkg/tlsclientconfig/testdata/unreadable-cert/client-cert-1.key
[ -e vendor/github.com/containers/image/v5/pkg/tlsclientconfig/testdata/missing-cert ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/v5/pkg/tlsclientconfig/testdata/missing-cert
[ -e vendor/github.com/containers/image/v5/pkg/tlsclientconfig/testdata/unreadable-ca ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/v5/pkg/tlsclientconfig/testdata/unreadable-ca
[ -e vendor/github.com/containers/image/v5/pkg/tlsclientconfig/testdata/missing-key ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/v5/pkg/tlsclientconfig/testdata/missing-key
[ -e vendor/github.com/containers/image/v5/pkg/tlsclientconfig/testdata/unreadable-key ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/v5/pkg/tlsclientconfig/testdata/unreadable-key
[ -e vendor/github.com/containers/image/v5/pkg/tlsclientconfig/testdata/unreadable-cert ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/v5/pkg/tlsclientconfig/testdata/unreadable-cert
[ -e vendor/github.com/containers/image/v5/pkg/tlsclientconfig/testdata ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/v5/pkg/tlsclientconfig/testdata
rm -f vendor/github.com/containers/image/v5/signature/fixtures/dir-img-manifest-digest-error/manifest.json
rm -f vendor/github.com/containers/image/v5/signature/fixtures/dir-img-manifest-digest-error/signature-1
[ -e  vendor/github.com/containers/image/v5/signature/fixtures/dir-img-manifest-digest-error ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/v5/signature/fixtures/dir-img-manifest-digest-error
rm -f vendor/github.com/containers/image/v5/signature/fixtures/dir-img-mixed/manifest.json
rm -f vendor/github.com/containers/image/v5/signature/fixtures/dir-img-mixed/signature-1
rm -f vendor/github.com/containers/image/v5/signature/fixtures/dir-img-mixed/signature-2
[ -e vendor/github.com/containers/image/v5/signature/fixtures/dir-img-mixed/ ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/v5/signature/fixtures/dir-img-mixed/
rm -f vendor/github.com/containers/image/v5/signature/fixtures/dir-img-valid-2/manifest.json
rm -f vendor/github.com/containers/image/v5/signature/fixtures/dir-img-valid-2/signature-1
[ -e vendor/github.com/containers/image/v5/signature/fixtures/dir-img-valid-2/ ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/v5/signature/fixtures/dir-img-valid-2/
rm -f vendor/github.com/containers/image/v5/signature/fixtures/dir-img-valid/manifest.json
[ -e vendor/github.com/containers/image/v5/signature/fixtures/dir-img-valid/ ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/v5/signature/fixtures/dir-img-valid/
rm -f vendor/github.com/containers/image/v5/signature/fixtures/dir-img-no-manifest/signature-1
[ -e vendor/github.com/containers/image/v5/signature/fixtures/dir-img-no-manifest/ ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/v5/signature/fixtures/dir-img-no-manifest/
rm -f vendor/github.com/containers/image/v5/signature/fixtures/dir-img-unsigned/manifest.json
[ -e vendor/github.com/containers/image/v5/signature/fixtures/dir-img-unsigned/ ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/v5/signature/fixtures/dir-img-unsigned/
rm -f vendor/github.com/containers/image/v5/signature/fixtures/dir-img-modified-manifest/signature-1
[ -e vendor/github.com/containers/image/v5/signature/fixtures/dir-img-modified-manifest/ ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/v5/signature/fixtures/dir-img-modified-manifest/
[ -e vendor/github.com/containers/image/v5/signature/fixtures/ ] && rmdir --ignore-fail-on-non-empty vendor/github.com/containers/image/v5/signature/fixtures/
