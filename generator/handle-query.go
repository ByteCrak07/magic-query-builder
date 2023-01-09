package generator

import "fmt"

func HandleQuery(querySlice []string) string {
	var finalQuery string

	for index, query := range querySlice {
		// get token stream for a query
		tokenStream := generateTokens(query)

		// check syntax of query using the token stream
		isValidInput, msg := simpleSyntaxCheck(tokenStream)

		if !isValidInput {
			return msg
		}

		finalQuery += generateQuery(tokenStream)

		if index != len(querySlice)-1 {
			finalQuery += ",\n\t"
		}
	}

	// wrap output in braces and assign __query__
	output := fmt.Sprintf("{\n\t\"__query__\": %v\n}", finalQuery)

	return output
}
