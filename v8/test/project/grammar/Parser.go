/*
................................................................................
.    Copyright (c) 2009-2026 Crater Dog Technologies™.  All Rights Reserved.   .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

/*
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│              This class file was automatically generated using:              │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package grammar

import (
	fmt "fmt"
	com "github.com/craterdog/go-essential-composites/v8"
	uti "github.com/craterdog/go-essential-utilities/v8"
	ast "github.com/craterdog/go-example-syntax/ast"
	mat "math"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func ParserClass() ParserClassLike {
	return parserClass()
}

// Constructor Methods

func (c *parserClass_) Parser() ParserLike {
	var instance = &parser_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *parser_) GetClass() ParserClassLike {
	return parserClass()
}

func (v *parser_) ParseSource(
	source string,
) ast.DocumentLike {
	v.source_ = sts.ReplaceAll(source, "\t", "    ")
	v.tokens_ = com.Queue[TokenLike]()
	v.next_ = com.Stack[TokenLike]()

	// The scanner runs in a separate Go routine.
	ScannerClass().Scanner(v.source_, v.tokens_)

	// Attempt to parse the document.
	var document, token, ok = v.parseDocument()
	if !ok || v.tokens_.GetSize() > 1 {
		var message = v.formatError("$Document", token)
		panic(message)
	}
	return document
}

// PROTECTED INTERFACE

// Private Methods

func (v *parser_) parseAdditional() (
	additional ast.AdditionalLike,
	token TokenLike,
	ok bool,
) {
	var tokens = com.List[TokenLike]()

	// Attempt to parse a single "," literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter(",")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Additional rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Additional", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Component rule.
	var component ast.ComponentLike
	component, token, ok = v.parseComponent()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Component rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Additional", token)
		panic(message)
	}

	// Found a single Additional rule.
	ok = true
	v.remove(tokens)
	additional = ast.AdditionalClass().Additional(
		delimiter,
		component,
	)
	return
}

func (v *parser_) parseComponent() (
	component ast.ComponentLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Intrinsic Component.
	var intrinsic ast.IntrinsicLike
	intrinsic, token, ok = v.parseIntrinsic()
	if ok {
		// Found a single Intrinsic Component.
		component = ast.ComponentClass().Component(intrinsic)
		return
	}

	// Attempt to parse a single List Component.
	var list ast.ListLike
	list, token, ok = v.parseList()
	if ok {
		// Found a single List Component.
		component = ast.ComponentClass().Component(list)
		return
	}

	// Attempt to parse a single Keyword Component.
	var keyword ast.KeywordLike
	keyword, token, ok = v.parseKeyword()
	if ok {
		// Found a single Keyword Component.
		component = ast.ComponentClass().Component(keyword)
		return
	}

	// This is not a single Component rule.
	return
}

func (v *parser_) parseDocument() (
	document ast.DocumentLike,
	token TokenLike,
	ok bool,
) {
	var tokens = com.List[TokenLike]()

	// Attempt to parse multiple Component rules.
	var components = com.List[ast.ComponentLike]()
componentsLoop:
	for count_ := 0; count_ < mat.MaxInt; count_++ {
		var component ast.ComponentLike
		component, token, ok = v.parseComponent()
		if !ok {
			switch {
			case count_ >= 3:
				break componentsLoop
			case uti.IsDefined(tokens):
				// This is not multiple Component rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Document", token)
				message += "3 or more Component rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		components.AppendValue(component)
	}

	// Found a single Document rule.
	ok = true
	v.remove(tokens)
	document = ast.DocumentClass().Document(components)
	return
}

func (v *parser_) parseIntrinsic() (
	intrinsic ast.IntrinsicLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single number Intrinsic.
	var number string
	number, token, ok = v.parseToken(NumberToken)
	if ok {
		// Found a single number Intrinsic.
		intrinsic = ast.IntrinsicClass().Intrinsic(number)
		return
	}

	// Attempt to parse a single rune Intrinsic.
	var rune_ string
	rune_, token, ok = v.parseToken(RuneToken)
	if ok {
		// Found a single rune Intrinsic.
		intrinsic = ast.IntrinsicClass().Intrinsic(rune_)
		return
	}

	// Attempt to parse a single text Intrinsic.
	var text string
	text, token, ok = v.parseToken(TextToken)
	if ok {
		// Found a single text Intrinsic.
		intrinsic = ast.IntrinsicClass().Intrinsic(text)
		return
	}

	// This is not a single Intrinsic rule.
	return
}

func (v *parser_) parseKeyword() (
	keyword ast.KeywordLike,
	token TokenLike,
	ok bool,
) {
	var delimiter string

	// Attempt to parse a single "if" delimiter.
	delimiter, token, ok = v.parseDelimiter("if")
	if ok {
		// Found a single "if" delimiter.
		keyword = ast.KeywordClass().Keyword(delimiter)
		return
	}

	// Attempt to parse a single "then" delimiter.
	delimiter, token, ok = v.parseDelimiter("then")
	if ok {
		// Found a single "then" delimiter.
		keyword = ast.KeywordClass().Keyword(delimiter)
		return
	}

	// Attempt to parse a single "else" delimiter.
	delimiter, token, ok = v.parseDelimiter("else")
	if ok {
		// Found a single "else" delimiter.
		keyword = ast.KeywordClass().Keyword(delimiter)
		return
	}

	// This is not a single Keyword rule.
	return
}

func (v *parser_) parseList() (
	list ast.ListLike,
	token TokenLike,
	ok bool,
) {
	var tokens = com.List[TokenLike]()

	// Attempt to parse a single "[" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single List rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$List", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Component rule.
	var component ast.ComponentLike
	component, token, ok = v.parseComponent()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Component rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$List", token)
		panic(message)
	}

	// Attempt to parse multiple Additional rules.
	var additionals = com.List[ast.AdditionalLike]()
additionalsLoop:
	for count_ := 0; count_ < mat.MaxInt; count_++ {
		var additional ast.AdditionalLike
		additional, token, ok = v.parseAdditional()
		if !ok {
			switch {
			case count_ >= 0:
				break additionalsLoop
			case uti.IsDefined(tokens):
				// This is not multiple Additional rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$List", token)
				message += "0 or more Additional rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		additionals.AppendValue(additional)
	}

	// Attempt to parse a single "]" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single List rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$List", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single List rule.
	ok = true
	v.remove(tokens)
	list = ast.ListClass().List(
		delimiter1,
		component,
		additionals,
		delimiter2,
	)
	return
}

func (v *parser_) parseDelimiter(
	literal string,
) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single delimiter.
	value, token, ok = v.parseToken(DelimiterToken)
	if ok {
		if value == literal {
			// Found the desired delimiter.
			return
		}
		v.next_.AddValue(token)
		ok = false
	}

	// This is not the desired delimiter.
	return
}

func (v *parser_) parseToken(
	tokenType TokenType,
) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a specific token type.
	var tokens = com.List[TokenLike]()
	token = v.getNextToken()
	for token != nil {
		tokens.AppendValue(token)
		switch token.GetType() {
		case tokenType:
			// Found the desired token type.
			value = token.GetValue()
			ok = true
			return
		case SpaceToken, NewlineToken:
			// Ignore any unrequested whitespace.
			token = v.getNextToken()
		default:
			// This is not the desired token type.
			v.putBack(tokens)
			return
		}
	}

	// We are at the end-of-file marker.
	return
}

func (v *parser_) formatError(
	ruleName string,
	token TokenLike,
) string {
	// Format the error message.
	var message = fmt.Sprintf(
		"An unexpected token was received by the parser: %v\n",
		ScannerClass().FormatToken(token),
	)
	var line = token.GetLine()
	var lines = sts.Split(v.source_, "\n")

	// Append the source lines with the error in it.
	message += "\033[36m"
	for index := line - 3; index < line; index++ {
		if index > 1 {
			message += fmt.Sprintf("%04d: ", index) + string(lines[index-1]) + "\n"
		}
	}
	message += fmt.Sprintf("%04d: ", line) + string(lines[line-1]) + "\n"

	// Append an arrow pointing to the error.
	message += " \033[32m>>>─"
	var count uint
	for count < token.GetPosition() {
		message += "─"
		count++
	}
	message += "⌃\033[36m\n"

	// Append the following source line for context.
	if line < uti.ArraySize(lines) {
		message += fmt.Sprintf("%04d: ", line+1) + string(lines[line]) + "\n"
	}
	message += "\033[0m\n"
	if uti.IsDefined(ruleName) {
		message += "Was expecting:\n"
		message += fmt.Sprintf(
			"  \033[32m%v: \033[33m%v\033[0m\n\n",
			ruleName,
			v.getDefinition(ruleName),
		)
	}
	return message
}

func (v *parser_) getDefinition(
	ruleName string,
) string {
	var syntax = parserClass().syntax_
	var definition = syntax.GetValue(ruleName)
	return definition
}

func (v *parser_) getNextToken() TokenLike {
	// Check for any read, but unprocessed tokens.
	if !v.next_.IsEmpty() {
		return v.next_.RemoveLast()
	}

	// Read a new token from the token stream.
	var token, ok = v.tokens_.RemoveFirst() // This will wait for a token.
	if !ok {
		// The token channel has been closed.
		return nil
	}

	// Check for an error token.
	if token.GetType() == ErrorToken {
		var message = v.formatError("", token)
		panic(message)
	}

	return token
}

func (v *parser_) putBack(
	tokens com.Sequential[TokenLike],
) {
	var iterator = tokens.GetIterator()
	for iterator.ToEnd(); iterator.HasPrevious(); {
		var token = iterator.GetPrevious()
		v.next_.AddValue(token)
	}
}

// NOTE:
// This method does nothing but must exist to satisfy the lint check on the
// generated parser code.  The generated code must call this method is some
// cases to make it look that the tokens variable is being used somewhere.
func (v *parser_) remove(
	tokens com.Sequential[TokenLike],
) {
}

// Instance Structure

type parser_ struct {
	// Declare the instance attributes.
	source_ string                   // The original source code.
	tokens_ com.QueueLike[TokenLike] // A queue of unread tokens from the scanner.
	next_   com.StackLike[TokenLike] // A stack of read, but unprocessed tokens.
}

// Class Structure

type parserClass_ struct {
	// Declare the class constants.
	syntax_ com.CatalogLike[string, string]
}

// Class Reference

func parserClass() *parserClass_ {
	return parserClassReference_
}

var parserClassReference_ = &parserClass_{
	// Initialize the class constants.
	syntax_: com.CatalogFromMap[string, string](
		map[string]string{
			"$Document": `Component{3..}  ! An inline comment.`,
			"$Component": `
    Intrinsic
    List
    Keyword`,
			"$Intrinsic": `
    number
    rune  ! A multiline comment.
    text`,
			"$List":       `"[" Component Additional* "]"`,
			"$Additional": `"," Component`,
			"$Keyword": `
    "if"
    "then"
    "else"`,
		},
	),
}
