package file

import (
	"../../base"
	"bufio"
)

type FileHandler struct {
	Writer         *bufio.Writer
	LineProcessors []func(string) string
	LineCount      int
}

func NewHandler(writer *bufio.Writer, lineProcessors []func(string) string) *FileHandler {
	return &FileHandler{writer, lineProcessors, 0}
}

func (this *FileHandler) BeforeProcess(ctx base.Context) error {
	return nil
}

func (this *FileHandler) BeforeHandleDir(ctx base.Context, dir string) error {
	return nil
}

func (this *FileHandler) BeforeHandleFile(ctx base.Context, file string) error {
	_, err := this.Writer.WriteString("---\n")
	return err
}

func (this *FileHandler) HandleFileContent(ctx base.Context, line string) error {
	if len(this.LineProcessors) != 0 {
		for _, processor := range this.LineProcessors {
			line = processor(line)
		}
	}

	if len(line) == 0 && ctx.GetDest().ShouldRemoveComment() {
		return nil
	}

	_, err := this.Writer.WriteString(line)
	this.LineCount++
	return err
}

func (this *FileHandler) AfterHandleFile(ctx base.Context, file string) error {
	return nil
}

func (this *FileHandler) AfterHandleDir(ctx base.Context, dir string) error {
	return nil
}

func (this *FileHandler) AfterProcess(ctx base.Context) error {
	return this.Writer.Flush()
}
