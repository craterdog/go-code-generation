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

func ProcessorSynthesizerClass() ProcessorSynthesizerClassLike {
	return processorSynthesizerClass()
}

// Constructor Methods

func (c *processorSynthesizerClass_) ProcessorSynthesizer(
	syntax not.SyntaxLike,
) ProcessorSynthesizerLike {
	var instance = &processorSynthesizer_{
		// Initialize the instance attributes.
		analyzer_: ana.SyntaxAnalyzerClass().SyntaxAnalyzer(syntax),
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *processorSynthesizer_) GetClass() ProcessorSynthesizerClassLike {
	return processorSynthesizerClass()
}

// TemplateDriven Methods

func (v *processorSynthesizer_) CreateLegalNotice() string {
	var legalNotice = v.analyzer_.GetLegalNotice()
	return legalNotice
}

func (v *processorSynthesizer_) CreateWarningMessage() string {
	var class = processorSynthesizerClass()
	var warningMessage = class.warningMessage_
	return warningMessage
}

func (v *processorSynthesizer_) CreateImportedPackages() string {
	var class = processorSynthesizerClass()
	var importedPackages = class.importedPackages_
	return importedPackages
}

func (v *processorSynthesizer_) CreateAccessFunction() string {
	var class = processorSynthesizerClass()
	var accessFunction = class.accessFunction_
	return accessFunction
}

func (v *processorSynthesizer_) CreateConstructorMethods() string {
	var class = processorSynthesizerClass()
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
	var class = processorSynthesizerClass()
	var principalMethods = class.principalMethods_
	return principalMethods
}

func (v *processorSynthesizer_) CreateAttributeMethods() string {
	var attributeMethods string
	return attributeMethods
}

func (v *processorSynthesizer_) CreateAspectMethods() string {
	var class = processorSynthesizerClass()
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
	var class = processorSynthesizerClass()
	var privateMethods = class.privateMethods_
	return privateMethods
}

func (v *processorSynthesizer_) CreateInstanceStructure() string {
	var class = processorSynthesizerClass()
	var instanceStructure = class.instanceStructure_
	return instanceStructure
}

func (v *processorSynthesizer_) CreateClassStructure() string {
	var class = processorSynthesizerClass()
	var classStructure = class.classStructure_
	return classStructure
}

func (v *processorSynthesizer_) CreateClass() string {
	var class = processorSynthesizerClass()
	var classReference = class.classReference_
	return classReference
}

func (v *processorSynthesizer_) PerformGlobalUpdates(
	existing string,
	generated string,
) string {
	return generated
}

// PROTECTED INTERFACE

// Private Methods

func (v *processorSynthesizer_) createProcessToken(
	tokenName string,
) string {
	var processToken string
	var class = processorSynthesizerClass()
	processToken = class.processToken_
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

func (v *processorSynthesizer_) createProcessRule(
	ruleName string,
) string {
	var class = processorSynthesizerClass()
	var processRule = class.processRule_
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

// Instance Structure

type processorSynthesizer_ struct {
	// Declare the instance attributes.
	analyzer_ ana.SyntaxAnalyzerLike
}

// Class Structure

type processorSynthesizerClass_ struct {
	// Declare the class constants.
	warningMessage_     string
	importedPackages_   string
	accessFunction_     string
	constructorMethods_ string
	principalMethods_   string
	aspectMethods_      string
	processToken_       string
	processRule_        string
	privateMethods_     string
	instanceStructure_  string
	classStructure_     string
	classReference_     string
}

// Class Reference

func processorSynthesizerClass() *processorSynthesizerClass_ {
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
	return processorClass()
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
	return processorClass()
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

	processRule_: `
func (v *processor_) Preprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
	index uint,
	count uint,
) {
}

func (v *processor_) Process<~RuleName>Slot(
	slot uint,
) {
}

func (v *processor_) Postprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
	index uint,
	count uint,
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

func processorClass() *processorClass_ {
	return processorClassReference_
}

var processorClassReference_ = &processorClass_{
	// Initialize the class constants.
}
`,
}
