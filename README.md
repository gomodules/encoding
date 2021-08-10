[![CI](https://github.com/gomodules/encoding/actions/workflows/ci.yml/badge.svg)](https://github.com/gomodules/encoding/actions/workflows/ci.yml)
[![PkgGoDev](https://pkg.go.dev/badge/gomodules.xyz/encoding)](https://pkg.go.dev/gomodules.xyz/encoding)

# encoding

This library has been forked from https://github.com/kubernetes/apimachinery/tree/v0.22.0/pkg/util so that it can easily used in non-k8s related projects.

## Why?

Go `encoding/json` package will automatically convert any number into `float64`, when any json string is unmarshled into a `map[string]interface{}`. This is a common use-case for Kubernetes libraries. To fix this issue, Kubernetes uses their own version of json library (a wrapper) that correctly converts to `int64/float64`. This library is forked from Kubernetes so that it can be used in various places without the whole Kubernetes dependency.
