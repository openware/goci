package versions

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v3"
)

// Versions store and manage versions.yaml
type Versions struct {
	data     map[string]interface{}
	filename string
}

// Load values.yaml
func Load(filename string) (*Versions, error) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	v := Versions{
		data:     make(map[string]interface{}),
		filename: filename,
	}
	err = yaml.Unmarshal(dat, v.data)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

// SetTag is a helper to set a tag for a given component
func (v *Versions) SetTag(component, value string) {
	v.data[component].(map[string]interface{})["image"].(map[string]interface{})["tag"] = value
}

// Save dump the values in yaml to the file
func (v *Versions) Save() error {
	var b bytes.Buffer
	yamlEncoder := yaml.NewEncoder(&b)
	yamlEncoder.SetIndent(2)
	yamlEncoder.Encode(&v.data)

	return ioutil.WriteFile(v.filename, b.Bytes(), 0644)
}

// Get any value from the structure
func (v *Versions) Get(keys ...string) (interface{}, error) {
	var d interface{} = v.data
	for _, k := range keys {
		e, ok := d.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("Invalid path %s in values", strings.Join(keys, "."))
		}
		d = e[k]
	}
	return d, nil
}
