package boom

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/smallfish/simpleyaml"
)

var _ = Describe("Boom", func() {
	var (
		boom *Boom
	)

	Context("Mask", func() {
		BeforeEach(func() {
			boom = New(completeManifestPath, false)
		})
		It("clears all but the given list in root", func() {
			err := boom.Mask("jobs", "instances")
			Expect(err).NotTo(HaveOccurred())
			result, err := simpleyaml.NewYaml([]byte(boom.String()))
			Expect(err).NotTo(HaveOccurred())
			manifest, err := result.Map()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(manifest)).To(Equal(1))
			Expect(manifest["jobs"]).ToNot(BeNil())
		})
		Context("when the list doesn't exist", func() {
			It("returns an error", func() {
				err := boom.Mask("not-existing", "instances")
				Expect(err).To(MatchError("list `not-existing` not found"))
			})
		})
		Context("when  the list specified isn't a list", func() {
			It("returns an error", func() {
				err := boom.Mask("properties", "instances")
				Expect(err).To(MatchError("key `properties` is not a list"))
			})
		})
		Context("when the key doesn't exist", func() {
			It("does return empty values", func() {
				err := boom.Mask("jobs", "not-existing-key")
				Expect(err).NotTo(HaveOccurred())
				result, err := simpleyaml.NewYaml([]byte(boom.String()))
				Expect(err).NotTo(HaveOccurred())
				job := result.Get("jobs").GetIndex(0)
				Expect(job).NotTo(BeNil())
				jobMap, err := job.Map()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(jobMap)).To(Equal(2))
				Expect(job.Get("not-existing-key")).NotTo(BeNil())
			})
		})
		Context("when only the list is specified", func() {
			It("clears all but name", func() {
				err := boom.Mask("jobs", "")
				Expect(err).NotTo(HaveOccurred())
				result, err := simpleyaml.NewYaml([]byte(boom.String()))
				Expect(err).NotTo(HaveOccurred())
				job := result.Get("jobs").GetIndex(0)
				Expect(job).NotTo(BeNil())
				jobMap, err := job.Map()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(jobMap)).To(Equal(1))
			})
		})
		Context("Instances", func() {
			It("clears all but name and instances", func() {
				err := boom.Mask("jobs", "instances")
				Expect(err).NotTo(HaveOccurred())
				result, err := simpleyaml.NewYaml([]byte(boom.String()))
				Expect(err).NotTo(HaveOccurred())
				job := result.Get("jobs").GetIndex(0)
				Expect(job).NotTo(BeNil())
				jobMap, err := job.Map()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(jobMap)).To(Equal(2))
			})
		})
		Context("Instances", func() {
			It("clears all but name and instances", func() {
				err := boom.Mask("jobs", "instances")
				Expect(err).NotTo(HaveOccurred())
				result, err := simpleyaml.NewYaml([]byte(boom.String()))
				Expect(err).NotTo(HaveOccurred())
				job := result.Get("jobs").GetIndex(0)
				Expect(job).NotTo(BeNil())
				jobMap, err := job.Map()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(jobMap)).To(Equal(2))
			})
		})

		Context("Versions", func() {
			It("clears all but name and version", func() {
				err := boom.Mask("releases", "version")
				Expect(err).NotTo(HaveOccurred())
				result, err := simpleyaml.NewYaml([]byte(boom.String()))
				Expect(err).NotTo(HaveOccurred())
				job := result.Get("releases").GetIndex(0)
				Expect(job).NotTo(BeNil())
				jobMap, err := job.Map()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(jobMap)).To(Equal(2))
			})
		})
	})
})
