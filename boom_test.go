package boom_test

import (
	. "github.com/jszroberto/boom"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Boom", func() {
	var (
		validManifest Manifest
		boom          Boom
	)
	BeforeEach(func() {
		validManifest = Manifest{
			Jobs: []Job{Job{
				Name:      "brain",
				Instances: 1,
			},
				Job{
					Name:      "cell",
					Instances: 1,
				},
			},
		}
	})

	Context("SetInstances", func() {
		Context("when the job is found", func() {
			It("updates the value", func() {
				err := boom.SetInstances(&validManifest, "cell", 2)
				Expect(err).NotTo(HaveOccurred())
				Expect(validManifest.Jobs[1].Instances).To(Equal(2))
			})
			It("returns an error if instance is not found", func() {

				err := boom.SetInstances(&validManifest, "not-existing", 2)
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("job `not-existing` not found"))

			})
		})
	})
})
