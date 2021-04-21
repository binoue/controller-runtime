package integration_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/binoue/controller-runtime/pkg/envtest/printer"
)

func TestIntegration(t *testing.T) {
	t.Parallel()
	RegisterFailHandler(Fail)
	suiteName := "Integration Framework Unit Tests"
	RunSpecsWithDefaultAndCustomReporters(t, suiteName, []Reporter{printer.NewlineReporter{}, printer.NewProwReporter(suiteName)})
}
