package checks

import (
	"context"

	"projectforge.dev/projectforge/app/doctor"
	"projectforge.dev/projectforge/app/util"
)

var choco = &doctor.Check{
	Key:       "choco",
	Section:   "build",
	Title:     "Chocolatey",
	Summary:   "Used to install other dependencies",
	URL:       "https://chocolatey.org",
	UsedBy:    "Package manager for Windows",
	Platforms: []string{"windows"},
	Fn:        simpleOut(".", "choco", []string{"--help"}, noop),
	Solve:     solveChoco,
}

func solveChoco(_ context.Context, r *doctor.Result, _ util.Logger) *doctor.Result {
	if r.Errors.Find("missing") != nil || r.Errors.Find("exitcode") != nil {
		r.AddSolution("#https://chocolatey.org/install")
	}
	return r
}
