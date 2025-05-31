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

func GrammarSynthesizerClass() GrammarSynthesizerClassLike {
	return grammarSynthesizerClass()
}

// Constructor Methods

func (c *grammarSynthesizerClass_) GrammarSynthesizer(
	syntax not.SyntaxLike,
) GrammarSynthesizerLike {
	var instance = &grammarSynthesizer_{
		// Initialize the instance attributes.
		analyzer_: ana.SyntaxAnalyzerClass().SyntaxAnalyzer(syntax),
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *grammarSynthesizer_) GetClass() GrammarSynthesizerClassLike {
	return grammarSynthesizerClass()
}

// TemplateDriven Methods

func (v *grammarSynthesizer_) CreateLegalNotice() string {
	var legalNotice = v.analyzer_.GetLegalNotice()
	return legalNotice
}

func (v *grammarSynthesizer_) CreateWarningMessage() string {
	var class = grammarSynthesizerClass()
	var warningMessage = class.warningMessage_
	return warningMessage
}

func (v *grammarSynthesizer_) CreatePackageDescription() string {
	var class = grammarSynthesizerClass()
	var packageDescription = class.packageDescription_
	return packageDescription
}

func (v *grammarSynthesizer_) CreateImportedPackages() string {
	var class = grammarSynthesizerClass()
	var importedPackages = class.importedPackages_
	return importedPackages
}

func (v *grammarSynthesizer_) CreateTypeDeclarations() string {
	var class = grammarSynthesizerClass()
	var typeDeclarations = class.typeDeclarations_
	var tokenTypes = v.generateTokenTypes()
	typeDeclarations = uti.ReplaceAll(
		typeDeclarations,
		"tokenTypes",
		tokenTypes,
	)
	return typeDeclarations
}

func (v *grammarSynthesizer_) CreateFunctionalDeclarations() string {
	var functionalDeclarations string
	return functionalDeclarations
}

func (v *grammarSynthesizer_) CreateClassDeclarations() string {
	var class = grammarSynthesizerClass()
	var classDeclarations = class.classDeclarations_
	return classDeclarations
}

func (v *grammarSynthesizer_) CreateInstanceDeclarations() string {
	var syntaxName = v.analyzer_.GetSyntaxName()
	var class = grammarSynthesizerClass()
	var instanceDeclarations = class.instanceDeclarations_
	instanceDeclarations = uti.ReplaceAll(
		instanceDeclarations,
		"syntaxName",
		syntaxName,
	)
	return instanceDeclarations
}

func (v *grammarSynthesizer_) CreateAspectDeclarations() string {
	var class = grammarSynthesizerClass()
	var aspectDeclarations = class.aspectDeclarations_
	var processTokens = v.generateProcessTokens()
	aspectDeclarations = uti.ReplaceAll(
		aspectDeclarations,
		"processTokens",
		processTokens,
	)
	var processRules = v.generateProcessRules()
	aspectDeclarations = uti.ReplaceAll(
		aspectDeclarations,
		"processRules",
		processRules,
	)
	return aspectDeclarations
}

func (v *grammarSynthesizer_) PerformGlobalUpdates(
	existing string,
	generated string,
) string {
	return generated
}

// PROTECTED INTERFACE

// Private Methods

func (v *grammarSynthesizer_) generateProcessRule(
	ruleName string,
) string {
	var class = grammarSynthesizerClass()
	var processRule = class.processRule_
	processRule = uti.ReplaceAll(
		processRule,
		"ruleName",
		ruleName,
	)
	return processRule
}

func (v *grammarSynthesizer_) generateProcessRules() string {
	var processRules string
	var ruleNames = v.analyzer_.GetRules().GetIterator()
	for ruleNames.HasNext() {
		var ruleName = ruleNames.GetNext()
		var processRule = v.generateProcessRule(ruleName)
		processRules += processRule
	}
	return processRules
}

func (v *grammarSynthesizer_) generateProcessToken(
	tokenName string,
) string {
	var class = grammarSynthesizerClass()
	var processToken = class.processToken_
	processToken = uti.ReplaceAll(
		processToken,
		"tokenName",
		tokenName,
	)
	return processToken
}

func (v *grammarSynthesizer_) generateProcessTokens() string {
	var processTokens string
	var tokenNames = v.analyzer_.GetExpressions().GetIterator()
	for tokenNames.HasNext() {
		var tokenName = tokenNames.GetNext()
		var processToken = v.generateProcessToken(tokenName)
		processTokens += processToken
	}
	return processTokens
}

func (v *grammarSynthesizer_) generateTokenTypes() string {
	var tokenTypes string
	var tokenNames = v.analyzer_.GetExpressions().GetIterator()
	for tokenNames.HasNext() {
		var tokenName = tokenNames.GetNext()
		var class = grammarSynthesizerClass()
		var tokenType = class.tokenType_
		tokenType = uti.ReplaceAll(
			tokenType,
			"tokenName",
			tokenName,
		)
		tokenTypes += tokenType
	}
	return tokenTypes
}

// Instance Structure

type grammarSynthesizer_ struct {
	// Declare the instance attributes.
	analyzer_ ana.SyntaxAnalyzerLike
}

// Class Structure

type grammarSynthesizerClass_ struct {
	// Declare the class constants.
	warningMessage_       string
	packageDescription_   string
	importedPackages_     string
	typeDeclarations_     string
	tokenType_            string
	classDeclarations_    string
	instanceDeclarations_ string
	aspectDeclarations_   string
	processToken_         string
	processRule_          string
}

// Class Reference

func grammarSynthesizerClass() *grammarSynthesizerClass_ {
	return grammarSynthesizerClassReference_
}

var grammarSynthesizerClassReference_ = &grammarSynthesizerClass_{
	// Initialize the class constants.
	warningMessage_: `
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│            This "package_api.go" file was automatically generated.           │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
`,

	packageDescription_: `
Package "grammar" provides the following grammar classes that operate on the
abstract syntax tree (AST) for this module:
  - Token captures the attributes associated with a parsed token.
  - Scanner is used to scan the source byte stream and recognize matching tokens.
  - Parser is used to process the token stream and generate the AST.
  - Validator is used to validate the semantics associated with an AST.
  - Formatter is used to format an AST back into a canonical version of its source.
  - Visitor walks the AST and calls processor methods for each node in the tree.
  - Processor provides empty processor methods to be inherited by the processors.`,

	importedPackages_: `
	ast "<ModuleName>/ast"
	col "github.com/craterdog/go-collection-framework/v7"
`,

	typeDeclarations_: `
/*
TokenType is a constrained type representing any token type recognized by a
scanner.
*/
type TokenType uint8

const (
	ErrorToken TokenType = iota<TokenTypes>
)
`,

	tokenType_: `
	<~TokenName>Token`,

	classDeclarations_: `
/*
FormatterClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete formatter-like class.
*/
type FormatterClassLike interface {
	// Constructor Methods
	Formatter() FormatterLike
}

/*
ParserClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete parser-like class.
*/
type ParserClassLike interface {
	// Constructor Methods
	Parser() ParserLike
}

/*
ProcessorClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete processor-like class.
*/
type ProcessorClassLike interface {
	// Constructor Methods
	Processor() ProcessorLike
}

/*
ScannerClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete scanner-like class.  The following functions are supported:

FormatToken() returns a formatted string containing the attributes of the token.

FormatType() returns the string version of the token type.

MatchesType() determines whether or not a token value is of a specified type.
*/
type ScannerClassLike interface {
	// Constructor Methods
	Scanner(
		source string,
		tokens col.QueueLike[TokenLike],
	) ScannerLike

	// Function Methods
	FormatToken(
		token TokenLike,
	) string
	FormatType(
		tokenType TokenType,
	) string
	MatchesType(
		tokenValue string,
		tokenType TokenType,
	) bool
}

/*
TokenClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete visitor-like class.
*/
type TokenClassLike interface {
	// Constructor Methods
	Token(
		line uint,
		position uint,
		type_ TokenType,
		value string,
	) TokenLike
}

/*
ValidatorClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete validator-like class.
*/
type ValidatorClassLike interface {
	// Constructor Methods
	Validator() ValidatorLike
}

/*
VisitorClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete visitor-like class.
*/
type VisitorClassLike interface {
	// Constructor Methods
	Visitor(
		processor Methodical,
	) VisitorLike
}
`,

	instanceDeclarations_: `
/*
FormatterLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete formatter-like class.
*/
type FormatterLike interface {
	// Principal Methods
	GetClass() FormatterClassLike
	Format<~SyntaxName>(
		<syntaxName_> ast.<~SyntaxName>Like,
	) string

	// Aspect Interfaces
	Methodical
}

/*
ParserLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete parser-like class.
*/
type ParserLike interface {
	// Principal Methods
	GetClass() ParserClassLike
	ParseSource(
		source string,
	) ast.<~SyntaxName>Like
}

/*
ProcessorLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete processor-like class.
*/
type ProcessorLike interface {
	// Principal Methods
	GetClass() ProcessorClassLike

	// Aspect Interfaces
	Methodical
}

/*
ScannerLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete scanner-like class.
*/
type ScannerLike interface {
	// Principal Methods
	GetClass() ScannerClassLike
}

/*
TokenLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete token-like class.
*/
type TokenLike interface {
	// Principal Methods
	GetClass() TokenClassLike

	// Attribute Methods
	GetLine() uint
	GetPosition() uint
	GetType() TokenType
	GetValue() string
}

/*
ValidatorLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete validator-like class.
*/
type ValidatorLike interface {
	// Principal Methods
	GetClass() ValidatorClassLike
	Validate<~SyntaxName>(
		<syntaxName_> ast.<~SyntaxName>Like,
	)

	// Aspect Interfaces
	Methodical
}

/*
VisitorLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete visitor-like class.
*/
type VisitorLike interface {
	// Principal Methods
	GetClass() VisitorClassLike
	Visit<~SyntaxName>(
		<syntaxName_> ast.<~SyntaxName>Like,
	)
}
`,

	aspectDeclarations_: `
/*
Methodical declares the set of method signatures that must be supported by
all methodical processors.
*/
type Methodical interface {<ProcessTokens><ProcessRules>
}
`,

	processToken_: `
	Process<~TokenName>(
		<tokenName_> string,
	)`,

	processRule_: `
	Preprocess<~RuleName>(
		<ruleName_> ast.<~RuleName>Like,
		index_ uint,
		count_ uint,
	)
	Postprocess<~RuleName>(
		<ruleName_> ast.<~RuleName>Like,
		index_ uint,
		count_ uint,
	)
	Process<~RuleName>Slot(
		<ruleName_> ast.<~RuleName>Like,
		slot_ uint,
	)`,
}
