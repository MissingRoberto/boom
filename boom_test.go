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
				Instances: 2,
			},
				Job{
					Name:      "cell",
					Instances: 20,
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
		})
		Context("when the job is not found", func() {
			It("returns an error", func() {
				err := boom.SetInstances(&validManifest, "not-existing", 2)
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("job `not-existing` not found"))

			})
		})
	})

	Context("ScaleInstances", func() {
		Context("when the job is found", func() {
			Context("when the value is not round", func() {
				It("don't update the value", func() {
					err := boom.ScaleInstances(&validManifest, "cell", 1)
					Expect(err).NotTo(HaveOccurred())
					Expect(validManifest.Jobs[1].Instances).To(Equal(20))
				})
			})
			It("decreases the value", func() {
				err := boom.ScaleInstances(&validManifest, "cell", 0.5)
				Expect(err).NotTo(HaveOccurred())
				Expect(validManifest.Jobs[1].Instances).To(Equal(10))
			})
			It("increases the value", func() {
				err := boom.ScaleInstances(&validManifest, "cell", 2)
				Expect(err).NotTo(HaveOccurred())
				Expect(validManifest.Jobs[1].Instances).To(Equal(40))
			})
			Context("when the value is not round", func() {
				It("updates the value", func() {
					err := boom.ScaleInstances(&validManifest, "cell", 1.5)
					Expect(err).NotTo(HaveOccurred())
					Expect(validManifest.Jobs[1].Instances).To(Equal(30))
				})
			})
			Context("when the factor is 0", func() {
				It("returns an error", func() {
					err := boom.ScaleInstances(&validManifest, "cell", 0)
					Expect(err).To(MatchError("factor 0 is not permitted"))
				})
			})
			Context("when the factor is too low to modify a unit", func() {
				Context("when force mode", func() {
					It("decreases the value", func() {
						boom := Boom{Force: true}
						err := boom.ScaleInstances(&validManifest, "brain", 0.8)
						Expect(err).NotTo(HaveOccurred())
						Expect(validManifest.Jobs[0].Instances).To(Equal(1))
					})
					It("increases the value", func() {
						boom := Boom{Force: true}
						err := boom.ScaleInstances(&validManifest, "brain", 1.2)
						Expect(err).NotTo(HaveOccurred())
						Expect(validManifest.Jobs[0].Instances).To(Equal(3))
					})
				})
				Context("when force mode isn't used", func() {
					It("increases the value", func() {
						err := boom.ScaleInstances(&validManifest, "brain", 1.2)
						Expect(err).NotTo(HaveOccurred())
						Expect(validManifest.Jobs[0].Instances).To(Equal(2))
					})
				})
			})
		})

		Context("when the job is not found", func() {
			It("returns an error", func() {
				err := boom.SetInstances(&validManifest, "not-existing", 2)
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("job `not-existing` not found"))

			})
		})
	})
})
