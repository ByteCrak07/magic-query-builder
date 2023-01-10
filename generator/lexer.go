package generator

import "regexp"

type Token struct {
	Type  string
	Value string
}

func generateTokens(query string) []Token {
	var tokenStream []Token

	tokenTypes := []string{"STRING", "ID", "CONST", "DOT", "EQ", "SQBRAC"}

	// map of regular expressions used
	var regExps = make(map[string]string)
	regExps["STRING"] = `\".+\"`
	regExps["ID"] = `[_a-zA-Z][_a-zA-Z0-9]*`
	regExps["CONST"] = `[0-9]+`
	regExps["DOT"] = `\.`
	regExps["EQ"] = `\=`
	regExps["SQBRAC"] = `\[\]`

	// combining all regular expressions by or for splitting string into tokens
	completeExp := ""
	for _, exp := range regExps {
		completeExp += exp + "|"
	}
	completeExp = completeExp[0 : len(completeExp)-1] // for removing the extra |(or) symbol
	splitString := regexp.MustCompile(completeExp)
	splitStringsSlice := splitString.FindAllString(query, -1) // splice containing all tokens

	// taking each token and checking its token type
	var isToken bool
	for _, token := range splitStringsSlice {
		for _, tokenType := range tokenTypes {
			isToken, _ = regexp.MatchString(regExps[tokenType], token)

			// if token type is matched add the token to the token stream along with its type
			if isToken {
				var tempToken Token
				tempToken.Type = tokenType
				tempToken.Value = token
				tokenStream = append(tokenStream, tempToken)
				break
			}
		}

	}

	return tokenStream
}
