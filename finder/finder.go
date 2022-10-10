package finder

import (
	"fmt"
	"github.com/romariuse/go/utils/errorEnum"
	"golang.org/x/tools/go/packages"
	"path/filepath"
	"sort"
)

const (
	currentPath   = "."
	recursivePath = "./..."
)

func New(src PathConfig) *module {
	return &module{
		src: src,
	}
}

type module struct {
	src               PathConfig
	thatAreLookingFor map[string]Result
	used              map[string]Result
}

func (m *module) initThatAreLookingFor() error {
	if m.thatAreLookingFor != nil {
		return nil
	}

	abs, err := filepath.Abs(m.src.Path)
	if err != nil {
		return errorEnum.FailedGetAbsolutePath{
			Path: m.src.Path,
		}
	}

	m.src.Path = abs

	pkgList, err := findStructuresPackages(m.src)
	if err != nil {
		return err
	}

	m.thatAreLookingFor = make(map[string]Result)

	for _, search := range pkgList {
		for _, name := range search.Types.Scope().Names() {
			t := search.Types.Scope().Lookup(name)
			if t.Type().String() != fmt.Sprintf("%s.%s", search.ID, name) {
				continue
			}

			r := Result{
				PkgFull:  search.ID,
				PkgShort: t.Pkg().Name(),
				Struct:   name,
			}

			m.thatAreLookingFor[r.String()] = r
		}
	}

	return nil
}

func (m *module) findUsed(dst PathConfig) error {
	abs, err := filepath.Abs(dst.Path)
	if err != nil {
		return errorEnum.FailedGetAbsolutePath{
			Path: dst.Path,
		}
	}

	dst.Path = abs

	pattern := currentPath
	if dst.Recursive {
		pattern = recursivePath
	}

	pkgList, err := packages.Load(
		&packages.Config{
			Dir:  dst.Path,
			Mode: packages.NeedTypes | packages.NeedTypesInfo | packages.NeedExportFile,
		}, pattern)
	if err != nil {
		return errorEnum.FailedLoadPackageUnexpected{
			Path:    dst.Path,
			Pattern: pattern,
			Err:     err,
		}
	}

	var pkgListWithoutExclude []*packages.Package
	for _, pkg := range pkgList {
		//TODO Compare with m.src
		pkgListWithoutExclude = append(pkgListWithoutExclude, pkg)
	}

	if len(pkgList) == 0 {
		return errorEnum.FailedLoadPackageSourceEmpty{
			Path:    dst.Path,
			Pattern: pattern,
		}
	}

	m.used = make(map[string]Result)

	for k, search := range m.thatAreLookingFor {
		id := search.PkgFull

		for _, pkg := range pkgList {
			if len(pkg.Errors) != 0 {
				continue
			}

			for _, obj := range pkg.TypesInfo.Uses {
				if obj.Pkg() == nil || obj.Parent() == nil || obj.Pkg().Path() != id {
					continue
				}

				m.used[k] = search
			}
		}
	}

	return nil
}

func (m *module) ThatAreLookingFor() ([]Result, error) {
	if err := m.initThatAreLookingFor(); err != nil {
		return nil, err
	}

	return sortedSlice(m.thatAreLookingFor), nil
}

func (m *module) Used(dst PathConfig) ([]Result, error) {
	return m.compare(true, dst)
}

func (m *module) Unused(dst PathConfig) ([]Result, error) {
	return m.compare(false, dst)
}

func (m *module) compare(compare bool, dst PathConfig) ([]Result, error) {
	if err := m.initThatAreLookingFor(); err != nil {
		return nil, err
	}

	if err := m.findUsed(dst); err != nil {
		return nil, err
	}

	var diff = make(map[string]Result)

	for k, v := range m.thatAreLookingFor {
		if _, ok := m.used[k]; ok == compare {
			diff[k] = v
		}
	}

	return sortedSlice(diff), nil
}

func sortedSlice(diff map[string]Result) []Result {
	diffLen := len(diff)

	var diffKeys = make([]string, 0, diffLen)
	for k := range diff {
		diffKeys = append(diffKeys, k)
	}

	sort.Strings(diffKeys)
	result := make([]Result, 0, diffLen)

	for _, k := range diffKeys {
		result = append(result, diff[k])
	}

	return result
}

func findStructuresPackages(cfg PathConfig) ([]*packages.Package, error) {
	pattern := currentPath
	if cfg.Recursive {
		pattern = recursivePath
	}

	pkgList, err := packages.Load(
		&packages.Config{
			Dir:  cfg.Path,
			Mode: packages.NeedTypes,
		}, pattern)
	if err != nil {
		return nil, errorEnum.FailedLoadPackageUnexpected{
			Path: cfg.Path,
			Err:  err,
		}
	}

	var result = make([]*packages.Package, 0, len(pkgList))
	for _, pkg := range pkgList {
		if len(pkg.Errors) != 0 {
			continue
		}

		result = append(result, pkg)
	}

	if len(result) == 0 {
		return nil, errorEnum.FailedLoadPackageStructuresEmpty{
			Path: cfg.Path,
		}
	}

	return result, nil
}
