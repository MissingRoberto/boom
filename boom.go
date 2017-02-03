package boom

import (
	"errors"
	"fmt"
)

type Boom struct {
}

type Manifest struct {
	Jobs []Job `yaml:"jobs"`
}

type Job struct {
	Name      string
	Instances int
}

func (b *Boom) SetInstances(manifest *Manifest, name string, value int) error {
	for k, job := range manifest.Jobs {
		if name == job.Name {
			manifest.Jobs[k].Instances = value
			return nil
		}
	}
	return errors.New(fmt.Sprintf("job `%v` not found", name))

}
