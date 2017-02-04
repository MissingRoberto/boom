package boom

import (
	"errors"
	"fmt"
)

type Boom struct {
	Force bool
}

type Manifest struct {
	Jobs          []Job          `yaml:"jobs"`
	ResourcePools []ResourcePool `yaml:"resource_pools"`
}

type Job struct {
	Name         string `yaml:"name"`
	Instances    int    `yaml:"instances"`
	ResourcePool string `yaml:"resource_pool"`
}

type ResourcePool struct {
	Name string `yaml:"name"`
	Size int    `yaml:"size"`
}

func (b *Boom) ScaleInstances(manifest *Manifest, name string, factor float32) error {

	if factor == 0 {
		return errors.New("factor 0 is not permitted")
	}
	index, err := findJob(manifest, name)
	if err != nil {
		return err
	}
	oldValue := manifest.Jobs[index].Instances
	newValue := int(float32(oldValue) * factor)
	if b.Force && oldValue == newValue {
		if factor > 0 {
			newValue++
		} else {
			newValue--
		}
	}

	indexResourcePool, err := findResourcePool(manifest, manifest.Jobs[index].ResourcePool)
	if err == nil {
		manifest.ResourcePools[indexResourcePool].Size = newValue
	}
	manifest.Jobs[index].Instances = newValue
	return nil
}

func (b *Boom) SetInstances(manifest *Manifest, name string, value int) error {
	index, err := findJob(manifest, name)
	if err != nil {
		return err
	}

	indexResourcePool, err := findResourcePool(manifest, manifest.Jobs[index].ResourcePool)
	if err == nil {
		manifest.ResourcePools[indexResourcePool].Size = value
	}

	manifest.Jobs[index].Instances = value
	return nil
}

func findJob(manifest *Manifest, name string) (int, error) {

	for index, job := range manifest.Jobs {
		if name == job.Name {
			return index, nil
		}
	}
	return -1, errors.New(fmt.Sprintf("job `%v` not found", name))

}

func findResourcePool(manifest *Manifest, name string) (int, error) {

	for index, resourcePool := range manifest.ResourcePools {
		if name == resourcePool.Name {
			return index, nil
		}
	}
	return -1, errors.New(fmt.Sprintf("resource_pool `%v` not found", name))

}
