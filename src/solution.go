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
)

var out = base.OutFlag()
var dest = base.DestFlag()
var src = base.SrcFlag()
var comment = base.CommentFlag()
var on = base.OldNamepsaceFlag()
var nn = base.NewNamepsaceFlag()

func main() {
	flag.Parse()

	ctx := base.NewContext(src, out, dest, comment, on, nn)

	if !ctx.IsValid() {
		flag.Usage()
		return
	}

	err := safeVisit(ctx)
	if err != nil {
		panic(err)
	}
	//fmt.Println("done processing")
}
func safeVisit(ctx base.Context) error {
	if ctx.IsStd() {
		return visitor.Visit(ctx, std.NewHandler())
	}

	if ctx.IsFile() {
		out, err := os.Create(ctx.GetDest())
		if err != nil {
			return err
		}
		defer out.Close()

		writer := bufio.NewWriter(out)
		return visitor.Visit(ctx, file.NewHandler(writer))
	}

	if ctx.IsDir() {
		return visitor.Visit(ctx, dir.NewHandler())
	}

	return fmt.Errorf("no handler found")
}
