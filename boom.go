package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Boom struct {
	Force    bool
	Manifest map[string]interface{}
}

func New(manifestPath string) *Boom {

	manifest, err := loadYML(manifestPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	return &Boom{Manifest: manifest}
}

func (b *Boom) ScaleInstances(name string, factor float32) error {
	if factor == 0 {
		return errors.New("factor 0 is not permitted")
	}
	jobs := b.Manifest["jobs"].([]interface{})
	job, index, err := findByName(jobs, name)
	if err != nil {
		return err
	}
	oldValue := job["instances"].(int)
	variation := int(float32(oldValue)*factor) - oldValue
	if b.Force && variation == 0 {
		if factor > 1 {
			variation++
		} else {
			variation--
		}
	}
	job["instances"] = variation + oldValue
	jobs[index] = job
	b.Manifest["jobs"] = jobs

	resourcePools := b.Manifest["resource_pools"].([]interface{})

	pool, indexResourcePool, err := findByName(resourcePools, job["resource_pool"].(string))
	if err == nil {
		pool["size"] = variation + pool["size"].(int)
		resourcePools[indexResourcePool] = pool
		b.Manifest["resource_pools"] = resourcePools
	}

	return nil
}

func (b *Boom) SetInstances(name string, value int) error {
	jobs := b.Manifest["jobs"].([]interface{})
	job, index, err := findByName(jobs, name)
	if err != nil {
		return err
	}

	oldValue := job["instances"].(int)
	job["instances"] = value
	jobs[index] = job
	b.Manifest["jobs"] = jobs

	if b.Manifest["resource_pools"] != nil {
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

func findByName(list []interface{}, name string) (map[string]interface{}, int, error) {
	for index, value := range list {
		element := value.(map[string]interface{})
		if name == element["name"] {
			return element, index, nil
		}
	}
	return nil, -1, errors.New(fmt.Sprintf("element `%v` not found", name))
}
func loadYML(path string) (map[string]interface{}, error) {
	var manifest interface{}
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, &manifest)
	if err != nil {
		return nil, err
	}

	return convert(manifest).(map[string]interface{}), nil

}
func convert(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = convert(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = convert(v)
		}
	}
	return i
}
