package registry

import (
	"fmt"
)
type Module interface {
	Identity() string
}

type Registry struct {
	Modules map[string]Module
}

var Default Registry

func init() {
	Default = Registry{
		Modules: make(map[string]Module),
	}
}

func (r *Registry) Register(m Module) {
	fmt.Printf("Registering %s\n", m.Identity())
	r.Modules[m.Identity()] = m
}
