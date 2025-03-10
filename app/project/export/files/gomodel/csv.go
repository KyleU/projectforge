package gomodel

import (
	"strings"

	"projectforge.dev/projectforge/app/lib/metamodel/model"
	"projectforge.dev/projectforge/app/project/export/files/helper"
	"projectforge.dev/projectforge/app/project/export/golang"
)

func modelStrings(g *golang.File, m *model.Model) *golang.Block {
	ret := golang.NewBlock("Strings", "func")
	ret.WF("func (%s *%s) Strings() []string {", m.FirstLetter(), m.Proper())
	x := m.Columns.NotDerived().ToGoStrings(m.FirstLetter()+".", true, 160)
	if strings.Contains(x, "fmt.") {
		g.AddImport(helper.ImpFmt)
	}
	ret.WF("\treturn []string{%s}", x)
	ret.W("}")
	return ret
}

func modelToCSV(m *model.Model) *golang.Block {
	ret := golang.NewBlock("ToCSV", "func")
	ret.WF("func (%s *%s) ToCSV() ([]string, [][]string) {", m.FirstLetter(), m.Proper())
	ret.WF("\treturn %sFieldDescs.Keys(), [][]string{%s.Strings()}", m.Proper(), m.FirstLetter())
	ret.W("}")
	return ret
}

func modelArrayToCSV(m *model.Model) *golang.Block {
	ret := golang.NewBlock("ToCSV", "func")
	ret.WF("func (%s %s) ToCSV() ([]string, [][]string) {", m.FirstLetter(), m.ProperPlural())
	ret.WF("\treturn %sFieldDescs.Keys(), lo.Map(%s, func(x *%s, _ int) []string {", m.Proper(), m.FirstLetter(), m.Proper())
	ret.W("\t\treturn x.Strings()")
	ret.W("\t})")
	ret.W("}")
	return ret
}
