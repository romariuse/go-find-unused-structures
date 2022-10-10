package errorEnum

import "fmt"

type FailedLoadPackageUnexpected struct {
	Path    string
	Pattern string
	Err     error
}

func (m FailedLoadPackageUnexpected) Error() string {
	return fmt.Sprintf("failed load packages in %q, with err = %s", m.Path, m.Err)
}

type FailedLoadPackageStructuresEmpty struct {
	Path string
}

func (m FailedLoadPackageStructuresEmpty) Error() string {
	return fmt.Sprintf("not found structures in %q", m.Path)
}

type FailedLoadPackageSourceEmpty struct {
	Path    string
	Pattern string
}

func (m FailedLoadPackageSourceEmpty) Error() string {
	return fmt.Sprintf("not found sources in %q", m.Path)
}

type FailedGetAbsolutePath struct {
	Path string
}

func (m FailedGetAbsolutePath) Error() string {
	return fmt.Sprintf("can't get absolute path for %q", m.Path)
}
