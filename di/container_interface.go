package di

// ContainerInterface service container interface
type ContainerInterface interface {
	// AddParameter add a parameter
	AddParameter(name, param string)

	// GetParameter a parameter
	GetParameter(name, param string) (string, error)

	// AddService add a service
	AddService(name string, service interface{})

	// GetService get service
	GetService(name string) (interface{}, error)
}
