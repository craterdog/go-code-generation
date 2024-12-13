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

func ProcessorSynthesizerClass() ProcessorSynthesizerClassLike {
	return processorSynthesizerClassReference()
}

// Constructor Methods

func (c *processorSynthesizerClass_) Make(
	syntax not.SyntaxLike,
) ProcessorSynthesizerLike {
	var instance = &processorSynthesizer_{
		// Initialize the instance attributes.
		analyzer_: ana.SyntaxAnalyzerClass().Make(syntax),
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *processorSynthesizer_) GetClass() ProcessorSynthesizerClassLike {
	return processorSynthesizerClassReference()
}

// TemplateDriven Methods

func (v *processorSynthesizer_) CreateLegalNotice() string {
	var legalNotice = v.analyzer_.GetLegalNotice()
	return legalNotice
}

func (v *processorSynthesizer_) CreateWarningMessage() string {
	var class = processorSynthesizerClassReference()
	var warningMessage = class.warningMessage_
	return warningMessage
}

func (v *processorSynthesizer_) CreateAccessFunction() string {
	var class = processorSynthesizerClassReference()
	var accessFunction = class.accessFunction_
	return accessFunction
}

func (v *processorSynthesizer_) CreateConstructorMethods() string {
	var class = processorSynthesizerClassReference()
	var constructorMethods = class.constructorMethods_
	return constructorMethods
}

func (v *processorSynthesizer_) CreateConstantMethods() string {
	var constantMethods string
	return constantMethods
}

func (v *processorSynthesizer_) CreateFunctionMethods() string {
	var functionMethods string
	return functionMethods
}

func (v *processorSynthesizer_) CreatePrincipalMethods() string {
	var class = processorSynthesizerClassReference()
	var principalMethods = class.principalMethods_
	return principalMethods
}

func (v *processorSynthesizer_) CreateAttributeMethods() string {
	var attributeMethods string
	return attributeMethods
}

func (v *processorSynthesizer_) CreateAspectMethods() string {
	var class = processorSynthesizerClassReference()
	var aspectMethods = class.aspectMethods_
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

func (v *processorSynthesizer_) CreatePrivateMethods() string {
	var class = processorSynthesizerClassReference()
	var privateMethods = class.privateMethods_
	return privateMethods
}

func (v *processorSynthesizer_) CreateInstanceStructure() string {
	var class = processorSynthesizerClassReference()
	var instanceStructure = class.instanceStructure_
	return instanceStructure
}

func (v *processorSynthesizer_) CreateClassStructure() string {
	var class = processorSynthesizerClassReference()
	var classStructure = class.classStructure_
	return classStructure
}

func (v *processorSynthesizer_) CreateClassReference() string {
	var class = processorSynthesizerClassReference()
	var classReference = class.classReference_
	return classReference
}

func (v *processorSynthesizer_) PerformGlobalUpdates(
	source string,
) string {
	var class = processorSynthesizerClassReference()
	var importedPackages = class.importedPackages_
	source = uti.ReplaceAll(
		source,
		"importedPackages",
		importedPackages,
	)
	return source
}

// PROTECTED INTERFACE

// Private Methods

func (v *processorSynthesizer_) createProcessRule(
	ruleName string,
) string {
	var class = processorSynthesizerClassReference()
	var processRule = class.processRule_
	if v.analyzer_.IsPlural(ruleName) {
		processRule = class.processIndexedRule_
	}
	processRule = uti.ReplaceAll(
		processRule,
		"ruleName",
		ruleName,
	)
	return processRule
}

func (v *processorSynthesizer_) createProcessRules() string {
	var processRules string
	var ruleNames = v.analyzer_.GetRuleNames().GetIterator()
	for ruleNames.HasNext() {
		var ruleName = ruleNames.GetNext()
		var processRule = v.createProcessRule(ruleName)
		processRules += processRule
	}
	return processRules
}

func (v *processorSynthesizer_) createProcessToken(
	tokenName string,
) string {
	var processToken string
	if tokenName == "delimiter" {
		return processToken
	}
	var class = processorSynthesizerClassReference()
	processToken = class.processToken_
	if v.analyzer_.IsPlural(tokenName) {
		processToken = class.processIndexedToken_
	}
	processToken = uti.ReplaceAll(
		processToken,
		"tokenName",
		tokenName,
	)
	return processToken
}

func (v *processorSynthesizer_) createProcessTokens() string {
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

type processorSynthesizer_ struct {
	// Declare the instance attributes.
	analyzer_ ana.SyntaxAnalyzerLike
}

// Class Structure

type processorSynthesizerClass_ struct {
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

func processorSynthesizerClassReference() *processorSynthesizerClass_ {
	return processorSynthesizerClassReference_
}

var processorSynthesizerClassReference_ = &processorSynthesizerClass_{
	// Initialize the class constants.
	warningMessage_: `
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│                 This class file was automatically generated.                 │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
`,

	importedPackages_: `
	ast "<ModuleName>/ast"
`,

	accessFunction_: `
// Access Function

func ProcessorClass() ProcessorClassLike {
	return processorClassReference()
}
`,

	constructorMethods_: `
// Constructor Methods

func (c *processorClass_) Processor() ProcessorLike {
	var instance = &processor_{
		// Initialize the instance attributes.
	}
	return instance
}
`,

	principalMethods_: `
// Principal Methods

func (v *processor_) GetClass() ProcessorClassLike {
	return processorClassReference()
}
`,

	aspectMethods_: `
// Methodical Methods
<ProcessTokens><ProcessRules>`,

	processToken_: `
func (v *processor_) Process<~TokenName>(
	<tokenName_> string,
) {
}
`,

	processIndexedToken_: `
func (v *processor_) Process<~TokenName>(
	<tokenName_> string,
	index uint,
	size uint,
) {
}
`,

	processRule_: `
func (v *processor_) Preprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
) {
}

func (v *processor_) Process<~RuleName>Slot(
	slot uint,
) {
}

func (v *processor_) Postprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
) {
}
`,

	processIndexedRule_: `
func (v *processor_) Preprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
	index uint,
	size uint,
) {
}

func (v *processor_) Process<~RuleName>Slot(
	slot uint,
) {
}

func (v *processor_) Postprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
	index uint,
	size uint,
) {
}
`,

	privateMethods_: `
// Private Methods
`,

	instanceStructure_: `
// Instance Structure

type processor_ struct {
	// Declare the instance attributes.
}
`,

	classStructure_: `
// Class Structure

type processorClass_ struct {
	// Declare the class constants.
}
`,

	classReference_: `
// Class Reference

func processorClassReference() *processorClass_ {
	return processorClassReference_
}

var processorClassReference_ = &processorClass_{
	// Initialize the class constants.
}
`,
}
