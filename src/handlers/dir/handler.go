package dir

import (
	"bufio"
	"../../base"
	"os"
	"../file"
	"../../util"
)

type DirHandler struct {
	fileHandler    *file.FileHandler
	currentFile    *os.File
	LineProcessors []func(string) string
}

func NewHandler(lineProcessors []func(string) string) *DirHandler {
	return &DirHandler{LineProcessors: lineProcessors}
}

func (this *DirHandler) BeforeProcess(ctx base.Context) error {
	return util.CreateDir(ctx.GetDest().GetName())
}

func (this *DirHandler) BeforeHandleDir(ctx base.Context, dir string) error {
	return util.CreateDir(dir)
}

func (this *DirHandler) BeforeHandleFile(ctx base.Context, name string) error {
	out, err := os.Create(name)
	if err != nil {
		return err
	}

	fileHandler := file.NewHandler(bufio.NewWriter(out), this.LineProcessors)

	this.currentFile = out
	this.fileHandler = fileHandler

	return this.fileHandler.BeforeHandleFile(ctx, name)
}

func (this *DirHandler) HandleFileContent(ctx base.Context, line string) error {
	return this.fileHandler.HandleFileContent(ctx, line)
}

func (this *DirHandler) AfterHandleFile(ctx base.Context, file string) error {
	err := this.fileHandler.AfterHandleFile(ctx, file)
	if err != nil {
		return err
	}
	err = this.fileHandler.AfterProcess(ctx)

	if err != nil {
		return err
	}
	err = this.currentFile.Close()

	if err != nil {
		return err
	}

	this.fileHandler = nil
	this.currentFile = nil

	return nil
}

func (this *DirHandler) AfterHandleDir(ctx base.Context, dir string) error {
	return nil
}

func (this *DirHandler) AfterProcess(ctx base.Context) error {
	return nil
}
