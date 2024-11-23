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
)

// CLASS INTERFACE

// Access Function

func ScannerSynthesizer() ScannerSynthesizerClassLike {
	return scannerSynthesizerReference()
}

// Constructor Methods

func (c *scannerSynthesizerClass_) Make(
	syntax not.SyntaxLike,
) ScannerSynthesizerLike {
	var instance = &scannerSynthesizer_{
		// Initialize the instance attributes.
		analyzer_: ana.SyntaxAnalyzer().Make(syntax),
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *scannerSynthesizer_) GetClass() ScannerSynthesizerClassLike {
	return scannerSynthesizerReference()
}

// TemplateDriven Methods

func (v *scannerSynthesizer_) CreateLegalNotice() string {
	var legalNotice = v.analyzer_.GetLegalNotice()
	return legalNotice
}

func (v *scannerSynthesizer_) CreateAccessFunction() string {
	var accessFunction = scannerSynthesizerReference().accessFunction_
	return accessFunction
}

func (v *scannerSynthesizer_) CreateConstantMethods() string {
	var constantMethods string
	return constantMethods
}

func (v *scannerSynthesizer_) CreateConstructorMethods() string {
	var constructorMethods = scannerSynthesizerReference().constructorMethods_
	return constructorMethods
}

func (v *scannerSynthesizer_) CreateFunctionMethods() string {
	var functionMethods = scannerSynthesizerReference().functionMethods_
	return functionMethods
}

func (v *scannerSynthesizer_) CreatePrincipalMethods() string {
	var principalMethods = scannerSynthesizerReference().principalMethods_
	return principalMethods
}

func (v *scannerSynthesizer_) CreateAttributeMethods() string {
	var attributeMethods string
	return attributeMethods
}

func (v *scannerSynthesizer_) CreateAspectMethods() string {
	var aspectMethods string
	return aspectMethods
}

func (v *scannerSynthesizer_) CreatePrivateMethods() string {
	var privateMethods = scannerSynthesizerReference().privateMethods_
	var foundCases = v.createFoundCases()
	privateMethods = uti.ReplaceAll(
		privateMethods,
		"foundCases",
		foundCases,
	)
	return privateMethods
}

func (v *scannerSynthesizer_) CreateInstanceStructure() string {
	var instanceStructure = scannerSynthesizerReference().instanceStructure_
	return instanceStructure
}

func (v *scannerSynthesizer_) CreateClassStructure() string {
	var classStructure = scannerSynthesizerReference().classStructure_
	return classStructure
}

func (v *scannerSynthesizer_) CreateClassReference() string {
	var classReference = scannerSynthesizerReference().classReference_
	var tokenIdentifiers = v.createTokenIdentifiers()
	classReference = uti.ReplaceAll(
		classReference,
		"tokenIdentifiers",
		tokenIdentifiers,
	)
	var tokenMatchers = v.createTokenMatchers()
	classReference = uti.ReplaceAll(
		classReference,
		"tokenMatchers",
		tokenMatchers,
	)
	var regularExpressions = v.createRegularExpressions()
	classReference = uti.ReplaceAll(
		classReference,
		"regularExpressions",
		regularExpressions,
	)
	return classReference
}

func (v *scannerSynthesizer_) PerformGlobalUpdates(
	source string,
) string {
	var classImports = scannerSynthesizerReference().classImports_
	source = uti.ReplaceAll(
		source,
		"classImports",
		classImports,
	)
	return source
}

// PROTECTED INTERFACE

// Private Methods

func (v *scannerSynthesizer_) createRegularExpressions() string {
	var regularExpressions string
	var expressions = v.analyzer_.GetExpressions().GetIterator()
	for expressions.HasNext() {
		var association = expressions.GetNext()
		var expressionName = association.GetKey()
		var expressionValue = association.GetValue()
		var regularExpression = scannerSynthesizerReference().regularExpression_
		regularExpression = uti.ReplaceAll(
			regularExpression,
			"expressionName",
			expressionName,
		)
		regularExpression = uti.ReplaceAll(
			regularExpression,
			"expressionValue",
			expressionValue,
		)
		regularExpressions += regularExpression
	}
	return regularExpressions
}

func (v *scannerSynthesizer_) createFoundCases() string {
	var foundCases string
	var tokenNames = v.analyzer_.GetTokenNames().GetIterator()
	for tokenNames.HasNext() {
		var tokenName = tokenNames.GetNext()
		var foundCase = scannerSynthesizerReference().foundCase_
		foundCase = uti.ReplaceAll(
			foundCase,
			"tokenName",
			tokenName,
		)
		foundCases += foundCase
	}
	return foundCases
}

func (v *scannerSynthesizer_) createTokenIdentifiers() string {
	var tokenIdentifiers = `ErrorToken: "error",`
	var tokenNames = v.analyzer_.GetTokenNames().GetIterator()
	for tokenNames.HasNext() {
		var tokenName = tokenNames.GetNext()
		var tokenIdentifier = scannerSynthesizerReference().tokenIdentifier_
		tokenIdentifier = uti.ReplaceAll(
			tokenIdentifier,
			"tokenName",
			tokenName,
		)
		tokenIdentifiers += tokenIdentifier
	}
	return tokenIdentifiers
}

func (v *scannerSynthesizer_) createTokenMatchers() string {
	var tokenMatchers string
	var tokenNames = v.analyzer_.GetTokenNames().GetIterator()
	for tokenNames.HasNext() {
		var tokenName = tokenNames.GetNext()
		var tokenMatcher = scannerSynthesizerReference().tokenMatcher_
		tokenMatcher = uti.ReplaceAll(
			tokenMatcher,
			"tokenName",
			tokenName,
		)
		tokenMatchers += tokenMatcher
	}
	return tokenMatchers
}

// Instance Structure

type scannerSynthesizer_ struct {
	// Declare the instance attributes.
	analyzer_ ana.SyntaxAnalyzerLike
}

// Class Structure

type scannerSynthesizerClass_ struct {
	// Declare the class constants.
	classImports_        string
	accessFunction_      string
	constructorMethods_  string
	functionMethods_     string
	principalMethods_    string
	methodicalMethods_   string
	processToken_        string
	processIndexedToken_ string
	processRule_         string
	processIndexedRule_  string
	privateMethods_      string
	foundCase_           string
	instanceStructure_   string
	classStructure_      string
	classReference_      string
	tokenIdentifier_     string
	tokenMatcher_        string
	regularExpression_   string
}

// Class Reference

func scannerSynthesizerReference() *scannerSynthesizerClass_ {
	return scannerSynthesizerReference_
}

var scannerSynthesizerReference_ = &scannerSynthesizerClass_{
	// Initialize the class constants.
	classImports_: `
import (
	fmt "fmt"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
	reg "regexp"
	sts "strings"
	uni "unicode"
)
`,

	accessFunction_: `
// Access Function

func Scanner() ScannerClassLike {
	return scannerReference()
}
`,

	constructorMethods_: `
// Constructor Methods

func (c *scannerClass_) Make(
	source string,
	tokens abs.QueueLike[TokenLike],
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
		c.tokens_[token.GetType()],
		token.GetLine(),
		token.GetPosition(),
		value,
	)
}

func (c *scannerClass_) FormatType(
	tokenType TokenType,
) string {
	return c.tokens_[tokenType]
}

func (c *scannerClass_) MatchesType(
	tokenValue string,
	tokenType TokenType,
) bool {
	var matcher = c.matchers_[tokenType]
	var match = matcher.FindString(tokenValue)
	return uti.IsDefined(match)
}
`,

	principalMethods_: `
// Principal Methods

func (v *scanner_) GetClass() ScannerClassLike {
	return scannerReference()
}
`,

	methodicalMethods_: `
// Methodical Methods
<ProcessTokens><ProcessRules>`,

	processToken_: `
func (v *scanner_) Process<~TokenName>(
	<tokenName_> string,
) {
	v.validateToken(<tokenName_>, <~TokenName>Token)
}
`,

	processIndexedToken_: `
func (v *scanner_) Process<~TokenName>(
	<tokenName_> string,
	index uint,
	size uint,
) {
	v.validateToken(<tokenName_>, <~TokenName>Token)
}
`,

	processRule_: `
func (v *scanner_) Preprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
) {
	// TBD - Add any validation checks.
}

func (v *scanner_) Process<~RuleName>Slot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *scanner_) Postprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
) {
	// TBD - Add any validation checks.
}
`,

	processIndexedRule_: `
func (v *scanner_) Preprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *scanner_) Process<~RuleName>Slot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *scanner_) Postprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}
`,

	privateMethods_: `
// Private Methods

func (v *scanner_) emitToken(tokenType TokenType) {
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
	var token = Token().Make(v.line_, v.position_, tokenType, value)
	//fmt.Println(Scanner().FormatToken(token)) // Uncomment when debugging.
	v.tokens_.AddValue(token) // This will block if the queue is full.
}

func (v *scanner_) foundError() {
	v.next_++
	v.emitToken(ErrorToken)
}

func (v *scanner_) foundToken(tokenType TokenType) bool {
	// Attempt to match the specified token type.
	var text = string(v.runes_[v.next_:])
	var matcher = scannerReference().matchers_[tokenType]
	var match = matcher.FindString(text)
	if uti.IsUndefined(match) {
		return false
	}

	// Check for false delimiter matches.
	var token = []rune(match)
	var length = uint(len(token))
	var previous = token[length-1]
	if tokenType == DelimiterToken && uint(len(v.runes_)) > v.next_+length {
		var next = v.runes_[v.next_+length]
		if (uni.IsLetter(previous) || uni.IsNumber(previous)) &&
			(uni.IsLetter(next) || uni.IsNumber(next) || next == '_') {
			return false
		}
	}

	// Found the requested token type.
	v.next_ += length
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

func (v *scanner_) indexOfLastEol(runes []rune) uint {
	var length = uint(len(runes))
	for index := length; index > 0; index-- {
		if runes[index-1] == '\n' {
			return length - index + 1
		}
	}
	return 0
}

func (v *scanner_) scanTokens() {
loop:
	for v.next_ < uint(len(v.runes_)) {
		switch {
		// Find the next token type.<FoundCases>
		default:
			v.foundError()
			break loop
		}
	}
	v.tokens_.CloseQueue()
}
`,

	foundCase_: `
		case v.foundToken(<~TokenName>Token):`,

	instanceStructure_: `
// Instance Structure

type scanner_ struct {
	// Declare the instance attributes.
	first_    uint // A zero based index of the first possible rune in the next token.
	next_     uint // A zero based index of the next possible rune in the next token.
	line_     uint // The line number in the source string of the next rune.
	position_ uint // The position in the current line of the next rune.
	runes_    []rune
	tokens_   abs.QueueLike[TokenLike]
}
`,

	classStructure_: `
// Class Structure

type scannerClass_ struct {
	// Declare the class constants.
	tokens_   map[TokenType]string
	matchers_ map[TokenType]*reg.Regexp
}
`,

	classReference_: `
// Class Reference

func scannerReference() *scannerClass_ {
	return scannerReference_
}

var scannerReference_ = &scannerClass_{
	// Initialize the class constants.
	tokens_: map[TokenType]string{
		// Define identifiers for each type of token.
		<TokenIdentifiers>
	},
	matchers_: map[TokenType]*reg.Regexp{
		// Define pattern matchers for each type of token.<TokenMatchers>
	},
}

// Private Constants

/*
NOTE:
These private constants define the regular expression sub-patterns that make up
the intrinsic types and token types.  Unfortunately there is no way to make them
private to the scanner class since they must be TRUE Go constants to be used in
this way.  We append an underscore to each name to lessen the chance of a name
collision with other private Go class constants in this package.
*/
const (
	// Define the regular expression patterns for each intrinsic type.
	any_     = "." // This does NOT include newline characters.
	control_ = "\\p{Cc}"
	digit_   = "\\p{Nd}"
	eol_     = "\\r?\\n"
	lower_   = "\\p{Ll}"
	upper_   = "\\p{Lu}"

	// Define the regular expression patterns for each token type.<RegularExpressions>
)
`,

	tokenIdentifier_: `
		<~TokenName>Token: "<~tokenName>",`,

	tokenMatcher_: `
		<~TokenName>Token: reg.MustCompile("^" + <~tokenName>_),`,

	regularExpression_: `
	<~expressionName>_ = <ExpressionValue>`,
}
