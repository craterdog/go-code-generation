/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies.  All Rights Reserved.    .
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
	ana "github.com/craterdog/go-code-generation/v7/analyzer"
	uti "github.com/craterdog/go-missing-utilities/v7"
	not "github.com/craterdog/go-syntax-notation/v7"
)

// CLASS INTERFACE

// Access Function

func ParserSynthesizerClass() ParserSynthesizerClassLike {
	return parserSynthesizerClass()
}

// Constructor Methods

func (c *parserSynthesizerClass_) ParserSynthesizer(
	syntax not.SyntaxLike,
) ParserSynthesizerLike {
	var instance = &parserSynthesizer_{
		// Initialize the instance attributes.
		analyzer_: ana.SyntaxAnalyzerClass().SyntaxAnalyzer(syntax),
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *parserSynthesizer_) GetClass() ParserSynthesizerClassLike {
	return parserSynthesizerClass()
}

// TemplateDriven Methods

func (v *parserSynthesizer_) CreateLegalNotice() string {
	var legalNotice = v.analyzer_.GetLegalNotice()
	return legalNotice
}

func (v *parserSynthesizer_) CreateWarningMessage() string {
	var class = parserSynthesizerClass()
	var warningMessage = class.warningMessage_
	return warningMessage
}

func (v *parserSynthesizer_) CreateImportedPackages() string {
	var class = parserSynthesizerClass()
	var importedPackages = class.importedPackages_
	return importedPackages
}

func (v *parserSynthesizer_) CreateAccessFunction() string {
	var class = parserSynthesizerClass()
	var accessFunction = class.accessFunction_
	return accessFunction
}

func (v *parserSynthesizer_) CreateConstantMethods() string {
	var constantMethods string
	return constantMethods
}

func (v *parserSynthesizer_) CreateConstructorMethods() string {
	var class = parserSynthesizerClass()
	var constructorMethods = class.constructorMethods_
	return constructorMethods
}

func (v *parserSynthesizer_) CreateFunctionMethods() string {
	var functionMethods string
	return functionMethods
}

func (v *parserSynthesizer_) CreatePrincipalMethods() string {
	var class = parserSynthesizerClass()
	var principalMethods = class.principalMethods_
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
	var class = parserSynthesizerClass()
	var privateMethods = class.privateMethods_
	privateMethods = uti.ReplaceAll(
		privateMethods,
		"parseMethods",
		parseMethods,
	)
	return privateMethods
}

func (v *parserSynthesizer_) CreateInstanceStructure() string {
	var class = parserSynthesizerClass()
	var instanceStructure = class.instanceStructure_
	return instanceStructure
}

func (v *parserSynthesizer_) CreateClassStructure() string {
	var class = parserSynthesizerClass()
	var classStructure = class.classStructure_
	return classStructure
}

func (v *parserSynthesizer_) CreateClass() string {
	var class = parserSynthesizerClass()
	var classReference = class.classReference_
	return classReference
}

func (v *parserSynthesizer_) PerformGlobalUpdates(
	existing string,
	generated string,
) string {
	var syntaxMap = v.analyzer_.GetSyntaxMap()
	generated = uti.ReplaceAll(
		generated,
		"syntaxMap",
		syntaxMap,
	)
	var syntaxName = v.analyzer_.GetSyntaxName()
	generated = uti.ReplaceAll(
		generated,
		"syntaxName",
		syntaxName,
	)
	return generated
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

// NOTE:
// This method is recursive.  It walks the AST to determine if a rule definition
// contains any inline rule definitions.  If so, the look-ahead token buffer
// must be turned off.
func (v *parserSynthesizer_) containsTermSequence(
	ruleName string,
) bool {
	var definitions = v.analyzer_.GetDefinitions()
	switch definitions.GetValue(ruleName).GetAny().(type) {
	case not.LiteralAlternativesLike:
		// The rule name refers to a multi-literal rule.
	case not.TokenAlternativesLike:
		// The rule name refers to a multi-expression rule.
	case not.RuleAlternativesLike:
		// The rule name refers to a multi-rule rule.
		var options = v.analyzer_.GetRuleNames(ruleName).GetIterator()
		for options.HasNext() {
			var option = options.GetNext()
			var uppercase = option.GetUppercase()
			if v.containsTermSequence(uppercase) {
				// Once true, it can never go back to false.
				return true
			}
		}
	case not.TermSequenceLike:
		// The rule name refers to an inline rule.
		return true
	}
	return false
}

func (v *parserSynthesizer_) createRuleTerm(
	term not.RuleTermLike,
) string {
	var implementation string
	var cardinality = term.GetOptionalCardinality()
	var actual = term.GetComponent().GetAny().(string)
	switch {
	case not.MatchesType(actual, not.LiteralToken):
		implementation = v.createLiteral(actual, cardinality)
	case not.MatchesType(actual, not.LowercaseToken):
		implementation = v.createExpression(actual, cardinality)
	case not.MatchesType(actual, not.UppercaseToken):
		implementation = v.createRule(actual, cardinality)
	}
	return implementation
}

func (v *parserSynthesizer_) createTermSequence(
	ruleName string,
) string {
	var class = parserSynthesizerClass()
	var implementation = class.declareTokens_
	var terms = v.analyzer_.GetRuleTerms(ruleName).GetIterator()
	var variables = v.analyzer_.GetVariables(ruleName).GetIterator()
	for terms.HasNext() {
		var term = terms.GetNext()
		var parseRuleTerm = v.createRuleTerm(term)
		var variableName = variables.GetNext()
		parseRuleTerm = uti.ReplaceAll(
			parseRuleTerm,
			"variableName",
			variableName,
		)
		implementation += parseRuleTerm
	}
	var found = class.found_
	var arguments = v.createArguments(ruleName)
	implementation += uti.ReplaceAll(
		found,
		"arguments",
		arguments,
	)
	return implementation
}

func (v *parserSynthesizer_) createLiteralAlternatives(
	ruleName string,
) string {
	var class = parserSynthesizerClass()
	var implementation = class.declareDelimiter_
	var literalValues = v.analyzer_.GetLiteralValues(ruleName).GetIterator()
	for literalValues.HasNext() {
		var literal = literalValues.GetNext().GetLiteral()
		var option = class.literalValue_
		implementation += uti.ReplaceAll(
			option,
			"literal",
			literal,
		)
	}
	implementation += class.defaultOption_
	return implementation
}

func (v *parserSynthesizer_) createTokenAlternatives(
	ruleName string,
) string {
	var implementation string
	var class = parserSynthesizerClass()
	var tokenNames = v.analyzer_.GetTokenNames(ruleName).GetIterator()
	for tokenNames.HasNext() {
		var lowercase = tokenNames.GetNext().GetLowercase()
		var parseExpressionName = class.expressionName_
		implementation += uti.ReplaceAll(
			parseExpressionName,
			"lowercase",
			lowercase,
		)
	}
	implementation += class.defaultOption_
	return implementation
}

func (v *parserSynthesizer_) createRuleAlternatives(
	ruleName string,
) string {
	var implementation string
	var class = parserSynthesizerClass()
	var ruleNames = v.analyzer_.GetRuleNames(ruleName).GetIterator()
	for ruleNames.HasNext() {
		var uppercase = ruleNames.GetNext().GetUppercase()
		var parseRuleName = class.ruleName_
		implementation += uti.ReplaceAll(
			parseRuleName,
			"uppercase",
			uppercase,
		)
	}
	implementation += class.defaultOption_
	return implementation
}

func (v *parserSynthesizer_) createParseMethod(
	ruleName string,
) string {
	var methodImplementation string
	var definitions = v.analyzer_.GetDefinitions()
	var definition = definitions.GetValue(ruleName)
	switch definition.GetAny().(type) {
	case not.LiteralAlternativesLike:
		methodImplementation = v.createLiteralAlternatives(ruleName)
	case not.TokenAlternativesLike:
		methodImplementation = v.createTokenAlternatives(ruleName)
	case not.RuleAlternativesLike:
		methodImplementation = v.createRuleAlternatives(ruleName)
	case not.TermSequenceLike:
		methodImplementation = v.createTermSequence(ruleName)
	}
	var class = parserSynthesizerClass()
	var parseMethod = class.parseMethod_
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
	var ruleNames = v.analyzer_.GetRules().GetIterator()
	for ruleNames.HasNext() {
		var ruleName = ruleNames.GetNext()
		parseMethods += v.createParseMethod(ruleName)
	}
	return parseMethods
}

func (v *parserSynthesizer_) createRule(
	uppercase string,
	cardinality not.CardinalityLike,
) string {
	var class = parserSynthesizerClass()
	var required = class.requiredNestedExpression_
	var optional = class.optionalNestedExpression_
	var repeated = class.repeatedNestedExpression_
	if v.containsTermSequence(uppercase) {
		required = class.requiredRule_
		optional = class.optionalRule_
		repeated = class.repeatedRule_
	}
	var implementation = v.createCardinality(
		cardinality,
		optional,
		required,
		repeated,
	)
	implementation = uti.ReplaceAll(
		implementation,
		"uppercase",
		uppercase,
	)
	return implementation
}

func (v *parserSynthesizer_) createLiteral(
	literal string,
	cardinality not.CardinalityLike,
) string {
	var class = parserSynthesizerClass()
	var optional = class.optionalLiteral_
	var required = class.requiredLiteral_
	var repeated = class.repeatedLiteral_
	var implementation = v.createCardinality(
		cardinality,
		optional,
		required,
		repeated,
	)
	implementation = uti.ReplaceAll(
		implementation,
		"literal",
		literal,
	)
	return implementation
}

func (v *parserSynthesizer_) createExpression(
	lowercase string,
	cardinality not.CardinalityLike,
) string {
	var class = parserSynthesizerClass()
	var optional = class.optionalExpression_
	var required = class.requiredExpression_
	var repeated = class.repeatedExpression_
	var implementation = v.createCardinality(
		cardinality,
		optional,
		required,
		repeated,
	)
	implementation = uti.ReplaceAll(
		implementation,
		"lowercase",
		lowercase,
	)
	return implementation
}

// Instance Structure

type parserSynthesizer_ struct {
	// Declare the instance attributes.
	analyzer_ ana.SyntaxAnalyzerLike
}

// Class Structure

type parserSynthesizerClass_ struct {
	// Declare the class constants.
	warningMessage_           string
	importedPackages_         string
	accessFunction_           string
	constructorMethods_       string
	parseMethod_              string
	declareDelimiter_         string
	declareTokens_            string
	optionalLiteral_          string
	requiredLiteral_          string
	repeatedLiteral_          string
	optionalExpression_       string
	requiredExpression_       string
	repeatedExpression_       string
	optionalNestedExpression_ string
	requiredNestedExpression_ string
	repeatedNestedExpression_ string
	optionalRule_             string
	requiredRule_             string
	repeatedRule_             string
	found_                    string
	literalValue_             string
	expressionName_           string
	ruleName_                 string
	defaultOption_            string
	principalMethods_         string
	privateMethods_           string
	instanceStructure_        string
	classStructure_           string
	classReference_           string
}

// Class Reference

func parserSynthesizerClass() *parserSynthesizerClass_ {
	return parserSynthesizerClassReference_
}

var parserSynthesizerClassReference_ = &parserSynthesizerClass_{
	// Initialize the class constants.
	warningMessage_: `
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│                 This class file was automatically generated.                 │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
`,

	importedPackages_: `
	fmt "fmt"
	ast "<ModuleName>/ast"
	col "github.com/craterdog/go-collection-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
	mat "math"
	sts "strings"
`,

	accessFunction_: `
// Access Function

func ParserClass() ParserClassLike {
	return parserClass()
}
`,

	constructorMethods_: `
// Constructor Methods

func (c *parserClass_) Parser() ParserLike {
	var instance = &parser_{
		// Initialize the instance attributes.
	}
	return instance
}
`,

	principalMethods_: `
// Principal Methods

func (v *parser_) GetClass() ParserClassLike {
	return parserClass()
}

func (v *parser_) ParseSource(
	source string,
) ast.<~SyntaxName>Like {
	v.source_ = sts.ReplaceAll(source, "\t", "    ")
	v.tokens_ = col.Queue[TokenLike]()
	v.next_ = col.Stack[TokenLike]()

	// The scanner runs in a separate Go routine.
	ScannerClass().Scanner(v.source_, v.tokens_)

	// Attempt to parse the <~syntaxName>.
	var <syntaxName_>, token, ok = v.parse<~SyntaxName>()
	if !ok || !v.tokens_.IsEmpty() {
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
	tokens col.Sequential[TokenLike],
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
	tokens col.Sequential[TokenLike],
) {
}
`,

	parseMethod_: `
func (v *parser_) parse<~RuleName>() (
	<ruleName_> ast.<~RuleName>Like,
	token TokenLike,
	ok bool,
) {<MethodImplementation>}
`,

	declareDelimiter_: `
	var delimiter string
`,

	declareTokens_: `
	var tokens = col.List[TokenLike]()
`,

	optionalLiteral_: `
	// Attempt to parse an optional <literal> literal.
	var <variableName_> string
	<variableName_>, token, ok = v.parseDelimiter(<literal>)
	if ok {
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	} else {
		<variableName_> = "" // Reset this to undefined.
	}
`,

	requiredLiteral_: `
	// Attempt to parse a single <literal> literal.
	var <variableName_> string
	<variableName_>, token, ok = v.parseDelimiter(<literal>)
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

	repeatedLiteral_: `
	// Attempt to parse multiple <literal> literals.
	var <variableName_> = col.List[string]()
<~variableName>Loop:
	for count_ := 0; count_ < <last>; count_++ {
		var literal string
		literal, token, ok = v.parseDelimiter(<literal>)
		if !ok {
			switch {
			case count_ >= <first>:
				break <~variableName>Loop
			case uti.IsDefined(tokens):
				// This is not multiple <literal> literals.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$<~RuleName>", token)
				message += "<first> or more <literal> literals are required."
				panic(message)
			}
		}
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
		<variableName_>.AppendValue(literal)
	}
`,

	optionalExpression_: `
	// Attempt to parse an optional <~lowercase> token.
	var <variableName_> string
	<variableName_>, token, ok = v.parseToken(<~Lowercase>Token)
	if ok {
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	} else {
		<variableName_> = "" // Reset this to undefined.
	}
`,

	requiredExpression_: `
	// Attempt to parse a single <~lowercase> token.
	var <variableName_> string
	<variableName_>, token, ok = v.parseToken(<~Lowercase>Token)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single <~lowercase> token.
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

	repeatedExpression_: `
	// Attempt to parse multiple <~lowercase> tokens.
	var <variableName_> = col.List[string]()
<~variableName>Loop:
	for count_ := 0; count_ < <last>; count_++ {
		var <lowercase_> string
		<lowercase_>, token, ok = v.parseToken(<~Lowercase>Token)
		if !ok {
			switch {
			case count_ >= <first>:
				break <~variableName>Loop
			case uti.IsDefined(tokens):
				// This is not multiple <~lowercase> tokens.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$<~RuleName>", token)
				message += "<first> or more <~lowercase> tokens are required."
				panic(message)
			}
		}
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
		<variableName_>.AppendValue(<lowercase_>)
	}
`,

	optionalNestedExpression_: `
	// Attempt to parse an optional <~Uppercase> rule.
	var <variableName_> ast.<~Uppercase>Like
	<variableName_>, token, ok = v.parse<~Uppercase>()
	if ok {
		// Found a multiexpression token.
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	}
`,

	requiredNestedExpression_: `
	// Attempt to parse a single <~Uppercase> rule.
	var <variableName_> ast.<~Uppercase>Like
	<variableName_>, token, ok = v.parse<~Uppercase>()
	switch {
	case ok:
		// Found a multiexpression token.
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	case uti.IsDefined(tokens):
		// This is not a single <~Uppercase> rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$<~RuleName>", token)
		panic(message)
	}
`,

	repeatedNestedExpression_: `
	// Attempt to parse multiple <~Uppercase> rules.
	var <variableName_> = col.List[ast.<~Uppercase>Like]()
<~variableName>Loop:
	for count_ := 0; count_ < <last>; count_++ {
		var <uppercase_> ast.<~Uppercase>Like
		<uppercase_>, token, ok = v.parse<~Uppercase>()
		if !ok {
			switch {
			case count_ >= <first>:
				break <~variableName>Loop
			case uti.IsDefined(tokens):
				// This is not multiple <~Uppercase> rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$<~RuleName>", token)
				message += "<first> or more <~Uppercase> rules are required."
				panic(message)
			}
		}
		// Found a multiexpression token.
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
		<variableName_>.AppendValue(<uppercase_>)
	}
`,

	optionalRule_: `
	// Attempt to parse an optional <~Uppercase> rule.
	var <variableName_> ast.<~Uppercase>Like
	<variableName_>, _, ok = v.parse<~Uppercase>()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}
`,

	requiredRule_: `
	// Attempt to parse a single <~Uppercase> rule.
	var <variableName_> ast.<~Uppercase>Like
	<variableName_>, token, ok = v.parse<~Uppercase>()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single <~Uppercase> rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$<~RuleName>", token)
		panic(message)
	}
`,

	repeatedRule_: `
	// Attempt to parse multiple <~Uppercase> rules.
	var <variableName_> = col.List[ast.<~Uppercase>Like]()
<~variableName>Loop:
	for count_ := 0; count_ < <last>; count_++ {
		var <uppercase_> ast.<~Uppercase>Like
		<uppercase_>, token, ok = v.parse<~Uppercase>()
		if !ok {
			switch {
			case count_ >= <first>:
				break <~variableName>Loop
			case uti.IsDefined(tokens):
				// This is not multiple <~Uppercase> rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$<~RuleName>", token)
				message += "<first> or more <~Uppercase> rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		<variableName_>.AppendValue(<uppercase_>)
	}
`,

	found_: `
	// Found a single <~RuleName> rule.
	ok = true
	v.remove(tokens)
	<ruleName_> = ast.<~RuleName>Class().<~RuleName>(<Arguments>)
	return
`,

	literalValue_: `
	// Attempt to parse a single <literal> delimiter.
	delimiter, token, ok = v.parseDelimiter(<literal>)
	if ok {
		// Found a single <literal> delimiter.
		<ruleName_> = ast.<~RuleName>Class().<~RuleName>(delimiter)
		return
	}
`,

	expressionName_: `
	// Attempt to parse a single <~lowercase> <~RuleName>.
	var <lowercase_> string
	<lowercase_>, token, ok = v.parseToken(<~Lowercase>Token)
	if ok {
		// Found a single <~lowercase> <~RuleName>.
		<ruleName_> = ast.<~RuleName>Class().<~RuleName>(<lowercase_>)
		return
	}
`,

	ruleName_: `
	// Attempt to parse a single <~Uppercase> <~RuleName>.
	var <uppercase_> ast.<~Uppercase>Like
	<uppercase_>, token, ok = v.parse<~Uppercase>()
	if ok {
		// Found a single <~Uppercase> <~RuleName>.
		<ruleName_> = ast.<~RuleName>Class().<~RuleName>(<uppercase_>)
		return
	}
`,

	defaultOption_: `
	// This is not a single <~RuleName> rule.
	return
`,

	instanceStructure_: `
// Instance Structure

type parser_ struct {
	// Declare the instance attributes.
	source_ string                   // The original source code.
	tokens_ col.QueueLike[TokenLike] // A queue of unread tokens from the scanner.
	next_   col.StackLike[TokenLike] // A stack of read, but unprocessed tokens.
}
`,

	classStructure_: `
// Class Structure

type parserClass_ struct {
	// Declare the class constants.
	syntax_ col.CatalogLike[string, string]
}
`,

	classReference_: `
// Class Reference

func parserClass() *parserClass_ {
	return parserClassReference_
}

var parserClassReference_ = &parserClass_{
	// Initialize the class constants.
	syntax_: col.CatalogFromMap[string, string](
		map[string]string{<SyntaxMap>
		},
	),
}
`,
}
