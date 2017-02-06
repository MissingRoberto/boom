package boom

import (
	"errors"
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

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
