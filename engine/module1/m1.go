package module1

import (
	"fmt"

	"github.com/szaydel/goregistry1/engine/registry"
)
type Module1 struct {
	identity string
}

func (m *Module1) Identity() string {
	return m.identity
}

func init() {
	fmt.Printf("XXX\n")
	var m1 = &Module1{
		identity: "Module_1",
	}
	registry.Default.Register(m1)
}
