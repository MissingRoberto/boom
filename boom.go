package boom

import (
	"errors"
	"fmt"
	"log"
	"os"

	yaml "github.com/geofffranks/yaml"
)

type Boom struct {
	Force    bool
	Manifest map[string]interface{}
}

func New(manifestPath string, force bool) *Boom {

	manifest, err := loadYML(manifestPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return &Boom{Manifest: manifest, Force: force}
}

func (b *Boom) ScaleInstances(name string, factor float64) error {
	if factor == 0 {
		return errors.New("factor 0 is not permitted")
	}

	if b.Manifest["jobs"] == nil {
		return nil
	}
	jobs := b.Manifest["jobs"].([]interface{})
	job, _, err := findByName(jobs, name)
	if err != nil {
		return err
	}
	oldValue := job["instances"].(int)
	newValue := int(float64(oldValue) * factor)
	if b.Force && newValue == oldValue {
		if factor > 1 {
			newValue++
		} else {
			newValue--
		}
	}
	return b.SetInstances(name, newValue)
}

func (b *Boom) SetInstances(name string, value int) error {
	if b.Manifest["jobs"] == nil {
		return nil
	}
	jobs := b.Manifest["jobs"].([]interface{})
	job, index, err := findByName(jobs, name)
	if err != nil {
		return err
	}

	oldValue := job["instances"].(int)
	job["instances"] = value
	jobs[index] = job
	b.Manifest["jobs"] = jobs

	if job["resource_pool"] != nil && b.Manifest["resource_pools"] != nil {
		resourcePools := b.Manifest["resource_pools"].([]interface{})
		pool, indexResourcePool, err := findByName(resourcePools, job["resource_pool"].(string))

		if err == nil {
			pool["size"] = value - oldValue + pool["size"].(int)
			resourcePools[indexResourcePool] = pool
			b.Manifest["resource_pools"] = resourcePools
		}
	}
	return nil
}

func (b *Boom) String() string {
	d, err := yaml.Marshal(b.Manifest)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return fmt.Sprintf("---\n%s\n\n", string(d))
}

func (b *Boom) Print() {
	fmt.Printf("%s", b.String())
}
