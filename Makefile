test-all:
	TRACE=1 ./test.sh

# Detect go environment to build tools into env specific directories.
# This is helpful for those running tools locally and within containers across different OS/architechture combinations.
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
SCRIPT_DIR=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
OUTPUT_DIR=$(SCRIPT_DIR)/_output/bin/$(GOOS)/$(GOARCH)

.PHONY:applyconfiguration-gen
applyconfiguration-gen: $(OUTPUT_DIR)/applyconfiguration-gen

$(OUTPUT_DIR)/applyconfiguration-gen:
	go build -mod=vendor -o $(OUTPUT_DIR)/applyconfiguration-gen ./vendor/k8s.io/code-generator/cmd/applyconfiguration-gen

generate-applyconfigs: applyconfiguration-gen
	$(OUTPUT_DIR)/applyconfiguration-gen \
	-v 3 \
	--output-package sigs.k8s.io/controller-tools/pkg/applyconfigurations/testdata/cronjob/applyconfigurations \
	--input-dirs sigs.k8s.io/controller-tools/pkg/applyconfigurations/testdata/cronjob \
	--go-header-file ./hack/empty.txt \
	--external-applyconfigurations k8s.io/api/batch/v1.JobSpec:k8s.io/client-go/applyconfigurations/batch/v1 \
	--external-applyconfigurations k8s.io/api/batch/v1beta1.JobTemplateSpec:k8s.io/client-go/applyconfigurations/batch/v1beta1 \
	sigs.k8s.io/controller-tools/pkg/applyconfigurations/testdata/cronjob
