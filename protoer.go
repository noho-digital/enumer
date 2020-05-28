package main

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"strings"
)

func (g *Generator) buildProtoEnum(values []Value, typeName string) {
	tmpl := `enum %s {
%s
}
`
	var el []string
	for i, value := range values {
		n := value.name
		if i == 0 {
			n = strings.Join([]string{n, strcase.ToScreamingSnake(typeName)}, "_")
		}
		el = append(el, fmt.Sprintf("    %s = %s;", n, value.str))
	}
	g.Printf(fmt.Sprintf(tmpl, typeName, strings.Join(el, "\n")))
}
