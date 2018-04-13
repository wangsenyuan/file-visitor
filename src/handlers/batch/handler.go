package batch

import (
	"../../base"
	"os"
	"../file"
	"../../util"
	"bufio"
	"fmt"
)

type BatchHandler struct {
	fileHandler    *file.FileHandler
	currentFile    *os.File
	LineProcessors []func(string) string
	batchIndex     int
}

func NewHandler(lineProcessors []func(string) string) *BatchHandler {
	return &BatchHandler{LineProcessors: lineProcessors}
}

func (this *BatchHandler) BeforeProcess(ctx base.Context) error {
	err := util.CreateDir(ctx.GetDest().GetName())
	if err != nil {
		return err
	}

	return this.createFileHandler(ctx)
}

func (this *BatchHandler) createFileHandler(ctx base.Context) error {
	out, err := os.Create(fmt.Sprintf("%s/batch_%d.%s", ctx.GetDest().GetName(), this.batchIndex, ctx.GetSrc().GetSuffix()))
	if err != nil {
		return err
	}

	fileHandler := file.NewHandler(bufio.NewWriter(out), this.LineProcessors)

	this.currentFile = out
	this.fileHandler = fileHandler
	this.batchIndex++

	return nil
}

func (this *BatchHandler) BeforeHandleDir(ctx base.Context, dir string) error {
	return nil
}

func (this *BatchHandler) BeforeHandleFile(ctx base.Context, name string) error {
	return nil
}

func (this *BatchHandler) HandleFileContent(ctx base.Context, line string) error {
	return this.fileHandler.HandleFileContent(ctx, line)
}

func (this *BatchHandler) AfterHandleFile(ctx base.Context, file string) error {
	//fmt.Printf("[debug] after handle file %s\n", file)
	this.fileHandler.AfterHandleFile(ctx, file)
	this.fileHandler.AfterProcess(ctx)

	if ctx.GetDest().GetFileLineLimit() <= 0 {
		return nil
	}
	if this.fileHandler.LineCount > ctx.GetDest().GetFileLineLimit() {
		err := this.currentFile.Close()

		if err != nil {
			return err
		}

		this.fileHandler = nil
		this.currentFile = nil

		return this.createFileHandler(ctx)
	}
	return nil
}

func (this *BatchHandler) AfterHandleDir(ctx base.Context, dir string) error {
	return nil
}

func (this *BatchHandler) AfterProcess(ctx base.Context) error {
	if this.fileHandler != nil {
		this.fileHandler.AfterProcess(ctx)
		this.fileHandler = nil
		this.currentFile.Close()
	}
	return nil
}
