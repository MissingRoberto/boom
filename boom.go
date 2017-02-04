package boom

import (
	"errors"
	"fmt"
)

type Boom struct {
	Force bool
}

type Manifest struct {
	Jobs []Job `yaml:"jobs"`
}

type Job struct {
	Name      string
	Instances int
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

	manifest.Jobs[index].Instances = newValue
	return nil
}

func (b *Boom) SetInstances(manifest *Manifest, name string, value int) error {
	index, err := findJob(manifest, name)
	if err != nil {
		return err
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
