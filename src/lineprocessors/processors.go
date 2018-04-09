package lineprocessors

import (
	"../base"
	"strings"
	"fmt"
)

func CreateLineProcessors(ctx base.Context) []func(string) string {
	return []func(string) string{
		createNSProcess(ctx.GetDest().GetOldNS(), ctx.GetDest().GetNewNS()),
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
