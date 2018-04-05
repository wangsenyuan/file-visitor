package base

import (
	"fmt"
	"flag"
)

type Context struct {
	src     *Src
	out     *Out
	dest    *Dest
	comment *Comment
	on      *OldNamespace
	nn      *NewNamepsace
}

func NewContext(src *Src, out *Out, dest *Dest, c *Comment, on *OldNamespace, nn *NewNamepsace) Context {
	return Context{src, out, dest, c, on, nn}
}

func (ctx Context) IsStd() bool {
	return ctx.out.tp == "std"
}

func (ctx Context) IsDir() bool {
	return ctx.out.tp == "dir"
}

func (ctx Context) IsFile() bool {
	return ctx.out.tp == "file"
}

func (ctx Context) IsComment() bool {
	return ctx.comment.val
}

func (ctx Context) GetSrc() string {
	if ctx.src == nil {
		return "."
	}
	return ctx.src.name
}

func (ctx Context) GetDest() string {
	return ctx.dest.name
}

func (ctx Context) GetOldNamespace() string {
	return ctx.on.val
}

func (ctx Context) GetNewNamespace() string {
	return ctx.nn.val
}
func (ctx Context) IsValid() bool {
	if len(ctx.GetDest()) == 0 {
		return false
	}

	if !ctx.IsDir() && !ctx.IsFile() && !ctx.IsStd() {
		return false
	}

	if len(ctx.GetSrc()) == 0 {
		return false
	}

	return true
}

type Out struct {
	tp string
}

func (o *Out) Set(s string) error {
	if s != "std" && s != "dir" && s != "file" {
		return fmt.Errorf("invalid out type, only std, dir or file are supported")
	}
	o.tp = s
	return nil
}

func (o *Out) String() string {
	return fmt.Sprintf("-out %s", o.tp)
}

func OutFlag() *Out {
	o := &Out{}

	flag.CommandLine.Var(o, "out", "输出结果到std, dir或者file")

	return o
}

type Dest struct {
	name string
}

func (d *Dest) Set(s string) error {
	d.name = s
	return nil
}

func (d *Dest) String() string {
	return fmt.Sprintf("-dest %s", d.name)
}

func DestFlag() *Dest {
	d := &Dest{}
	flag.CommandLine.Var(d, "dest", "输出的目录或文件名，相对当前位置")
	return d
}

type Src struct {
	name string
}

func (src *Src) Set(s string) error {
	if len(s) == 0 {
		src.name = "."
	} else {
		src.name = s
	}
	return nil
}

func (src *Src) String() string {
	return fmt.Sprintf("-src %s", src.name)
}

func SrcFlag() *Src {
	s := &Src{}
	flag.CommandLine.Var(s, "src", "输入文件的目录或者文件名， 相对当前位置")
	return s
}

type Comment struct {
	val bool
}

func (c *Comment) Set(s string) error {
	if len(s) == 0 {
		return nil
	}
	fmt.Sscanf(s, "%t", &c.val)
	return nil
}

func (c *Comment) String() string {
	return fmt.Sprintf("-comment %t", c.val)
}

func CommentFlag() *Comment {
	c := &Comment{}

	flag.CommandLine.Var(c, "comment", "是否在结果中输出目录和文件名")

	return c
}

type OldNamespace struct {
	val string
}

func (on *OldNamespace) Set(s string) error {
	if len(s) == 0 {
		on.val = "maycur"
	} else {
		on.val = s
	}
	return nil
}

func (on *OldNamespace) String() string {
	return fmt.Sprintf("-old-namespace %s", on.val)
}

func OldNamepsaceFlag() *OldNamespace {
	on := &OldNamespace{"maycur"}
	flag.CommandLine.Var(on, "old-namespace", "要替换的namespace")
	return on
}

type NewNamepsace struct {
	val string
}

func (on *NewNamepsace) Set(s string) error {
	if len(s) == 0 {
		on.val = "maycur"
	} else {
		on.val = s
	}
	return nil
}

func (on *NewNamepsace) String() string {
	return fmt.Sprintf("-new-namespace %s", on.val)
}

func NewNamepsaceFlag() *NewNamepsace {
	on := &NewNamepsace{"maycur"}
	flag.CommandLine.Var(on, "new-namespace", "要替换的namespace")
	return on
}
