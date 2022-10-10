package main

import (
	"flag"
	"fmt"
	"github.com/romariuse/go/utils/finder"
	"os"
	"strings"
)

const (
	keySrc = iota
	keyDst
	_keyEND
)

const (
	recursiveSuffix = "/..."
)

func main() {
	var cfg = make([]finder.PathConfig, _keyEND)

	flag.StringVar(&cfg[keySrc].Path, "src", "", "Path for structures declaration")
	flag.StringVar(&cfg[keyDst].Path, "dst", "", "Path for sources with used structures")
	flag.Usage = func() {
		w := flag.CommandLine.Output()
		_, _ = fmt.Fprintln(w, "For recursive search use (...)-suffix")
		_, _ = fmt.Fprintf(w, "Example: %s -src ./errorEnum' -dst ./...\nParams:\n", os.Args[0])
		flag.PrintDefaults()

	}
	flag.Parse()

	for k, v := range cfg {
		if v.Path == "" {
			flag.Usage()
			os.Exit(1)
		}
		if strings.HasSuffix(v.Path, recursiveSuffix) {
			cfg[k].Recursive = true
			cfg[k].Path = strings.TrimSuffix(v.Path, recursiveSuffix)
		}
	}

	missing, err := finder.New(cfg[keySrc]).Unused(cfg[keyDst])

	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	missingCount := len(missing)
	if missingCount == 0 {
		os.Exit(0)
	}

	var result = make([]string, 0, missingCount)
	for _, r := range missing {
		result = append(result, r.String())
	}

	fmt.Println(strings.Join(result, "\n"))
}
