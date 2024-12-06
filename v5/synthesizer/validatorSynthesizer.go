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

func ValidatorSynthesizer() ValidatorSynthesizerClassLike {
	return validatorSynthesizerClassReference()
}

// Constructor Methods

func (c *validatorSynthesizerClass_) Make(
	syntax not.SyntaxLike,
) ValidatorSynthesizerLike {
	var instance = &validatorSynthesizer_{
		// Initialize the instance attributes.
		analyzer_: ana.SyntaxAnalyzer().Make(syntax),
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *validatorSynthesizer_) GetClass() ValidatorSynthesizerClassLike {
	return validatorSynthesizerClassReference()
}

// TemplateDriven Methods

func (v *validatorSynthesizer_) CreateLegalNotice() string {
	var legalNotice = v.analyzer_.GetLegalNotice()
	return legalNotice
}

func (v *validatorSynthesizer_) CreateWarningMessage() string {
	var warningMessage = validatorSynthesizerClassReference().warningMessage_
	return warningMessage
}

func (v *validatorSynthesizer_) CreateAccessFunction() string {
	var accessFunction = validatorSynthesizerClassReference().accessFunction_
	return accessFunction
}

func (v *validatorSynthesizer_) CreateConstructorMethods() string {
	var constructorMethods = validatorSynthesizerClassReference().constructorMethods_
	return constructorMethods
}

func (v *validatorSynthesizer_) CreateConstantMethods() string {
	var constantMethods string
	return constantMethods
}

func (v *validatorSynthesizer_) CreateFunctionMethods() string {
	var functionMethods string
	return functionMethods
}

func (v *validatorSynthesizer_) CreatePrincipalMethods() string {
	var principalMethods = validatorSynthesizerClassReference().principalMethods_
	return principalMethods
}

func (v *validatorSynthesizer_) CreateAttributeMethods() string {
	var attributeMethods string
	return attributeMethods
}

func (v *validatorSynthesizer_) CreateAspectMethods() string {
	var aspectMethods = validatorSynthesizerClassReference().aspectMethods_
	var processTokens = v.createProcessTokens()
	aspectMethods = uti.ReplaceAll(
		aspectMethods,
		"processTokens",
		processTokens,
	)
	var processRules = v.createProcessRules()
	aspectMethods = uti.ReplaceAll(
		aspectMethods,
		"processRules",
		processRules,
	)
	return aspectMethods
}

func (v *validatorSynthesizer_) CreatePrivateMethods() string {
	var privateMethods = validatorSynthesizerClassReference().privateMethods_
	return privateMethods
}

func (v *validatorSynthesizer_) CreateInstanceStructure() string {
	var instanceStructure = validatorSynthesizerClassReference().instanceStructure_
	return instanceStructure
}

func (v *validatorSynthesizer_) CreateClassStructure() string {
	var classStructure = validatorSynthesizerClassReference().classStructure_
	return classStructure
}

func (v *validatorSynthesizer_) CreateClassReference() string {
	var classReference = validatorSynthesizerClassReference().classReference_
	return classReference
}

func (v *validatorSynthesizer_) PerformGlobalUpdates(
	source string,
) string {
	var syntaxName = v.analyzer_.GetSyntaxName()
	source = uti.ReplaceAll(
		source,
		"syntaxName",
		syntaxName,
	)
	var importedPackages = validatorSynthesizerClassReference().importedPackages_
	source = uti.ReplaceAll(
		source,
		"importedPackages",
		importedPackages,
	)
	return source
}

// PROTECTED INTERFACE

// Private Methods

func (v *validatorSynthesizer_) createProcessRule(
	ruleName string,
) string {
	var processRule = validatorSynthesizerClassReference().processRule_
	if v.analyzer_.IsPlural(ruleName) {
		processRule = validatorSynthesizerClassReference().processIndexedRule_
	}
	processRule = uti.ReplaceAll(
		processRule,
		"ruleName",
		ruleName,
	)
	return processRule
}

func (v *validatorSynthesizer_) createProcessRules() string {
	var processRules string
	var ruleNames = v.analyzer_.GetRuleNames().GetIterator()
	for ruleNames.HasNext() {
		var ruleName = ruleNames.GetNext()
		var processRule = v.createProcessRule(ruleName)
		processRules += processRule
	}
	return processRules
}

func (v *validatorSynthesizer_) createProcessToken(
	tokenName string,
) string {
	var processToken string
	if tokenName == "delimiter" {
		return processToken
	}
	processToken = validatorSynthesizerClassReference().processToken_
	if v.analyzer_.IsPlural(tokenName) {
		processToken = validatorSynthesizerClassReference().processIndexedToken_
	}
	processToken = uti.ReplaceAll(
		processToken,
		"tokenName",
		tokenName,
	)
	return processToken
}

func (v *validatorSynthesizer_) createProcessTokens() string {
	var processTokens string
	var tokenNames = v.analyzer_.GetTokenNames().GetIterator()
	for tokenNames.HasNext() {
		var tokenName = tokenNames.GetNext()
		var processToken = v.createProcessToken(tokenName)
		processTokens += processToken
	}
	return processTokens
}

// Instance Structure

type validatorSynthesizer_ struct {
	// Declare the instance attributes.
	analyzer_ ana.SyntaxAnalyzerLike
}

// Class Structure

type validatorSynthesizerClass_ struct {
	// Declare the class constants.
	warningMessage_      string
	importedPackages_    string
	accessFunction_      string
	constructorMethods_  string
	principalMethods_    string
	aspectMethods_       string
	processToken_        string
	processIndexedToken_ string
	processRule_         string
	processIndexedRule_  string
	privateMethods_      string
	instanceStructure_   string
	classStructure_      string
	classReference_      string
}

// Class Reference

func validatorSynthesizerClassReference() *validatorSynthesizerClass_ {
	return validatorSynthesizerClassReference_
}

var validatorSynthesizerClassReference_ = &validatorSynthesizerClass_{
	// Initialize the class constants.
	warningMessage_: `
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│                 This class file was automatically generated.                 │
│ Updates to any section other than the Methodical Methods may be overwritten. │
└──────────────────────────────────────────────────────────────────────────────┘
`,

	importedPackages_: `
	fmt "fmt"
	ast "<ModuleName>/ast"
`,

	accessFunction_: `
// Access Function

func ValidatorClass() ValidatorClassLike {
	return validatorClassReference()
}
`,

	constructorMethods_: `
// Constructor Methods

func (c *validatorClass_) Make() ValidatorLike {
	var instance = &validator_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Methodical: Processor().Make(),
	}
	instance.visitor_ = Visitor().Make(instance)
	return instance
}
`,

	principalMethods_: `
// Principal Methods

func (v *validator_) GetClass() ValidatorClassLike {
	return validatorReference()
}

func (v *validator_) Validate<~SyntaxName>(
	<syntaxName_> ast.<~SyntaxName>Like,
) {
	v.visitor_.Visit<~SyntaxName>(<syntaxName_>)
}
`,

	aspectMethods_: `
// Methodical Methods
<ProcessTokens><ProcessRules>`,

	processToken_: `
func (v *validator_) Process<~TokenName>(
	<tokenName_> string,
) {
	v.validateToken(<tokenName_>, <~TokenName>Token)
}
`,

	processIndexedToken_: `
func (v *validator_) Process<~TokenName>(
	<tokenName_> string,
	index uint,
	size uint,
) {
	v.validateToken(<tokenName_>, <~TokenName>Token)
}
`,

	processRule_: `
func (v *validator_) Preprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) Process<~RuleName>Slot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) Postprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
) {
	// TBD - Add any validation checks.
}
`,

	processIndexedRule_: `
func (v *validator_) Preprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) Process<~RuleName>Slot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) Postprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}
`,

	privateMethods_: `
// Private Methods

func (v *validator_) validateToken(
	tokenValue string,
	tokenType TokenType,
) {
	if !Scanner().MatchesType(tokenValue, tokenType) {
		var message = fmt.Sprintf(
			"The following token value is not of type %v: %v",
			Scanner().FormatType(tokenType),
			tokenValue,
		)
		panic(message)
	}
}
`,

	instanceStructure_: `
// Instance Structure

type validator_ struct {
	// Declare the instance attributes.
	visitor_ VisitorLike

	// Declare the inherited aspects.
	Methodical
}
`,

	classStructure_: `
// Class Structure

type validatorClass_ struct {
	// Declare the class constants.
}
`,

	classReference_: `
// Class Reference

func validatorReference() *validatorClass_ {
	return validatorReference_
}

var validatorReference_ = &validatorClass_{
	// Initialize the class constants.
}
`,
}
