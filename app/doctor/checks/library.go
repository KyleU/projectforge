package checks

import (
	"context"

	"github.com/samber/lo"

	"projectforge.dev/projectforge/app/doctor"
	"projectforge.dev/projectforge/app/lib/telemetry"
	"projectforge.dev/projectforge/app/util"
)

var AllChecks = doctor.Checks{
	PF, Homebrew, Choco, Golang, Make, Node, NPM, Git, Air, Gofumpt,
	Golangcilint, Gotestsum, QTC, Imagemagick, Inkscape, Repo, Project,
}

func GetCheck(key string) *doctor.Check {
	return AllChecks.Get(key)
}

func Core(core bool) doctor.Checks {
	return lo.Filter(AllChecks, func(c *doctor.Check, _ int) bool {
		return c.Core == core
	})
}

func ForModules(modules []string) doctor.Checks {
	var ret doctor.Checks
	lo.ForEach(AllChecks, func(c *doctor.Check, _ int) {
		if !c.Applies() {
			return
		}
		hit := len(c.Modules) == 0 || lo.ContainsBy(c.Modules, func(mod string) bool {
			return lo.Contains(modules, mod)
		})
		if !hit {
			return
		}
		ret = append(ret, c)
	})
	return ret
}

func CheckAll(ctx context.Context, modules []string, logger util.Logger, exclude ...string) doctor.Results {
	ctx, span, logger := telemetry.StartSpan(ctx, "doctor:checkall", logger)
	defer span.Complete()

	return lo.FilterMap(ForModules(modules), func(c *doctor.Check, _ int) (*doctor.Result, bool) {
		if lo.Contains(exclude, c.Key) {
			return nil, false
		}
		return c.Check(ctx, logger), true
	})
}

func noop(_ context.Context, r *doctor.Result, _ string) *doctor.Result {
	return r
}
