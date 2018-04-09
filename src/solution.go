package main

import (
	"./base"
	"flag"
	"./handlers/std"
	"./handlers/file"
	"./handlers/dir"
	"./visitor"
	"os"
	"bufio"
	"fmt"
	"./base/yaml"
	"./lineprocessors"
)

var configFlag = yaml.ConfigOptionFlag()
var helpFlag = base.HelpFlag()

func main() {
	flag.Parse()

	if help(helpFlag) {
		return
	}

	ctx, err := yaml.NewContext(configFlag)

	if err != nil {
		panic(err)
	}

	err = safeVisit(ctx)
	if err != nil {
		panic(err)
	}
	//fmt.Println("done processing")
}
func help(help *base.Help) bool {
	if len(help.Value) == 0 {
		return false
	}

	if help.Value == "usage" {
		flag.Usage()
		return true
	}

	if help.Value == "config-example" {
		yaml.ShowExample()
		return true
	}

	return false
}
func safeVisit(ctx base.Context) error {
	lineProcessors := lineprocessors.CreateLineProcessors(ctx)
	if ctx.GetDest().GetType() == "std" {
		return visitor.Visit(ctx, std.NewHandler(lineProcessors))
	}

	if ctx.GetDest().GetType() == "file" {
		out, err := os.Create(ctx.GetDest().GetName())
		if err != nil {
			return err
		}
		defer out.Close()

		writer := bufio.NewWriter(out)
		return visitor.Visit(ctx, file.NewHandler(writer, lineProcessors))
	}

	if ctx.GetDest().GetType() == "dir" {
		return visitor.Visit(ctx, dir.NewHandler(lineProcessors))
	}

	return fmt.Errorf("no handler found")
}
