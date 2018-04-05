package visitor

import "../base"
import (
	"../util"
	"os"
	"bufio"
	"io/ioutil"
	"sort"
	"io"
)

func Visit(ctx base.Context, handler base.Handler) error {
	err := handler.BeforeProcess(ctx)
	if err != nil {
		return err
	}

	src := ctx.GetSrc()

	if ok, err := util.IsFile(src); err != nil {
		return err
	} else if ok {
		tmp := visitFile(ctx, handler, src, ctx.GetDest())
		if tmp != nil {
			return tmp
		}
	} else {
		tmp := visitDir(ctx, handler, src, ctx.GetDest())
		if tmp != nil {
			return tmp
		}
	}

	return handler.AfterProcess(ctx)
}

func visitDir(ctx base.Context, handler base.Handler, src, dest string) error {
	handler.BeforeHandleDir(ctx, dest)

	files, err := ioutil.ReadDir(src)

	if err != nil {
		return err
	}

	sort.Sort(util.Files(files))

	for _, file := range files {
		if file.IsDir() {
			err := visitDir(ctx, handler, src+"/"+file.Name(), dest+"/"+file.Name())
			if err != nil {
				return err
			}
		} else {
			err := visitFile(ctx, handler, src+"/"+file.Name(), dest+"/"+file.Name())
			if err != nil {
				return err
			}
		}
	}

	return handler.AfterHandleDir(ctx, dest)
}

func visitFile(ctx base.Context, handler base.Handler, src, dest string) error {
	var err error
	err = handler.BeforeHandleFile(ctx, dest)
	if err != nil {
		return err
	}

	file, fileErr := os.Open(src)
	if fileErr != nil {
		return fileErr
	}

	defer file.Close()
	defer handler.AfterHandleFile(ctx, dest)

	scanner := bufio.NewReader(file)

	for {
		line, err := scanner.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		//fmt.Print("[debug] ---" + line)

		err2 := handler.HandleFileContent(ctx, line)

		if err2 != nil {
			return err
		}
	}
	//fmt.Printf("[debug] --- file %s is processed\n", name)

	return nil
}
