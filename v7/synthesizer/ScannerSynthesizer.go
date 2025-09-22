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
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
	not "github.com/craterdog/go-syntax-notation/v7"
)

// CLASS INTERFACE

// Access Function

func ScannerSynthesizerClass() ScannerSynthesizerClassLike {
	return scannerSynthesizerClass()
}

// Constructor Methods

func (c *scannerSynthesizerClass_) ScannerSynthesizer(
	syntax not.SyntaxLike,
) ScannerSynthesizerLike {
	var instance = &scannerSynthesizer_{
		// Initialize the instance attributes.
		analyzer_: ana.SyntaxAnalyzerClass().SyntaxAnalyzer(syntax),
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *scannerSynthesizer_) GetClass() ScannerSynthesizerClassLike {
	return scannerSynthesizerClass()
}

// TemplateDriven Methods

func (v *scannerSynthesizer_) CreateLegalNotice() string {
	var legalNotice = v.analyzer_.GetLegalNotice()
	return legalNotice
}

func (v *scannerSynthesizer_) CreateWarningMessage() string {
	var class = scannerSynthesizerClass()
	var warningMessage = class.warningMessage_
	return warningMessage
}

func (v *scannerSynthesizer_) CreateImportedPackages() string {
	var class = scannerSynthesizerClass()
	var importedPackages = class.importedPackages_
	return importedPackages
}

func (v *scannerSynthesizer_) CreateAccessFunction() string {
	var class = scannerSynthesizerClass()
	var accessFunction = class.accessFunction_
	return accessFunction
}

func (v *scannerSynthesizer_) CreateConstantMethods() string {
	var constantMethods string
	return constantMethods
}

func (v *scannerSynthesizer_) CreateConstructorMethods() string {
	var class = scannerSynthesizerClass()
	var constructorMethods = class.constructorMethods_
	return constructorMethods
}

func (v *scannerSynthesizer_) CreateFunctionMethods() string {
	var class = scannerSynthesizerClass()
	var functionMethods = class.functionMethods_
	return functionMethods
}

func (v *scannerSynthesizer_) CreatePrincipalMethods() string {
	var class = scannerSynthesizerClass()
	var principalMethods = class.principalMethods_
	return principalMethods
}

func (v *scannerSynthesizer_) CreateAttributeMethods() string {
	var attributeMethods string
	return attributeMethods
}

func (v *scannerSynthesizer_) CreateAspectInterfaces() string {
	var aspectInterfaces string
	return aspectInterfaces
}

func (v *scannerSynthesizer_) CreatePrivateMethods() string {
	var class = scannerSynthesizerClass()
	var privateMethods = class.privateMethods_
	var foundCases = v.createFoundCases()
	privateMethods = uti.ReplaceAll(
		privateMethods,
		"foundCases",
		foundCases,
	)
	return privateMethods
}

func (v *scannerSynthesizer_) CreateInstanceStructure() string {
	var class = scannerSynthesizerClass()
	var instanceStructure = class.instanceStructure_
	return instanceStructure
}

func (v *scannerSynthesizer_) CreateClassStructure() string {
	var class = scannerSynthesizerClass()
	var classStructure = class.classStructure_
	return classStructure
}

func (v *scannerSynthesizer_) CreateClass() string {
	var class = scannerSynthesizerClass()
	var classReference = class.classReference_
	var expressionIdentifiers = v.createExpressionIdentifiers()
	classReference = uti.ReplaceAll(
		classReference,
		"expressionIdentifiers",
		expressionIdentifiers,
	)
	var expressionMatchers = v.createExpressionMatchers()
	classReference = uti.ReplaceAll(
		classReference,
		"expressionMatchers",
		expressionMatchers,
	)
	var expressions = v.createExpressions()
	classReference = uti.ReplaceAll(
		classReference,
		"expressions",
		expressions,
	)
	return classReference
}

func (v *scannerSynthesizer_) PerformGlobalUpdates(
	existing string,
	generated string,
) string {
	return generated
}

// PROTECTED INTERFACE

// Private Methods

// NOTE:
// The found cases must be in the same order as the expressions declared in
// the syntax file.  But the expressions include patterns that are only used
// in other expressions and not visible as tokens.  And the token names have
// been sorted so we must pull out the token names from the catalog of
// expressions in their declared order.
func (v *scannerSynthesizer_) createFoundCases() string {
	var foundCases string
	var synthesizerClass = scannerSynthesizerClass()
	var expressions = v.analyzer_.GetExpressions()
	var expressionNames = fra.CatalogFromSequence[string, string](
		v.analyzer_.GetPatterns(),
	).GetKeys().GetIterator()
	for expressionNames.HasNext() {
		var expressionName = expressionNames.GetNext()
		if expressions.ContainsValue(expressionName) {
			var foundCase = synthesizerClass.foundCase_
			foundCase = uti.ReplaceAll(
				foundCase,
				"expressionName",
				expressionName,
			)
			foundCases += foundCase
		}
	}
	return foundCases
}

func (v *scannerSynthesizer_) createExpressions() string {
	var expressions string
	var set = fra.Set[string]()
	var patterns = v.analyzer_.GetPatterns().GetIterator()
	for patterns.HasNext() {
		var association = patterns.GetNext()
		var expressionName = association.GetKey()
		var expressionValue = association.GetValue()
		var class = scannerSynthesizerClass()
		var expression = class.expression_
		expression = uti.ReplaceAll(
			expression,
			"expressionName",
			expressionName,
		)
		expression = uti.ReplaceAll(
			expression,
			"expressionValue",
			expressionValue,
		)
		set.AddValue(expression)
	}
	var iterator = set.GetIterator()
	for iterator.HasNext() {
		var expression = iterator.GetNext()
		expressions += expression
	}
	return expressions
}

func (v *scannerSynthesizer_) createExpressionIdentifiers() string {
	var class = scannerSynthesizerClass()
	// Create the error expression identifier.
	var expressionIdentifiers = class.expressionIdentifier_
	expressionIdentifiers = uti.ReplaceAll(
		expressionIdentifiers,
		"expressionName",
		"error",
	)

	// Create the rest of the expression identifiers.
	var expressionNames = v.analyzer_.GetExpressions().GetIterator()
	for expressionNames.HasNext() {
		var expressionName = expressionNames.GetNext()
		var expressionIdentifier = class.expressionIdentifier_
		expressionIdentifier = uti.ReplaceAll(
			expressionIdentifier,
			"expressionName",
			expressionName,
		)
		expressionIdentifiers += expressionIdentifier
	}

	return expressionIdentifiers
}

func (v *scannerSynthesizer_) createExpressionMatchers() string {
	var expressionMatchers string
	var expressionNames = v.analyzer_.GetExpressions().GetIterator()
	for expressionNames.HasNext() {
		var expressionName = expressionNames.GetNext()
		var class = scannerSynthesizerClass()
		var expressionMatcher = class.expressionMatcher_
		expressionMatcher = uti.ReplaceAll(
			expressionMatcher,
			"expressionName",
			expressionName,
		)
		expressionMatchers += expressionMatcher
	}
	return expressionMatchers
}

// Instance Structure

type scannerSynthesizer_ struct {
	// Declare the instance attributes.
	analyzer_ ana.SyntaxAnalyzerLike
}

// Class Structure

type scannerSynthesizerClass_ struct {
	// Declare the class constants.
	warningMessage_       string
	importedPackages_     string
	accessFunction_       string
	constructorMethods_   string
	functionMethods_      string
	principalMethods_     string
	methodicalMethods_    string
	processExpression_    string
	processRule_          string
	privateMethods_       string
	foundCase_            string
	instanceStructure_    string
	classStructure_       string
	classReference_       string
	expressionIdentifier_ string
	expressionMatcher_    string
	expression_           string
}

// Class Reference

func scannerSynthesizerClass() *scannerSynthesizerClass_ {
	return scannerSynthesizerClassReference_
}

var scannerSynthesizerClassReference_ = &scannerSynthesizerClass_{
	// Initialize the class constants.
	warningMessage_: `
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│              This class file was automatically generated using:              │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
`,

	importedPackages_: `
	fmt "fmt"
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
	reg "regexp"
	sts "strings"
	uni "unicode"
`,

	accessFunction_: `
// Access Function

func ScannerClass() ScannerClassLike {
	return scannerClass()
}
`,

	constructorMethods_: `
// Constructor Methods

func (c *scannerClass_) Scanner(
	source string,
	tokens fra.QueueLike[TokenLike],
) ScannerLike {
	if uti.IsUndefined(source) {
		panic("The \"source\" attribute is required by this class.")
	}
	if uti.IsUndefined(tokens) {
		panic("The \"tokens\" attribute is required by this class.")
	}
	var instance = &scanner_{
		// Initialize the instance attributes.
		line_:     1,
		position_: 1,
		runes_:    []rune(source),
		tokens_:   tokens,
	}
	go instance.scanTokens() // Start scanning tokens in the background.
	return instance
}
`,

	functionMethods_: `
// Function Methods

func (c *scannerClass_) FormatToken(
	token TokenLike,
) string {
	var value = token.GetValue()
	value = fmt.Sprintf("%q", value)
	if len(value) > 40 {
		value = fmt.Sprintf("%.40q...", value)
	}
	return fmt.Sprintf(
		"Token [type: %s, line: %d, position: %d]: %s",
		c.tokens_.GetValue(token.GetType()),
		token.GetLine(),
		token.GetPosition(),
		value,
	)
}

func (c *scannerClass_) FormatType(
	tokenType TokenType,
) string {
	return c.tokens_.GetValue(tokenType)
}

func (c *scannerClass_) MatchesType(
	tokenValue string,
	tokenType TokenType,
) bool {
	var matcher = c.matchers_.GetValue(tokenType)
	var match = matcher.FindString(tokenValue)
	return uti.IsDefined(match)
}
`,

	principalMethods_: `
// Principal Methods

func (v *scanner_) GetClass() ScannerClassLike {
	return scannerClass()
}
`,

	methodicalMethods_: `
// Methodical Methods
<ProcessExpressions><ProcessRules>`,

	processExpression_: `
func (v *scanner_) Process<~ExpressionName>(
	<expressionName_> string,
) {
	v.validateExpression(<expressionName_>, <~ExpressionName>Expression)
}
`,

	processRule_: `
func (v *scanner_) Preprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *scanner_) Postprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *scanner_) Process<~RuleName>Slot(
	<ruleName_> ast.<~RuleName>Like,
	slot_ uint,
) {
	// TBD - Add any validation checks.
}
`,

	privateMethods_: `
// Private Methods

func (v *scanner_) emitToken(
	tokenType TokenType,
) {
	var value = string(v.runes_[v.first_:v.next_])
	switch value {
	case "\x00":
		value = "<NULL>"
	case "\a":
		value = "<BELL>"
	case "\b":
		value = "<BKSP>"
	case "\t":
		value = "<HTAB>"
	case "\f":
		value = "<FMFD>"
	case "\r":
		value = "<CRTN>"
	case "\v":
		value = "<VTAB>"
	}
	var token = TokenClass().Token(v.line_, v.position_, tokenType, value)
	//fmt.Println(ScannerClass().FormatToken(token)) // Uncomment when debugging.
	v.tokens_.AddValue(token) // This will block if the queue is full.
}

func (v *scanner_) foundError() {
	v.next_++
	v.emitToken(ErrorToken)
}

func (v *scanner_) foundToken(
	tokenType TokenType,
) bool {
	// Attempt to match the specified token type.
	var class = scannerClass()
	var matcher = class.matchers_.GetValue(tokenType)
	var text = string(v.runes_[v.next_:])
	var match = matcher.FindString(text)
	if uti.IsUndefined(match) {
		return false
	}

	// Check for an exact delimiter match which takes precedence.
	if tokenType != DelimiterToken {
		matcher = class.matchers_.GetValue(DelimiterToken)
		var delimiter = matcher.FindString(match)
		if uti.IsDefined(delimiter) && delimiter == match {
			// This is a delimiter instead so ignore it.
			return false
		}
	}

	// Check for partial identifier matches.
	var token = []rune(match)
	var size = uti.ArraySize(token)
	var previous = token[size-1]
	if uti.ArraySize(v.runes_) > v.next_+size {
		var next = v.runes_[v.next_+size]
		if (uni.IsLetter(previous) || uni.IsNumber(previous)) &&
			(uni.IsLetter(next) || uni.IsNumber(next) || next == '_') {
			return false
		}
	}

	// Found the requested token type.
	v.next_ += size
	v.emitToken(tokenType)
	var count = uint(sts.Count(match, "\n"))
	if count > 0 {
		v.line_ += count
		v.position_ = v.indexOfLastEol(token)
	} else {
		v.position_ += v.next_ - v.first_
	}
	v.first_ = v.next_
	return true
}

func (v *scanner_) indexOfLastEol(
	runes []rune,
) uint {
	var size = uti.ArraySize(runes)
	for index := size; index > 0; index-- {
		if runes[index-1] == '\n' {
			return size - index + 1
		}
	}
	return 0
}

func (v *scanner_) scanTokens() {
loop:
	for v.next_ < uti.ArraySize(v.runes_) {
		switch {
		// Find the next token type.<FoundCases>
		default:
			v.foundError()
			break loop
		}
	}
	v.tokens_.CloseChannel()
}
`,

	foundCase_: `
		case v.foundToken(<~ExpressionName>Token):`,

	instanceStructure_: `
// Instance Structure

type scanner_ struct {
	// Declare the instance attributes.
	first_    uint // A zero based index of the first possible rune in the next token.
	next_     uint // A zero based index of the next possible rune in the next token.
	line_     uint // The line number in the source string of the next rune.
	position_ uint // The position in the current line of the next rune.
	runes_    []rune
	tokens_   fra.QueueLike[TokenLike]
}
`,

	classStructure_: `
// Class Structure

type scannerClass_ struct {
	// Declare the class constants.
	tokens_   fra.CatalogLike[TokenType, string]
	matchers_ fra.CatalogLike[TokenType, *reg.Regexp]
}
`,

	classReference_: `
// Class Reference

func scannerClass() *scannerClass_ {
	return scannerClassReference_
}

var scannerClassReference_ = &scannerClass_{
	// Initialize the class constants.
	tokens_: fra.CatalogFromMap[TokenType, string](
		map[TokenType]string{
			// Define token identifiers for each type of expression.<ExpressionIdentifiers>
		},
	),
	matchers_: fra.CatalogFromMap[TokenType, *reg.Regexp](
		map[TokenType]*reg.Regexp{
			// Define pattern matchers for each type of expression.<ExpressionMatchers>
		},
	),
}

// Private Constants

// NOTE:
// These private constants define the regular expression sub-patterns that make
// up the intrinsic types and expression types.  Unfortunately there is no way
// to make them private to the scanner class since they must be TRUE Go constants
// to be used in this way.  We append an underscore to each name to lessen the
// chance of a name collision with other private Go class constants in this
// package.
const (
	// Define the regular expressions for each intrinsic type.
	any_     = "." // This does NOT include newline characters.
	control_ = "\\p{Cc}"
	digit_   = "\\p{Nd}"
	eol_     = "\\r?\\n"
	lower_   = "\\p{Ll}"
	upper_   = "\\p{Lu}"

	// Define the regular expressions for each expression type.<Expressions>
)
`,

	expressionIdentifier_: `
			<~ExpressionName>Token: "<~expressionName>",`,

	expressionMatcher_: `
			<~ExpressionName>Token: reg.MustCompile("^" + <~expressionName>_),`,

	expression_: `
	<~expressionName>_ = <ExpressionValue>`,
}
