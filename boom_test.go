package boom_test

import (
	. "github.com/jszroberto/boom"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Boom", func() {
	var (
		validMap map[string]interface{}
		boom     Boom
	)
	BeforeEach(func() {
		validMap := Manifest{
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
				result, err := boom.SetInstances("cell", 2)
				Expect(err).NotTo(HaveOccurred())
				Expect(result[0].Instances).To(Equal(2))
			})

		})
	})
})
