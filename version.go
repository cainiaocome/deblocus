package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	app_name    = "deblocus"
	project_url = "https://github.com/spance/delocus"
	ver_major   = uint8(0)
	ver_minor   = uint8(9)
	ver_build   = 10*uint16(223) + uint16(0)
)

var build_flag string // -ldflags "-X main.build_flag -beta"

var version uint32

func init() {
	var ver uint32
	ver |= uint32(ver_major) << 24
	ver |= uint32(ver_minor) << 16
	ver |= uint32(ver_build)
	version = ver
}

func versionString() string {
	return fmt.Sprintf("%s version: %d.%d.%04d%s\n", app_name, ver_major, ver_minor, ver_build, build_flag)
}

type CArg struct {
	literal string
	usage   string
	example string
}

func showUsage() {
	fmt.Printf("Usage: %s [-OPTION=VALUE]...\n", filepath.Base(os.Args[0]))
	fmt.Printf("%s project: <%s>\n", app_name, project_url)
	fmt.Println(versionString())

	var group = map[string][]*CArg{}
	var common = "Common"
	var reBool, _ = regexp.Compile("(?i)true|false")
	flag.VisitAll(func(f *flag.Flag) {
		var literal string
		if len(f.DefValue) > 0 && !reBool.MatchString(f.DefValue) {
			literal = fmt.Sprintf("-%s=%s", f.Name, f.DefValue)
		} else {
			literal = "-" + f.Name
		}
		array := strings.SplitN(f.Usage, ";;", 2)
		if len(array) != 2 {
			array = []string{common, f.Usage}
		}
		cArg := &CArg{literal: literal, usage: array[1]}
		if strings.Index(array[1], "//") >= 0 {
			ue := strings.SplitN(array[1], "//", 2)
			cArg.usage, cArg.example = ue[0], ue[1]
		}
		group[array[0]] = append(group[array[0]], cArg)
	})
	for k, a := range group {
		fmt.Printf("%s options:\n", k)
		for _, i := range a {
			fmt.Printf("  %-12s %s\n", i.literal, i.usage)
			if i.example != "" {
				fmt.Printf("  %-12s %s\n", " ", i.example)
			}
		}
	}
}
