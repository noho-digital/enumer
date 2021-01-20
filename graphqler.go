package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
)

func (g *Generator) buildGraphQL(types []string, dir string) {
	var parts []string
	header := "# Code generated by \"enumer %s\"; DO NOT EDIT.\n\n"
	parts = append(parts, fmt.Sprintf(header, strings.Join(os.Args[1:], " ")))
	// Run generate for each type.
	for _, typeName := range types {
		enumName := strings.Title(g.pkg.name) + typeName
		if graphqlPrefix != nil && len(*graphqlPrefix) > 0 {
			enumName = *graphqlPrefix + typeName
		}
		values := g.values(typeName, typeName, "screaming-snake", false)
		stanza := g.createGraphQLEnumStanza(values, enumName, typeName)
		parts = append(parts, stanza)
	}

	// Format the output
	graphqlPath := *graphqlOutput
	if graphqlPath == "" {
		baseName := defaultBaseName
		if len(types) == 1 {
			baseName = types[0]
		}
		name := strcase.ToSnake(fmt.Sprintf("%s_gen.graphql", baseName))
		graphqlPath = filepath.Join(dir, *graphqlDir, name)
	}
	text := strings.Join(parts, "\n")
	write(graphqlPath, []byte(text))
}

func (g *Generator) createGraphQLEnumStanza(values []Value, enumName string, typeName string) string {
	tmpl := `enum %s {
%s
}
`
	var el []string
	for i, value := range values {
		n := value.name
		if i == 0 {
			if *gqlShouldSuffixUndef && strings.ToLower(value.name) == "undefined" {
				suffix := typeName
				if *gqlUndefSuffix != "" {
					suffix = *gqlUndefSuffix
				}
				n = strings.Join([]string{n, strcase.ToScreamingSnake(suffix)}, "_")
			}
		}
		el = append(el, fmt.Sprintf("    %s", n))
	}
	return fmt.Sprintf(tmpl, enumName, strings.Join(el, "\n"))
}
