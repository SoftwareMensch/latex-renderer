package di

import "fmt"

type container struct {
	parameters map[string]string
	services   map[string]interface{}
}

var containerInst *container

// GetContainer return new service container
func GetContainer() ContainerInterface {
	if containerInst == nil {
		containerInst = &container{}
	}

	return containerInst
}

// AddParameter add a parameter
func (c *container) AddParameter(name, param string) {
	if c.parameters == nil {
		c.parameters = make(map[string]string)
	}

	c.parameters[name] = param
}

// GetParameter a parameter
func (c *container) GetParameter(name, param string) (string, error) {
	if c.parameters == nil {
		return "", fmt.Errorf("no parameters defined")
	}

	if param, ok := c.parameters[name]; ok {
		return param, nil
	}

	return "", fmt.Errorf("parameter \"%s\" not found", name)
}

// AddService add a service
func (c *container) AddService(name string, service interface{}) {
	if c.services == nil {
		c.services = make(map[string]interface{})
	}

	c.services[name] = service
}

// GetService get service
func (c *container) GetService(name string) (interface{}, error) {
	if c.services == nil {
		return nil, fmt.Errorf("no services defined")
	}

	if service, ok := c.services[name]; ok {
		return service, nil
	}

	return nil, fmt.Errorf("service \"%s\" not found", name)
}
