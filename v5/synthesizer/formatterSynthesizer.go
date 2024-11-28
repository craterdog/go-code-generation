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

func FormatterSynthesizer() FormatterSynthesizerClassLike {
	return formatterSynthesizerReference()
}

// Constructor Methods

func (c *formatterSynthesizerClass_) Make(
	syntax not.SyntaxLike,
) FormatterSynthesizerLike {
	var instance = &formatterSynthesizer_{
		// Initialize the instance attributes.
		analyzer_: ana.SyntaxAnalyzer().Make(syntax),
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *formatterSynthesizer_) GetClass() FormatterSynthesizerClassLike {
	return formatterSynthesizerReference()
}

// TemplateDriven Methods

func (v *formatterSynthesizer_) CreateLegalNotice() string {
	var legalNotice = v.analyzer_.GetLegalNotice()
	return legalNotice
}

func (v *formatterSynthesizer_) CreateAccessFunction() string {
	var accessFunction = formatterSynthesizerReference().accessFunction_
	return accessFunction
}

func (v *formatterSynthesizer_) CreateConstructorMethods() string {
	var constructorMethods = formatterSynthesizerReference().constructorMethods_
	return constructorMethods
}

func (v *formatterSynthesizer_) CreateConstantMethods() string {
	var constantMethods string
	return constantMethods
}

func (v *formatterSynthesizer_) CreateFunctionMethods() string {
	var functionMethods string
	return functionMethods
}

func (v *formatterSynthesizer_) CreatePrincipalMethods() string {
	var principalMethods = formatterSynthesizerReference().principalMethods_
	return principalMethods
}

func (v *formatterSynthesizer_) CreateAttributeMethods() string {
	var attributeMethods string
	return attributeMethods
}

func (v *formatterSynthesizer_) CreateAspectMethods() string {
	var aspectMethods = formatterSynthesizerReference().aspectMethods_
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

func (v *formatterSynthesizer_) CreatePrivateMethods() string {
	var privateMethods = formatterSynthesizerReference().privateMethods_
	return privateMethods
}

func (v *formatterSynthesizer_) CreateInstanceStructure() string {
	var instanceStructure = formatterSynthesizerReference().instanceStructure_
	return instanceStructure
}

func (v *formatterSynthesizer_) CreateClassStructure() string {
	var classStructure = formatterSynthesizerReference().classStructure_
	return classStructure
}

func (v *formatterSynthesizer_) CreateClassReference() string {
	var classReference = formatterSynthesizerReference().classReference_
	return classReference
}

func (v *formatterSynthesizer_) PerformGlobalUpdates(
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
	var importedPackages = formatterSynthesizerReference().importedPackages_
	source = uti.ReplaceAll(
		source,
		"importedPackages",
		importedPackages,
	)
	return source
}

// PROTECTED INTERFACE

// Private Methods

func (v *formatterSynthesizer_) createProcessRule(
	ruleName string,
) string {
	var processRule = formatterSynthesizerReference().processRule_
	if v.analyzer_.IsPlural(ruleName) {
		processRule = formatterSynthesizerReference().processIndexedRule_
	}
	processRule = uti.ReplaceAll(
		processRule,
		"ruleName",
		ruleName,
	)
	return processRule
}

func (v *formatterSynthesizer_) createProcessRules() string {
	var processRules string
	var ruleNames = v.analyzer_.GetRuleNames().GetIterator()
	for ruleNames.HasNext() {
		var ruleName = ruleNames.GetNext()
		var processRule = v.createProcessRule(ruleName)
		processRules += processRule
	}
	return processRules
}

func (v *formatterSynthesizer_) createProcessToken(
	tokenName string,
) string {
	var processToken string
	if tokenName == "delimiter" {
		return processToken
	}
	processToken = formatterSynthesizerReference().processToken_
	if v.analyzer_.IsPlural(tokenName) {
		processToken = formatterSynthesizerReference().processIndexedToken_
	}
	if tokenName == "newline" {
		processToken = formatterSynthesizerReference().processNewline_
	}
	processToken = uti.ReplaceAll(
		processToken,
		"tokenName",
		tokenName,
	)
	return processToken
}

func (v *formatterSynthesizer_) createProcessTokens() string {
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

type formatterSynthesizer_ struct {
	// Declare the instance attributes.
	analyzer_ ana.SyntaxAnalyzerLike
}

// Class Structure

type formatterSynthesizerClass_ struct {
	// Declare the class constants.
	importedPackages_    string
	accessFunction_      string
	constructorMethods_  string
	principalMethods_    string
	aspectMethods_       string
	processToken_        string
	processNewline_      string
	processIndexedToken_ string
	processRule_         string
	processIndexedRule_  string
	privateMethods_      string
	instanceStructure_   string
	classStructure_      string
	classReference_      string
}

// Class Reference

func formatterSynthesizerReference() *formatterSynthesizerClass_ {
	return formatterSynthesizerReference_
}

var formatterSynthesizerReference_ = &formatterSynthesizerClass_{
	// Initialize the class constants.
	importedPackages_: `
	ast "<ModuleName>/ast"
	sts "strings"
`,

	accessFunction_: `
// Access Function

func Formatter() FormatterClassLike {
	return formatterReference()
}
`,

	constructorMethods_: `
// Constructor Methods

func (c *formatterClass_) Make() FormatterLike {
	var instance = &formatter_{
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

func (v *formatter_) GetClass() FormatterClassLike {
	return formatterReference()
}

func (v *formatter_) Format<~SyntaxName>(<syntaxName_> ast.<~SyntaxName>Like) string {
	v.visitor_.Visit<~SyntaxName>(<syntaxName_>)
	return v.getResult()
}
`,

	aspectMethods_: `
// Methodical Methods
<ProcessTokens><ProcessRules>`,

	processToken_: `
func (v *formatter_) Process<~TokenName>(
	<tokenName_> string,
) {
	v.appendString(<tokenName_>)
}
`,

	processNewline_: `
func (v *formatter_) Process<~TokenName>(
	<tokenName_> string,
) {
	v.appendNewline()
}
`,

	processIndexedToken_: `
func (v *formatter_) Process<~TokenName>(
	<tokenName_> string,
	index uint,
	size uint,
) {
	v.appendString(<tokenName_>)
}
`,

	processRule_: `
func (v *formatter_) Preprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) Process<~RuleName>Slot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) Postprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
) {
	// TBD - Add formatting of the delimited rule.
}
`,

	processIndexedRule_: `
func (v *formatter_) Preprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) Process<~RuleName>Slot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) Postprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}
`,

	privateMethods_: `
// Private Methods

func (v *formatter_) appendNewline() {
	var newline = "\n"
	var indentation = "\t"
	var level uint
	for ; level < v.depth_; level++ {
		newline += indentation
	}
	v.appendString(newline)
}

func (v *formatter_) appendString(
	string_ string,
) {
	v.result_.WriteString(string_)
}

func (v *formatter_) getResult() string {
	var result = v.result_.String()
	v.result_.Reset()
	return result
}
`,

	instanceStructure_: `
// Instance Structure

type formatter_ struct {
	// Declare the instance attributes.
	visitor_ VisitorLike
	depth_   uint
	result_  sts.Builder

	// Declare the inherited aspects.
	Methodical
}
`,

	classStructure_: `
// Class Structure

type formatterClass_ struct {
	// Declare the class constants.
}
`,

	classReference_: `
// Class Reference

func formatterReference() *formatterClass_ {
	return formatterReference_
}

var formatterReference_ = &formatterClass_{
	// Initialize the class constants.
}
`,
}
