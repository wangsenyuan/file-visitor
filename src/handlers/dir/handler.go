package dir

import (
	"bufio"
	"../../base"
	"os"
	"../file"
	"fmt"
)

type DirHandler struct {
	fileHandler *file.FileHandler
	currentFile *os.File
}

func NewHandler() *DirHandler {
	return &DirHandler{}
}

func (this *DirHandler) BeforeProcess(ctx base.Context) error {
	return createDir(ctx.GetDest())
}
func createDir(dest string) error {
	_, err := os.Stat(dest)
	if os.IsNotExist(err) {
		err = os.Mkdir(dest, os.ModePerm)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	fmt.Printf("[debug] directory %s\n", dest)

	return nil
}

func (this *DirHandler) BeforeHandleDir(ctx base.Context, dir string) error {
	return createDir(dir)
}

func (this *DirHandler) BeforeHandleFile(ctx base.Context, name string) error {
	out, err := os.Create(name)
	if err != nil {
		return err
	}

	fileHandler := file.NewHandler(bufio.NewWriter(out))

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
