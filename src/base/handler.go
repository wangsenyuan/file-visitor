package base

type Handler interface {
	BeforeProcess(context Context) error
	HandleFileContent(ctx Context, line string) error
	AfterProcess(ctx Context) error
	BeforeHandleFile(context Context, file string) error
	AfterHandleFile(ctx Context, file string) error
	BeforeHandleDir(ctx Context, dir string) error
	AfterHandleDir(ctx Context, dir string) error
}
