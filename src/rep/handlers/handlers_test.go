package handlers_test

import (
	"code.cloudfoundry.org/diego-release/locket/metrics/helpers/helpersfakes"
	"code.cloudfoundry.org/diego-release/rep"
	"code.cloudfoundry.org/diego-release/rep/evacuation/evacuation_context/fake_evacuation_context"
	"code.cloudfoundry.org/diego-release/rep/handlers"

	executorfakes "code.cloudfoundry.org/executor/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tedsuo/rata"
)

var _ = Describe("New", func() {
	var test_handlers rata.Handlers

	Context("an unsecure server", func() {
		BeforeEach(func() {
			fakeExecutorClient := new(executorfakes.FakeClient)
			fakeEvacuatable := new(fake_evacuation_context.FakeEvacuatable)
			fakeRequestMetrics := new(helpersfakes.FakeRequestMetrics)
			test_handlers = handlers.New(fakeLocalRep, fakeMetricCollector, fakeExecutorClient, fakeEvacuatable, fakeRequestMetrics, logger, false)
		})

		It("has no secure routes", func() {
			for _, route := range rep.RoutesNetworkAccessible {
				Expect(test_handlers[route.Name]).To(BeNil())
			}
		})

		It("has all unsecure routes", func() {
			for _, route := range rep.RoutesLocalhostOnly {
				Expect(test_handlers[route.Name]).NotTo(BeNil())
			}
		})
	})

	Context("a secure server", func() {
		BeforeEach(func() {
			fakeExecutorClient := new(executorfakes.FakeClient)
			fakeEvacuatable := new(fake_evacuation_context.FakeEvacuatable)
			fakeRequestMetrics := new(helpersfakes.FakeRequestMetrics)
			test_handlers = handlers.New(fakeLocalRep, fakeMetricCollector, fakeExecutorClient, fakeEvacuatable, fakeRequestMetrics, logger, true)
		})

		It("has all the secure routes", func() {
			for _, route := range rep.RoutesNetworkAccessible {
				Expect(test_handlers[route.Name]).NotTo(BeNil())
			}
		})

		It("has no unsecure routes", func() {
			for _, route := range rep.RoutesLocalhostOnly {
				Expect(test_handlers[route.Name]).To(BeNil())
			}
		})
	})
})
