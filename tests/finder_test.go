package finder_test

import (
	"fmt"
	"github.com/romariuse/go/utils/finder"
	"gopkg.in/check.v1"
	"path/filepath"
	"runtime"
	"sync"
	"testing"
)

type Suite struct {
	path    string
	mu      sync.Mutex
	missing []finder.Result
	err     error
}

func Test(t *testing.T) {
	check.Suite(&Suite{})
	check.TestingT(t)
}

func (m *Suite) SetUpSuite(c *check.C) {
	_, file, _, _ := runtime.Caller(0)
	m.path = filepath.Dir(file) + "/data"
}
func (m *Suite) TearDownSuite(c *check.C) {}

func (m *Suite) SetUpTest(c *check.C) {
	fmt.Println("===", c.TestName())
	m.mu.Lock()
}
func (m *Suite) TearDownTest(c *check.C) {
	m.mu.Unlock()
}

func (m *Suite) BenchmarkXXXX(c *check.C) {
	for i := 0; i < c.N; i++ {
		//For benchmark run go.check with -check.b program arguments
	}
}

func (m *Suite) GetShortList() []string {
	var result = make([]string, 0, len(m.missing))
	for _, v := range m.missing {
		result = append(result, v.Short())
	}

	return result
}

func (m *Suite) Test_01_01_0_0(c *check.C) {
	m.missing, m.err = finder.New(finder.PathConfig{
		Path:      m.path + "/src/v1",
		Recursive: false,
	}).Unused(finder.PathConfig{
		Path:      m.path + "/dst/v1",
		Recursive: false,
	})
	c.Assert(m.err, check.IsNil)
	c.Assert(len(m.missing), check.Equals, 0)
}
func (m *Suite) Test_01_01_0_1(c *check.C) {
	m.missing, m.err = finder.New(finder.PathConfig{
		Path:      m.path + "/src/v1",
		Recursive: false,
	}).Unused(finder.PathConfig{
		Path:      m.path + "/dst/v1",
		Recursive: true,
	})
	c.Assert(m.err, check.IsNil)
	c.Assert(len(m.missing), check.Equals, 0)
}
func (m *Suite) Test_01_01_1_0(c *check.C) {
	m.missing, m.err = finder.New(finder.PathConfig{
		Path:      m.path + "/src/v1",
		Recursive: true,
	}).Unused(finder.PathConfig{
		Path:      m.path + "/dst/v1",
		Recursive: false,
	})
	c.Assert(m.err, check.IsNil)
	c.Assert(len(m.missing), check.Equals, 0)
}
func (m *Suite) Test_01_01_1_1(c *check.C) {
	m.missing, m.err = finder.New(finder.PathConfig{
		Path:      m.path + "/src/v1",
		Recursive: true,
	}).Unused(finder.PathConfig{
		Path:      m.path + "/dst/v1",
		Recursive: true,
	})
	c.Assert(m.err, check.IsNil)
	c.Assert(len(m.missing), check.Equals, 0)
}
func (m *Suite) Test_01_02_0_0(c *check.C) {
	m.missing, m.err = finder.New(finder.PathConfig{
		Path:      m.path + "/src/v1",
		Recursive: false,
	}).Unused(finder.PathConfig{
		Path:      m.path + "/dst/v2",
		Recursive: false,
	})
	c.Assert(m.err, check.IsNil)
	c.Assert(m.GetShortList(), check.DeepEquals, []string{
		"v1.VA",
	})
}
func (m *Suite) Test_01_02_0_1(c *check.C) {
	m.missing, m.err = finder.New(finder.PathConfig{
		Path:      m.path + "/src/v1",
		Recursive: false,
	}).Unused(finder.PathConfig{
		Path:      m.path + "/dst/v2",
		Recursive: true,
	})
	c.Assert(m.err, check.IsNil)
	c.Assert(m.GetShortList(), check.DeepEquals, []string{
		"v1.VA",
	})
}
func (m *Suite) Test_01_02_1_0(c *check.C) {
	m.missing, m.err = finder.New(finder.PathConfig{
		Path:      m.path + "/src/v1",
		Recursive: true,
	}).Unused(finder.PathConfig{
		Path:      m.path + "/dst/v2",
		Recursive: false,
	})
	c.Assert(m.err, check.IsNil)
	c.Assert(m.GetShortList(), check.DeepEquals, []string{
		"v1.VA",
		"v11.VA",
		"v11.VB",
	})
}
func (m *Suite) Test_01_02_1_1(c *check.C) {
	m.missing, m.err = finder.New(finder.PathConfig{
		Path:      m.path + "/src/v1",
		Recursive: true,
	}).Unused(finder.PathConfig{
		Path:      m.path + "/dst/v2",
		Recursive: true,
	})
	c.Assert(m.err, check.IsNil)
	c.Assert(m.GetShortList(), check.DeepEquals, []string{
		"v1.VA",
		"v11.VA",
		"v11.VB",
	})
}
func (m *Suite) Test_01_03_0_0(c *check.C) {
	m.missing, m.err = finder.New(finder.PathConfig{
		Path:      m.path + "/src/v1",
		Recursive: false,
	}).Unused(finder.PathConfig{
		Path:      m.path + "/dst/v3",
		Recursive: false,
	})
	c.Assert(m.err, check.IsNil)
	c.Assert(m.GetShortList(), check.DeepEquals, []string{
		"v1.VA",
	})
}
func (m *Suite) Test_01_03_0_1(c *check.C) {
	m.missing, m.err = finder.New(finder.PathConfig{
		Path:      m.path + "/src/v1",
		Recursive: false,
	}).Unused(finder.PathConfig{
		Path:      m.path + "/dst/v3",
		Recursive: true,
	})
	c.Assert(m.err, check.IsNil)
	c.Assert(m.GetShortList(), check.DeepEquals, []string{
		"v1.VA",
	})
}
func (m *Suite) Test_01_03_1_0(c *check.C) {
	m.missing, m.err = finder.New(finder.PathConfig{
		Path:      m.path + "/src/v1",
		Recursive: true,
	}).Unused(finder.PathConfig{
		Path:      m.path + "/dst/v3",
		Recursive: false,
	})
	c.Assert(m.err, check.IsNil)
	c.Assert(m.GetShortList(), check.DeepEquals, []string{
		"v1.VA",
		"v11.VA",
		"v11.VB",
	})
}
func (m *Suite) Test_01_03_1_1(c *check.C) {
	m.missing, m.err = finder.New(finder.PathConfig{
		Path:      m.path + "/src/v1",
		Recursive: true,
	}).Unused(finder.PathConfig{
		Path:      m.path + "/dst/v3",
		Recursive: true,
	})
	c.Assert(m.err, check.IsNil)
	c.Assert(m.GetShortList(), check.DeepEquals, []string{
		"v1.VA",
		"v11.VA",
		"v11.VB",
	})
}
func (m *Suite) Test_01_04_0_0(c *check.C) {
	m.missing, m.err = finder.New(finder.PathConfig{
		Path:      m.path + "/src/v1",
		Recursive: false,
	}).Unused(finder.PathConfig{
		Path:      m.path + "/dst/v4",
		Recursive: false,
	})
	c.Assert(m.err, check.IsNil)
	c.Assert(len(m.missing), check.Equals, 0)
}
func (m *Suite) Test_01_04_0_1(c *check.C) {
	m.missing, m.err = finder.New(finder.PathConfig{
		Path:      m.path + "/src/v1",
		Recursive: false,
	}).Unused(finder.PathConfig{
		Path:      m.path + "/dst/v4",
		Recursive: true,
	})
	c.Assert(m.err, check.IsNil)
	c.Assert(len(m.missing), check.Equals, 0)
}
func (m *Suite) Test_01_04_1_0(c *check.C) {
	m.missing, m.err = finder.New(finder.PathConfig{
		Path:      m.path + "/src/v1",
		Recursive: true,
	}).Unused(finder.PathConfig{
		Path:      m.path + "/dst/v4",
		Recursive: false,
	})
	c.Assert(m.err, check.IsNil)
	c.Assert(m.GetShortList(), check.DeepEquals, []string{
		"v11.VA",
		"v11.VB",
	})
}
func (m *Suite) Test_01_04_1_1(c *check.C) {
	m.missing, m.err = finder.New(finder.PathConfig{
		Path:      m.path + "/src/v1",
		Recursive: true,
	}).Unused(finder.PathConfig{
		Path:      m.path + "/dst/v4",
		Recursive: true,
	})
	c.Assert(m.err, check.IsNil)
	c.Assert(m.GetShortList(), check.DeepEquals, []string{
		"v11.VA",
		"v11.VB",
	})
}
func (m *Suite) Test_01_05_0_0(c *check.C) {
	m.missing, m.err = finder.New(finder.PathConfig{
		Path:      m.path + "/src/v1",
		Recursive: false,
	}).Unused(finder.PathConfig{
		Path:      m.path + "/dst/v5",
		Recursive: false,
	})
	c.Assert(m.err, check.IsNil)
	c.Assert(len(m.missing), check.Equals, 0)
}
func (m *Suite) Test_01_05_0_1(c *check.C) {
	m.missing, m.err = finder.New(finder.PathConfig{
		Path:      m.path + "/src/v1",
		Recursive: false,
	}).Unused(finder.PathConfig{
		Path:      m.path + "/dst/v5",
		Recursive: true,
	})
	c.Assert(m.err, check.IsNil)
	c.Assert(len(m.missing), check.Equals, 0)
}
func (m *Suite) Test_01_05_1_0(c *check.C) {
	m.missing, m.err = finder.New(finder.PathConfig{
		Path:      m.path + "/src/v1",
		Recursive: true,
	}).Unused(finder.PathConfig{
		Path:      m.path + "/dst/v5",
		Recursive: false,
	})
	c.Assert(m.err, check.IsNil)
	c.Assert(len(m.missing), check.Equals, 0)
}
func (m *Suite) Test_01_05_1_1(c *check.C) {
	m.missing, m.err = finder.New(finder.PathConfig{
		Path:      m.path + "/src/v1",
		Recursive: true,
	}).Unused(finder.PathConfig{
		Path:      m.path + "/dst/v5",
		Recursive: true,
	})
	c.Assert(m.err, check.IsNil)
	c.Assert(len(m.missing), check.Equals, 0)
}

func (m *Suite) Test_XX_06_1_1(c *check.C) {
	m.missing, m.err = finder.New(finder.PathConfig{
		Path:      m.path + "/dst/v6/internal",
		Recursive: true,
	}).Unused(finder.PathConfig{
		Path:      m.path + "/dst/v6",
		Recursive: true,
	})
	c.Assert(m.err, check.IsNil)
	c.Assert(m.GetShortList(), check.DeepEquals, []string{
		"v11.VA",
		"v11.VB",
	})
}
