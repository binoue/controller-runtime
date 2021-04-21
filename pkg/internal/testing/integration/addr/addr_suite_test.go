package addr_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/binoue/controller-runtime/pkg/envtest/printer"
)

func TestAddr(t *testing.T) {
	t.Parallel()
	RegisterFailHandler(Fail)
	suiteName := "Addr Suite"
	RunSpecsWithDefaultAndCustomReporters(t, suiteName, []Reporter{printer.NewlineReporter{}, printer.NewProwReporter(suiteName)})
}
