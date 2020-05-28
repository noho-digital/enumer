package main

import (
	"fmt"
	"strings"
)

func (g *Generator) buildProtoEnum(runs [][]Value, typeName string, _ int)  {
	tmpl := `enum %s {
%s
}
`
	var el []string
	for _, values := range runs {
		for _, value := range values {
			el = append(el, fmt.Sprintf("    %s = %s;", value.name, value.str))
		}
	}
	g.Printf(fmt.Sprintf(tmpl, typeName, strings.Join(el, "\n")))
}