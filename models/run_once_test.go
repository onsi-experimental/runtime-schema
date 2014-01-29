package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudfoundry-incubator/runtime-schema/models"
)

var _ = Describe("RunOnce", func() {
	var runOnce RunOnce

	runOncePayload := `{"guid":"some-guid","executor_id":"executor","container_handle":"17fgsafdfcvc","failed":true,"failure_reason":"because i said so"}`

	BeforeEach(func() {
		runOnce = RunOnce{
			Guid:            "some-guid",
			ExecutorID:      "executor",
			ContainerHandle: "17fgsafdfcvc",
			Failed:          true,
			FailureReason:   "because i said so",
		}
	})

	Describe("ToJSON", func() {
		It("should JSONify", func() {
			json := runOnce.ToJSON()
			Ω(string(json)).Should(Equal(runOncePayload))
		})
	})

	Describe("NewRunOnceFromJSON", func() {
		It("returns a RunOnce with correct fields", func() {
			decodedRunOnce, err := NewRunOnceFromJSON([]byte(runOncePayload))
			Ω(err).ShouldNot(HaveOccurred())

			Ω(decodedRunOnce).Should(Equal(runOnce))
		})

		Context("with an invalid payload", func() {
			It("returns the error", func() {
				decodedRunOnce, err := NewRunOnceFromJSON([]byte("butts lol"))
				Ω(err).Should(HaveOccurred())

				Ω(decodedRunOnce).Should(BeZero())
			})
		})
	})
})
