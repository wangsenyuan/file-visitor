package file

import (
	"../../base"
	"fmt"
	"bufio"
	"strings"
)

type FileHandler struct {
	Writer *bufio.Writer
}

func NewHandler(writer *bufio.Writer) *FileHandler {
	return &FileHandler{Writer: writer}
}

func (this *FileHandler) BeforeProcess(ctx base.Context) error {
	if !ctx.IsComment() {
		return nil
	}
	_, err := this.Writer.WriteString("-------- begin --------\n")
	return err
}

func (this *FileHandler) BeforeHandleDir(ctx base.Context, dir string) error {
	if !ctx.IsComment() {
		return nil
	}
	_, err := this.Writer.WriteString(fmt.Sprintf("------- begin dir %s --------\n", dir))
	return err
}

func (this *FileHandler) BeforeHandleFile(ctx base.Context, file string) error {
	if !ctx.IsComment() {
		return nil
	}
	_, err := this.Writer.WriteString(fmt.Sprintf("------- begin file %s --------\n", file))
	return err
}

func (this *FileHandler) HandleFileContent(ctx base.Context, line string) error {
	_, err := this.Writer.WriteString(replaceNamespace(ctx, line))
	return err
}
func replaceNamespace(ctx base.Context, s string) string {
	if len(s) == 0 {
		return s
	}

	if len(ctx.GetOldNamespace()) == 0 || len(ctx.GetNewNamespace()) == 0 {
		return s
	}

	st := strings.Trim(s, " \n")
	tp := fmt.Sprintf("namespace: %s", ctx.GetOldNamespace())

	if st != tp {
		return s
	}

	return strings.Replace(s, tp, fmt.Sprintf("namespace: %s", ctx.GetNewNamespace()), 1)
}

func (this *FileHandler) AfterHandleFile(ctx base.Context, file string) error {
	if !ctx.IsComment() {
		return nil
	}
	_, err := this.Writer.WriteString(fmt.Sprintf("------- done file %s --------\n", file))
	return err
}

func (this *FileHandler) AfterHandleDir(ctx base.Context, dir string) error {
	if !ctx.IsComment() {
		return nil
	}
	_, err := this.Writer.WriteString(fmt.Sprintf("------- done dir %s --------\n", dir))
	return err
}

func (this *FileHandler) AfterProcess(ctx base.Context) error {
	if ctx.IsComment() {
		_, err := this.Writer.WriteString(fmt.Sprintf("------- done --------\n"))
		if err != nil {
			return err
		}
	}

	return this.Writer.Flush()
}
