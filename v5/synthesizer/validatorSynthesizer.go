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
	return validatorSynthesizerReference()
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

// Primary Methods

func (v *validatorSynthesizer_) GetClass() ValidatorSynthesizerClassLike {
	return validatorSynthesizerReference()
}

// TemplateDriven Methods

func (v *validatorSynthesizer_) CreateLegalNotice() string {
	var legalNotice = v.analyzer_.GetLegalNotice()
	return legalNotice
}

func (v *validatorSynthesizer_) CreateAccessFunction() string {
	var accessFunction = validatorSynthesizerReference().accessFunction_
	return accessFunction
}

func (v *validatorSynthesizer_) CreateConstructorMethods() string {
	var constructorMethods = validatorSynthesizerReference().constructorMethods_
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

func (v *validatorSynthesizer_) CreatePrimaryMethods() string {
	var primaryMethods = validatorSynthesizerReference().primaryMethods_
	return primaryMethods
}

func (v *validatorSynthesizer_) CreateAttributeMethods() string {
	var attributeMethods string
	return attributeMethods
}

func (v *validatorSynthesizer_) CreateAspectMethods() string {
	var aspectMethods = validatorSynthesizerReference().aspectMethods_
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
	var privateMethods = validatorSynthesizerReference().privateMethods_
	return privateMethods
}

func (v *validatorSynthesizer_) CreateInstanceStructure() string {
	var instanceStructure = validatorSynthesizerReference().instanceStructure_
	return instanceStructure
}

func (v *validatorSynthesizer_) CreateClassStructure() string {
	var classStructure = validatorSynthesizerReference().classStructure_
	return classStructure
}

func (v *validatorSynthesizer_) CreateClassReference() string {
	var classReference = validatorSynthesizerReference().classReference_
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
	var classImports = validatorSynthesizerReference().classImports_
	source = uti.ReplaceAll(
		source,
		"classImports",
		classImports,
	)
	return source
}

// PROTECTED INTERFACE

// Private Methods

func (v *validatorSynthesizer_) createProcessRule(
	ruleName string,
) string {
	var processRule = validatorSynthesizerReference().processRule_
	if v.analyzer_.IsPlural(ruleName) {
		processRule = validatorSynthesizerReference().processIndexedRule_
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
	processToken = validatorSynthesizerReference().processToken_
	if v.analyzer_.IsPlural(tokenName) {
		processToken = validatorSynthesizerReference().processIndexedToken_
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
	classImports_        string
	accessFunction_      string
	constructorMethods_  string
	primaryMethods_      string
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

func validatorSynthesizerReference() *validatorSynthesizerClass_ {
	return validatorSynthesizerReference_
}

var validatorSynthesizerReference_ = &validatorSynthesizerClass_{
	// Initialize the class constants.
	classImports_: `
import (
	fmt "fmt"
	ast "<ModuleName>/ast"
)
`,

	accessFunction_: `
// Access Function

func Validator() ValidatorClassLike {
	return validatorReference()
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

	primaryMethods_: `
// Primary Methods

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
