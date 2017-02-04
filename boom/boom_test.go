package boom_test

import (
	. "github.com/jszroberto/boom/boom"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Boom", func() {
	var (
		completeManifest             Manifest
		manifestWithOutResourcePools Manifest
		boom                         Boom
	)

	BeforeEach(func() {

		jobs := []Job{Job{
			Name:         "brain",
			ResourcePool: "shared-pool",
			Instances:    2,
		},
			Job{
				Name:         "cell",
				Instances:    20,
				ResourcePool: "dedicated-pool",
			},
		}
		resourcePools := []ResourcePool{
			ResourcePool{
				Name: "dedicated-pool",
				Size: 20,
			},
			ResourcePool{
				Name: "shared-pool",
				Size: 5,
			},
		}
		manifestWithOutResourcePools = Manifest{Jobs: jobs}
		completeManifest = Manifest{Jobs: jobs, ResourcePools: resourcePools}

	})

	Context("SetInstances", func() {
		Context("when the job is found", func() {
			Context("when resource_pools does not exist", func() {
				It("updates the value in the job", func() {
					err := boom.SetInstances(&completeManifest, "cell", 2)
					Expect(err).NotTo(HaveOccurred())
					Expect(completeManifest.Jobs[1].Instances).To(Equal(2))
				})
			})
			Context("when resource_pools exist", func() {
				It("updates the value in the job", func() {
					err := boom.SetInstances(&completeManifest, "cell", 2)
					Expect(err).NotTo(HaveOccurred())
					Expect(completeManifest.Jobs[1].Instances).To(Equal(2))
				})
				Context("when the resource pool is found", func() {
					Context("when the given value is the same", func() {
						It("does not modify resource pool", func() {
							previousSize := completeManifest.ResourcePools[0].Size
							err := boom.SetInstances(&completeManifest, "cell", 20)
							Expect(err).NotTo(HaveOccurred())
							Expect(completeManifest.Jobs[1].Instances).To(Equal(20))
							Expect(completeManifest.ResourcePools[0].Size).To(Equal(previousSize))
						})
					})
					Context("when the given value is lower", func() {
						It("decreases the number of instances in resource pool", func() {

							err := boom.SetInstances(&completeManifest, "cell", 15)
							Expect(err).NotTo(HaveOccurred())
							Expect(completeManifest.Jobs[1].Instances).To(Equal(15))
							Expect(completeManifest.ResourcePools[0].Size).To(Equal(15))
						})
					})

					Context("when the given value is greater", func() {
						It("increases the number of instances in resource pool", func() {
							err := boom.SetInstances(&completeManifest, "cell", 25)
							Expect(err).NotTo(HaveOccurred())
							Expect(completeManifest.Jobs[1].Instances).To(Equal(25))
							Expect(completeManifest.ResourcePools[0].Size).To(Equal(25))
						})
					})
				})
			})
		})
		Context("when the job is not found", func() {
			It("returns an error", func() {
				err := boom.SetInstances(&completeManifest, "not-existing", 2)
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("job `not-existing` not found"))

			})
		})
	})

	Context("ScaleInstances", func() {
		Context("when the job is found", func() {
			Context("when the value is not round", func() {
				It("don't update the value", func() {
					err := boom.ScaleInstances(&completeManifest, "cell", 1)
					Expect(err).NotTo(HaveOccurred())
					Expect(completeManifest.Jobs[1].Instances).To(Equal(20))
				})
			})
			It("decreases the value", func() {
				err := boom.ScaleInstances(&completeManifest, "cell", 0.5)
				Expect(err).NotTo(HaveOccurred())
				Expect(completeManifest.Jobs[1].Instances).To(Equal(10))
			})
			It("increases the value", func() {
				err := boom.ScaleInstances(&completeManifest, "cell", 2)
				Expect(err).NotTo(HaveOccurred())
				Expect(completeManifest.Jobs[1].Instances).To(Equal(40))
			})
			Context("when the value is not round", func() {
				It("updates the value", func() {
					err := boom.ScaleInstances(&completeManifest, "cell", 1.5)
					Expect(err).NotTo(HaveOccurred())
					Expect(completeManifest.Jobs[1].Instances).To(Equal(30))
				})
			})
			Context("when the factor is 0", func() {
				It("returns an error", func() {
					err := boom.ScaleInstances(&completeManifest, "cell", 0)
					Expect(err).To(MatchError("factor 0 is not permitted"))
				})
			})
			Context("when the factor is too low to modify a unit", func() {
				Context("when force mode", func() {
					It("decreases the value", func() {
						boom := Boom{Force: true}
						err := boom.ScaleInstances(&completeManifest, "brain", 0.8)
						Expect(err).NotTo(HaveOccurred())
						Expect(completeManifest.Jobs[0].Instances).To(Equal(1))
					})
					It("increases the value", func() {
						boom := Boom{Force: true}
						err := boom.ScaleInstances(&completeManifest, "brain", 1.2)
						Expect(err).NotTo(HaveOccurred())
						Expect(completeManifest.Jobs[0].Instances).To(Equal(3))
					})
				})
				Context("when force mode isn't used", func() {
					It("increases the value", func() {
						err := boom.ScaleInstances(&completeManifest, "brain", 1.2)
						Expect(err).NotTo(HaveOccurred())
						Expect(completeManifest.Jobs[0].Instances).To(Equal(2))
					})
				})
			})
		})

		Context("when the job is not found", func() {
			It("returns an error", func() {
				err := boom.SetInstances(&completeManifest, "not-existing", 2)
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("job `not-existing` not found"))

			})
		})
	})
})
