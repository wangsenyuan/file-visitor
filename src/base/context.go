package base

import (
	"fmt"
	"flag"
)

type Context interface {
	GetEnv() EnvInterface
	GetSrc() SrcInterface
	GetDest() DestInterface
}

type EnvInterface interface {
	ShouldIncludeFile(folder string, file string) bool
}

type SrcInterface interface {
	GetName() string
}

type DestInterface interface {
	GetName() string
	GetType() string
	GetOldNS() string
	GetNewNS() string
}

type Help struct {
	Value string
}

func (help *Help) Set(s string) error {
	help.Value = s
	return nil
}

func (help *Help) String() string {
	return fmt.Sprintf("-help config-example")
}

func HelpFlag() *Help {
	help := &Help{}
	flag.CommandLine.Var(help, "help", "show usage")
	return help
}
