package module1

import (
	"fmt"
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
	globalRegistry.Register(m1)
}
