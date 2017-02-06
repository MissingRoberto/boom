package main

import (
	"github.com/geofffranks/simpleyaml"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Boom", func() {
	var (
		boom *Boom
	)

	Context("ScaleInstances", func() {
		BeforeEach(func() {
			boom = New(completeManifestPath)
		})
		Context("when the job is found", func() {
			Context("when the value is not round", func() {
				It("don't update the value", func() {
					err := boom.ScaleInstances("cell", 1)
					Expect(err).NotTo(HaveOccurred())
					result, err := simpleyaml.NewYaml([]byte(boom.String()))
					Expect(err).NotTo(HaveOccurred())
					Expect(result.Get("jobs").GetIndex(0).Get("instances").Int()).To(Equal(20))
				})
			})
			It("decreases the value", func() {
				err := boom.ScaleInstances("cell", 0.5)
				Expect(err).NotTo(HaveOccurred())
				result, err := simpleyaml.NewYaml([]byte(boom.String()))
				Expect(err).NotTo(HaveOccurred())
				Expect(result.Get("jobs").GetIndex(0).Get("instances").Int()).To(Equal(10))
			})
			It("increases the value", func() {
				err := boom.ScaleInstances("cell", 2)
				Expect(err).NotTo(HaveOccurred())
				result, err := simpleyaml.NewYaml([]byte(boom.String()))
				Expect(err).NotTo(HaveOccurred())
				Expect(result.Get("jobs").GetIndex(0).Get("instances").Int()).To(Equal(40))
			})
			Context("when the value is not round", func() {
				It("updates the value", func() {
					err := boom.ScaleInstances("cell", 1.5)
					Expect(err).NotTo(HaveOccurred())
					result, err := simpleyaml.NewYaml([]byte(boom.String()))
					Expect(err).NotTo(HaveOccurred())
					Expect(result.Get("jobs").GetIndex(0).Get("instances").Int()).To(Equal(30))
				})
			})
			Context("when the factor is 0", func() {
				It("returns an error", func() {
					err := boom.ScaleInstances("cell", 0)
					Expect(err).To(MatchError("factor 0 is not permitted"))
				})
			})

			Context("when the factor is too low to modify a unit", func() {
				Context("when force mode", func() {
					It("decreases the value", func() {
						err := boom.ScaleInstances("brain", 0.8)
						Expect(err).NotTo(HaveOccurred())
						result, err := simpleyaml.NewYaml([]byte(boom.String()))
						Expect(err).NotTo(HaveOccurred())
						Expect(result.Get("jobs").GetIndex(1).Get("instances").Int()).To(Equal(1))
					})
					It("increases the value", func() {
						boom.Force = true
						err := boom.ScaleInstances("brain", 1.2)
						Expect(err).NotTo(HaveOccurred())
						result, err := simpleyaml.NewYaml([]byte(boom.String()))
						Expect(err).NotTo(HaveOccurred())
						Expect(result.Get("jobs").GetIndex(1).Get("instances").Int()).To(Equal(3))
					})
				})
				Context("when force mode isn't used", func() {
					It("don't increase the value", func() {
						err := boom.ScaleInstances("brain", 1.2)
						Expect(err).NotTo(HaveOccurred())
						result, err := simpleyaml.NewYaml([]byte(boom.String()))
						Expect(err).NotTo(HaveOccurred())
						Expect(result.Get("resource_pools").GetIndex(1).Get("size").Int()).To(Equal(5))
					})
				})
			})
			Context("when the resource_pool is found", func() {
				Context("when a dedicated resource pool is used", func() {
					It("modifies the resource pool size", func() {
						err := boom.ScaleInstances("cell", 0.5)
						Expect(err).NotTo(HaveOccurred())
						result, err := simpleyaml.NewYaml([]byte(boom.String()))
						Expect(err).NotTo(HaveOccurred())
						Expect(result.Get("resource_pools").GetIndex(0).Get("size").Int()).To(Equal(10))
					})
				})
				Context("when a shared resource pool is used", func() {
					It("modifies the resource pool size", func() {
						err := boom.ScaleInstances("brain", 2)
						Expect(err).NotTo(HaveOccurred())
						result, err := simpleyaml.NewYaml([]byte(boom.String()))
						Expect(err).NotTo(HaveOccurred())
						Expect(result.Get("resource_pools").GetIndex(1).Get("size").Int()).To(Equal(7))
					})
				})
			})
		})

		Context("when the job is not found", func() {
			It("returns an error", func() {
				err := boom.SetInstances("not-existing", 2)
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("element `not-existing` not found"))

			})
		})
	})
})
