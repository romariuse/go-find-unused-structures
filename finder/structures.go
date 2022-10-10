package finder

import (
	"fmt"
)

type Result struct {
	PkgFull  string
	PkgShort string
	Struct   string
}

func (m *Result) String() string {
	return fmt.Sprintf("%s.%s", m.PkgFull, m.Struct)
}
func (m *Result) Short() string {
	return fmt.Sprintf("%s.%s", m.PkgShort, m.Struct)
}

type PathConfig struct {
	Path      string
	Recursive bool
}
