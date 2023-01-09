package generator

import "fmt"

func HandleQuery(querySlice []string) string {
	finalQuery := ""

	for _, query := range querySlice {
		tokenStream := generateTokens(query)
		isValidInput, msg := simpleSyntaxCheck(tokenStream)

		if !isValidInput {
			return msg
		}

		finalQuery += query
	}

	output := fmt.Sprintf("{\n\t\"__query__\": %v\n}", finalQuery)

	return output
}
