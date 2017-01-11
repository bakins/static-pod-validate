static-pod-validate
================
static-pod-validate is a simple command line tool for ensuring that Kubernetes
pod manifests are valid.

It was written for static pod manifests as the error messages given by kubelet
are a bit obtuse.

Status
======

Alpha Quality

Building/Installing
===================

`go get bakins/static-pod-validate`

Usage
=====

`static-pod-validate <yaml-or-json-file`

will exit 0 if the manifest is valid.  Non-zreo if invalid, wtih a (hopefully useful)
error message.

License
=======
see [LICENSE](./LICENSE)
