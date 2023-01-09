package generator

func simpleSyntaxCheck(tokenStream []Token) (bool, string) {
	var msg string
	if len(tokenStream) < 3 {
		msg = "Syntax error: Input must be atleast <identifier> = <constant or string>"
		return false, msg
	}

	firstTokenType := tokenStream[0].Type
	lastTokenType := tokenStream[len(tokenStream)-1].Type
	secLastTokenType := tokenStream[len(tokenStream)-2].Type

	if firstTokenType != "ID" {
		msg = "Syntax error: Query input must begin with an identifier"
		return false, msg
	}

	if secLastTokenType != "EQ" {
		msg = "Syntax error: Query input must use = to assign it to a constant or a string"
		return false, msg
	}

	if lastTokenType != "STRING" && lastTokenType != "CONST" {
		msg = "Syntax error: Query input must end with a constant or a string"
		return false, msg
	}

	return true, ""
}
