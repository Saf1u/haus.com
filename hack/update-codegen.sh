#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail




 /Users/safwanahmed/go/pkg/mod/k8s.io/code-generator@v0.31.3/kube_codegen.sh "all" \
   github.com/saf1u/haus.com/pkg/generated \
   github.com/saf1u/haus.com/apis \
  "haus.com:v1" \
  --output-base ".." \
  --go-header-file ../hack/boilerplate.go.txt

