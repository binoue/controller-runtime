package cache_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"k8s.io/client-go/rest"

	"github.com/binoue/controller-runtime/pkg/cache"
	"github.com/binoue/controller-runtime/pkg/client/apiutil"
	"github.com/binoue/controller-runtime/pkg/manager"
)

var _ = Describe("informerCache", func() {
	It("should not require LeaderElection", func() {
		cfg := &rest.Config{}

		mapper, err := apiutil.NewDynamicRESTMapper(cfg, apiutil.WithLazyDiscovery)
		Expect(err).ToNot(HaveOccurred())

		c, err := cache.New(cfg, cache.Options{Mapper: mapper})
		Expect(err).ToNot(HaveOccurred())

		leaderElectionRunnable, ok := c.(manager.LeaderElectionRunnable)
		Expect(ok).To(BeTrue())
		Expect(leaderElectionRunnable.NeedLeaderElection()).To(BeFalse())
	})
})
