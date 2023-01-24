//go:build tools
// +build tools

package tools

import (
	_ "k8s.io/code-generator/cmd/applyconfiguration-gen"
	_ "sigs.k8s.io/controller-tools/pkg/applyconfigurations/testdata/cronjob" // Hack because this is a separate module, the context needs to be run from the module for import normally
)
