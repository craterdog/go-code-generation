/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package synthesizer

import (
	ana "github.com/craterdog/go-code-generation/v5/analyzer"
	uti "github.com/craterdog/go-missing-utilities/v2"
	not "github.com/craterdog/go-syntax-notation/v5"
	gra "github.com/craterdog/go-syntax-notation/v5/grammar"
)

// CLASS INTERFACE

// Access Function

func ParserSynthesizer() ParserSynthesizerClassLike {
	return parserSynthesizerReference()
}

// Constructor Methods

func (c *parserSynthesizerClass_) Make(
	syntax not.SyntaxLike,
) ParserSynthesizerLike {
	var instance = &parserSynthesizer_{
		// Initialize the instance attributes.
		analyzer_: ana.SyntaxAnalyzer().Make(syntax),
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *parserSynthesizer_) GetClass() ParserSynthesizerClassLike {
	return parserSynthesizerReference()
}

// TemplateDriven Methods

func (v *parserSynthesizer_) CreateLegalNotice() string {
	var legalNotice = v.analyzer_.GetLegalNotice()
	return legalNotice
}

func (v *parserSynthesizer_) CreateAccessFunction() string {
	var accessFunction = parserSynthesizerReference().accessFunction_
	return accessFunction
}

func (v *parserSynthesizer_) CreateConstantMethods() string {
	var constantMethods string
	return constantMethods
}

func (v *parserSynthesizer_) CreateConstructorMethods() string {
	var constructorMethods = parserSynthesizerReference().constructorMethods_
	return constructorMethods
}

func (v *parserSynthesizer_) CreateFunctionMethods() string {
	var functionMethods string
	return functionMethods
}

func (v *parserSynthesizer_) CreatePrincipalMethods() string {
	var principalMethods = parserSynthesizerReference().principalMethods_
	return principalMethods
}

func (v *parserSynthesizer_) CreateAttributeMethods() string {
	var attributeMethods string
	return attributeMethods
}

func (v *parserSynthesizer_) CreateAspectMethods() string {
	var aspectMethods string
	return aspectMethods
}

func (v *parserSynthesizer_) CreatePrivateMethods() string {
	var parseMethods = v.createParseMethods()
	var privateMethods = parserSynthesizerReference().privateMethods_
	privateMethods = uti.ReplaceAll(
		privateMethods,
		"parseMethods",
		parseMethods,
	)
	return privateMethods
}

func (v *parserSynthesizer_) CreateInstanceStructure() string {
	var instanceStructure = parserSynthesizerReference().instanceStructure_
	return instanceStructure
}

func (v *parserSynthesizer_) CreateClassStructure() string {
	var classStructure = parserSynthesizerReference().classStructure_
	return classStructure
}

func (v *parserSynthesizer_) CreateClassReference() string {
	var classReference = parserSynthesizerReference().classReference_
	return classReference
}

func (v *parserSynthesizer_) PerformGlobalUpdates(
	source string,
) string {
	var syntaxMap = v.analyzer_.GetSyntaxMap()
	source = uti.ReplaceAll(
		source,
		"syntaxMap",
		syntaxMap,
	)
	var syntaxName = v.analyzer_.GetSyntaxName()
	source = uti.ReplaceAll(
		source,
		"syntaxName",
		syntaxName,
	)
	var importedPackages = parserSynthesizerReference().importedPackages_
	source = uti.ReplaceAll(
		source,
		"importedPackages",
		importedPackages,
	)
	return source
}

// PROTECTED INTERFACE

// Private Methods

func (v *parserSynthesizer_) createArguments(
	ruleName string,
) (
	arguments string,
) {
	var variableNames = v.analyzer_.GetVariables(ruleName).GetIterator()
	if variableNames.IsEmpty() {
		return arguments
	}

	// Define the first argument.
	if variableNames.GetSize() > 1 {
		// Use the multiline argument style.
		arguments += "\n\t\t"
	}
	var argument = variableNames.GetNext()
	arguments += uti.ReplaceAll(
		"<argument_>",
		"argument",
		argument,
	)

	// Define any additional arguments.
	for variableNames.HasNext() {
		arguments += ",\n\t\t"
		argument = variableNames.GetNext()
		arguments += uti.ReplaceAll(
			"<argument_>",
			"argument",
			argument,
		)
	}
	if variableNames.GetSize() > 1 {
		// Use the multiline argument style.
		arguments += ",\n\t"
	}

	return arguments
}

func (v *parserSynthesizer_) createCardinality(
	cardinality not.CardinalityLike,
	optionalCardinality string,
	requiredCardinality string,
	repeatedCardinality string,
) string {
	var unlimited = "mat.MaxInt"
	var first string
	var last = unlimited
	var actualCardinality = requiredCardinality
	if uti.IsUndefined(cardinality) {
		return actualCardinality
	}
	switch actual := cardinality.GetAny().(type) {
	case not.ConstrainedLike:
		actualCardinality = repeatedCardinality
		switch actual.GetAny().(string) {
		case "?":
			// This is the "{0..1}" case.
			first = "0"
			last = "1"
			actualCardinality = optionalCardinality
		case "*":
			// This is the "{0..}" case.
			first = "0"
		case "+":
			// This is the "{1..}" case.
			first = "1"
		}
	case not.QuantifiedLike:
		first = actual.GetNumber()
		var limit = actual.GetOptionalLimit()
		if uti.IsUndefined(limit) {
			// This is the "{m}" case.
			last = first
		} else {
			last = limit.GetOptionalNumber()
			if uti.IsUndefined(last) {
				// This is the "{m..}" case.
				last = unlimited
			}
			// This is the "{m..n}" case.
		}
	}
	actualCardinality = uti.ReplaceAll(
		actualCardinality,
		"first",
		first,
	)
	actualCardinality = uti.ReplaceAll(
		actualCardinality,
		"last",
		last,
	)
	return actualCardinality
}

func (v *parserSynthesizer_) createDelimiterStep() string {
	var delimiterStep = parserSynthesizerReference().parseDelimiterStep_
	return delimiterStep
}

func (v *parserSynthesizer_) createInlineImplementation(
	ruleName string,
) string {
	var implementation = parserSynthesizerReference().declarationStep_
	var terms = v.analyzer_.GetTerms(ruleName).GetIterator()
	var variables = v.analyzer_.GetVariables(ruleName).GetIterator()
	for terms.HasNext() {
		var term = terms.GetNext()
		var parseStep string
		switch actual := term.GetAny().(type) {
		case not.ReferenceLike:
			var cardinality = actual.GetOptionalCardinality()
			var variableName = variables.GetNext()
			var identifier = actual.GetIdentifier().GetAny().(string)
			var scannerClass = gra.Scanner()
			switch {
			case scannerClass.MatchesType(identifier, gra.LowercaseToken):
				parseStep = v.createTokenStep(cardinality)
			case scannerClass.MatchesType(identifier, gra.UppercaseToken):
				parseStep = v.createRuleStep(cardinality)
			}
			parseStep = uti.ReplaceAll(
				parseStep,
				"variableName",
				variableName,
			)
			parseStep = uti.ReplaceAll(
				parseStep,
				"identifier",
				identifier,
			)
		case string:
			var delimiter = actual
			parseStep = v.createDelimiterStep()
			parseStep = uti.ReplaceAll(
				parseStep,
				"delimiter",
				delimiter,
			)
		}
		implementation += parseStep
	}
	var foundStep = parserSynthesizerReference().parseFoundStep_
	var arguments = v.createArguments(ruleName)
	implementation += uti.ReplaceAll(
		foundStep,
		"arguments",
		arguments,
	)
	return implementation
}

func (v *parserSynthesizer_) createMultilineImplementation(
	ruleName string,
) string {
	var implementation string
	var identifiers = v.analyzer_.GetIdentifiers(ruleName).GetIterator()
	for identifiers.HasNext() {
		var identifier = identifiers.GetNext().GetAny().(string)
		var parseCase string
		var scannerClass = gra.Scanner()
		switch {
		case scannerClass.MatchesType(identifier, gra.LowercaseToken):
			parseCase = parserSynthesizerReference().parseTokenCase_
		case scannerClass.MatchesType(identifier, gra.UppercaseToken):
			parseCase = parserSynthesizerReference().parseRuleCase_
		}
		implementation += uti.ReplaceAll(
			parseCase,
			"identifier",
			identifier,
		)
	}
	implementation += parserSynthesizerReference().parseDefaultCase_
	return implementation
}

func (v *parserSynthesizer_) createParseMethod(
	ruleName string,
) string {
	var methodImplementation string
	var references = v.analyzer_.GetReferences(ruleName)
	if uti.IsDefined(references) {
		methodImplementation = v.createInlineImplementation(ruleName)
	} else {
		methodImplementation = v.createMultilineImplementation(ruleName)
	}
	var parseMethod = parserSynthesizerReference().parseMethod_
	parseMethod = uti.ReplaceAll(
		parseMethod,
		"methodImplementation",
		methodImplementation,
	)
	parseMethod = uti.ReplaceAll(
		parseMethod,
		"ruleName",
		ruleName,
	)
	return parseMethod
}

func (v *parserSynthesizer_) createParseMethods() string {
	var parseMethods string
	var ruleNames = v.analyzer_.GetRuleNames().GetIterator()
	for ruleNames.HasNext() {
		var ruleName = ruleNames.GetNext()
		parseMethods += v.createParseMethod(ruleName)
	}
	return parseMethods
}

func (v *parserSynthesizer_) createRuleStep(
	cardinality not.CardinalityLike,
) string {
	var optionalStep = parserSynthesizerReference().parseOptionalRuleStep_
	var requiredStep = parserSynthesizerReference().parseRequiredRuleStep_
	var repeatedStep = parserSynthesizerReference().parseRepeatedRuleStep_
	var ruleStep = v.createCardinality(
		cardinality,
		optionalStep,
		requiredStep,
		repeatedStep,
	)
	return ruleStep
}

func (v *parserSynthesizer_) createTokenStep(
	cardinality not.CardinalityLike,
) string {
	var optionalStep = parserSynthesizerReference().parseOptionalTokenStep_
	var requiredStep = parserSynthesizerReference().parseRequiredTokenStep_
	var repeatedStep = parserSynthesizerReference().parseRepeatedTokenStep_
	var tokenStep = v.createCardinality(
		cardinality,
		optionalStep,
		requiredStep,
		repeatedStep,
	)
	return tokenStep
}

// Instance Structure

type parserSynthesizer_ struct {
	// Declare the instance attributes.
	analyzer_ ana.SyntaxAnalyzerLike
}

// Class Structure

type parserSynthesizerClass_ struct {
	// Declare the class constants.
	importedPackages_       string
	accessFunction_         string
	constructorMethods_     string
	parseMethod_            string
	parseDelimiterStep_     string
	declarationStep_        string
	parseRequiredTokenStep_ string
	parseOptionalTokenStep_ string
	parseRepeatedTokenStep_ string
	parseRequiredRuleStep_  string
	parseOptionalRuleStep_  string
	parseRepeatedRuleStep_  string
	parseFoundStep_         string
	parseTokenCase_         string
	parseRuleCase_          string
	parseDefaultCase_       string
	principalMethods_       string
	privateMethods_         string
	possibleDelimiter_      string
	requiredDelimiter_      string
	instanceStructure_      string
	classStructure_         string
	classReference_         string
}

// Class Reference

func parserSynthesizerReference() *parserSynthesizerClass_ {
	return parserSynthesizerReference_
}

var parserSynthesizerReference_ = &parserSynthesizerClass_{
	// Initialize the class constants.
	importedPackages_: `
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
	ast "<ModuleName>/ast"
	mat "math"
	sts "strings"
`,

	accessFunction_: `
// Access Function

func Parser() ParserClassLike {
	return parserReference()
}
`,

	constructorMethods_: `
// Constructor Methods

func (c *parserClass_) Make() ParserLike {
	var instance = &parser_{
		// Initialize the instance attributes.
	}
	return instance
}
`,

	parseMethod_: `
func (v *parser_) parse<~RuleName>() (
	<ruleName_> ast.<~RuleName>Like,
	token TokenLike,
	ok bool,
) {<MethodImplementation>}
`,

	parseDelimiterStep_: `
	// Attempt to parse a single <delimiter> delimiter.
	_, token, ok = v.parseDelimiter(<delimiter>)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single <~RuleName> rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$<~RuleName>", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}
`,

	declarationStep_: `
	var tokens = col.List[TokenLike]()
`,

	parseRequiredTokenStep_: `
	// Attempt to parse a single <~identifier> token.
	var <variableName_> string
	<variableName_>, token, ok = v.parseToken(<~Identifier>Token)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single <~RuleName> rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$<~RuleName>", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}
`,

	parseOptionalTokenStep_: `
	// Attempt to parse an optional <~identifier> token.
	var <variableName_> string
	<variableName_>, token, ok = v.parseToken(<~Identifier>Token)
	if ok && uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}
`,

	parseRepeatedTokenStep_: `
	// Attempt to parse multiple <~identifier> tokens.
	var <variableName_> = col.List[string]()
<~variableName>Loop:
	for count := 0; count < <last>; count++ {
		var <identifier_> string
		<identifier_>, token, ok = v.parseToken(<~Identifier>Token)
		if !ok {
			switch {
			case count >= <first>:
				break <~variableName>Loop
			case uti.IsDefined(tokens):
				// This is not multiple <~identifier> tokens.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$<~RuleName>", token)
				message += "<first> or more <~identifier> tokens are required."
				panic(message)
			}
		}
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
		<variableName_>.AppendValue(<identifier_>)
	}
`,

	parseRequiredRuleStep_: `
	// Attempt to parse a single <~Identifier> rule.
	var <variableName_> ast.<~Identifier>Like
	<variableName_>, token, ok = v.parse<~Identifier>()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single <~RuleName> rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$<~RuleName>", token)
		panic(message)
	}
`,

	parseOptionalRuleStep_: `
	// Attempt to parse an optional <~Identifier> rule.
	var <variableName_> ast.<~Identifier>Like
	<variableName_>, _, ok = v.parse<~Identifier>()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}
`,

	parseRepeatedRuleStep_: `
	// Attempt to parse multiple <~Identifier> rules.
	var <variableName_> = col.List[ast.<~Identifier>Like]()
<~variableName>Loop:
	for count := 0; count < <last>; count++ {
		var <identifier_> ast.<~Identifier>Like
		<identifier_>, token, ok = v.parse<~Identifier>()
		if !ok {
			switch {
			case count >= <first>:
				break <~variableName>Loop
			case uti.IsDefined(tokens):
				// This is not multiple <~Identifier> rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$<~RuleName>", token)
				message += "<first> or more <~Identifier> rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		<variableName_>.AppendValue(<identifier_>)
	}
`,

	parseFoundStep_: `
	// Found a single <~RuleName> rule.
	ok = true
	v.remove(tokens)
	<ruleName_> = ast.<~RuleName>().Make(<Arguments>)
	return
`,

	parseTokenCase_: `
	// Attempt to parse a single <~identifier> <~RuleName>.
	var <identifier_> string
	<identifier_>, token, ok = v.parseToken(<~Identifier>Token)
	if ok {
		// Found a single <~identifier> <~RuleName>.
		<ruleName_> = ast.<~RuleName>().Make(<identifier_>)
		return
	}
`,

	parseRuleCase_: `
	// Attempt to parse a single <~Identifier> <~RuleName>.
	var <identifier_> ast.<~Identifier>Like
	<identifier_>, token, ok = v.parse<~Identifier>()
	if ok {
		// Found a single <~Identifier> <~RuleName>.
		<ruleName_> = ast.<~RuleName>().Make(<identifier_>)
		return
	}
`,

	parseDefaultCase_: `
	// This is not a single <~RuleName> rule.
	return
`,

	principalMethods_: `
// Principal Methods

func (v *parser_) GetClass() ParserClassLike {
	return parserReference()
}

func (v *parser_) ParseSource(
	source string,
) ast.<~SyntaxName>Like {
	v.source_ = source
	v.tokens_ = col.Queue[TokenLike](parserReference().queueSize_)
	v.next_ = col.Stack[TokenLike](parserReference().stackSize_)

	// The scanner runs in a separate Go routine.
	Scanner().Make(v.source_, v.tokens_)

	// Attempt to parse the <~syntaxName>.
	var <syntaxName_>, token, ok = v.parse<~SyntaxName>()
	if !ok {
		var message = v.formatError("$<~SyntaxName>", token)
		panic(message)
	}
	return <syntaxName_>
}
`,

	privateMethods_: `
// Private Methods
<ParseMethods>
func (v *parser_) parseDelimiter(
	expectedValue string,
) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single delimiter.
	value, token, ok = v.parseToken(DelimiterToken)
	if ok {
		if value == expectedValue {
			// Found the desired delimiter.
			return
		}
		v.next_.AddValue(token)
		ok = false
	}

	// This is not the desired delimiter.
	return
}

func (v *parser_) parseToken(tokenType TokenType) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a specific token type.
	var tokens = col.List[TokenLike]()
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
		Scanner().FormatToken(token),
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
	if line < uint(len(lines)) {
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
	return parserReference().syntax_.GetValue(ruleName)
}

func (v *parser_) getNextToken() TokenLike {
	// Check for any read, but unprocessed tokens.
	if !v.next_.IsEmpty() {
		return v.next_.RemoveTop()
	}

	// Read a new token from the token stream.
	var token, ok = v.tokens_.RemoveHead() // This will wait for a token.
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
	tokens abs.Sequential[TokenLike],
) {
	var iterator = tokens.GetIterator()
	for iterator.ToEnd(); iterator.HasPrevious(); {
		var token = iterator.GetPrevious()
		v.next_.AddValue(token)
	}
}

func (v *parser_) remove(
	tokens abs.Sequential[TokenLike],
) {
}
`,

	possibleDelimiter_: `
	// Attempt to parse a single "<delimiter>" delimiter.
	_, token, ok = v.parseDelimiter("<delimiter>")
	if !ok {
		// This is not a single <~RuleName> rule.
		return <ruleName_>, token, false
	}
`,

	requiredDelimiter_: `
	// Attempt to parse a single "<delimiter>" delimiter.
	_, token, ok = v.parseDelimiter("<delimiter>")
	if !ok {
		// Encountered a syntax error.
		var message = v.formatError("$<~RuleName>", token)
		panic(message)
	}
`,

	instanceStructure_: `
// Instance Structure

type parser_ struct {
	// Declare the instance attributes.
	source_ string                   // The original source code.
	tokens_ abs.QueueLike[TokenLike] // A queue of unread tokens from the scanner.
	next_   abs.StackLike[TokenLike] // A stack of read, but unprocessed tokens.
}
`,

	classStructure_: `
// Class Structure

type parserClass_ struct {
	// Declare the class constants.
	queueSize_ uint
	stackSize_ uint
	syntax_    abs.CatalogLike[string, string]
}
`,

	classReference_: `
// Class Reference

func parserReference() *parserClass_ {
	return parserReference_
}

var parserReference_ = &parserClass_{
	// Initialize the class constants.
	queueSize_: 16,
	stackSize_: 16,
	syntax_: col.Catalog[string, string](
		map[string]string{<SyntaxMap>
		},
	),
}
`,
}
