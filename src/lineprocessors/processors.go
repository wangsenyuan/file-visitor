package lineprocessors

import (
	"../base"
	"strings"
	"fmt"
)

func CreateLineProcessors(ctx base.Context) []func(string) string {
	processors := make([]func(string) string, 0, 10)

	processors = append(processors, createNSProcess(ctx.GetDest().GetOldNS(), ctx.GetDest().GetNewNS()))

	if ctx.GetDest().IsRemoveCopyright() {
		processors = append(processors, createCopyrightProcess())
	}

	if ctx.GetDest().ShouldRemoveComment() {
		processors = append(processors, createRemoveCommentProcessor())
	}

	return processors
}

func createRemoveCommentProcessor() func(string) string {
	return func(s string) string {
		sm := strings.TrimLeft(s, " \t")
		if strings.HasPrefix(sm, "//") {
			return ""
		}
		return s
	}
}

func createNSProcess(OldNS, NewNS string) func(string) string {
	fn := func(s string) string {
		if len(s) == 0 {
			return s
		}

		if len(OldNS) == 0 || len(NewNS) == 0 {
			return s
		}

		st := strings.Trim(s, " \n\t")
		tp := fmt.Sprintf("namespace: %s", OldNS)

		if st != tp {
			return s
		}

		return strings.Replace(s, tp, fmt.Sprintf("namespace: %s", NewNS), 1)
	}

	return fn
}

func createCopyrightProcess() func(string) string {
	fn := func(s string) string {
		sm := strings.ToLower(s)
		if strings.Contains(sm, "copyright") {
			return strings.Replace(sm, "copyright", "", -1)
		}
		return s
	}

	return fn
}
