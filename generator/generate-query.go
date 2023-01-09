package generator

import (
	"fmt"
	"strings"
)

func generateQuery(tokenStream []Token) string {
	var query string
	var maxDepth uint

	// finding depth of tokens for properly aligning tabs
	for i := 0; i < len(tokenStream); i++ {
		token := tokenStream[i]

		if token.Type == "ID" || token.Type == "SQBRAC" {
			maxDepth++
		}
	}

	// taking last token and parsing in reverse order
	for i := len(tokenStream) - 1; i >= 0; i-- {
		token := tokenStream[i]
		tab1 := strings.Repeat("\t", int(maxDepth+1))
		tab2 := strings.Repeat("\t", int(maxDepth))

		if token.Type == "CONST" || token.Type == "STRING" {
			query = token.Value
		}

		// wrap previous generated query with braces and
		if token.Type == "SQBRAC" {
			if tokenStream[i+1].Type == "EQ" {
				// assign __eq__ for [] when [] = query
				query = fmt.Sprintf("{\n%v\"__eq__\": %v\n%v}", tab1, query, tab2)
			} else {
				// assign __match__ for [] when [].identifier
				query = fmt.Sprintf("{\n%v\"__match__\": %v\n%v}", tab1, query, tab2)
			}
			maxDepth--
		}

		// wrap previous generated query with braces and assign identifier
		if token.Type == "ID" {
			query = fmt.Sprintf("{\n%v\"%v\": %v\n%v}", tab1, token.Value, query, tab2)
			maxDepth--
		}
	}

	return query
}
