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

func ProcessorSynthesizer() ProcessorSynthesizerClassLike {
	return processorSynthesizerReference()
}

// Constructor Methods

func (c *processorSynthesizerClass_) Make(
	syntax not.SyntaxLike,
) ProcessorSynthesizerLike {
	var instance = &processorSynthesizer_{
		// Initialize the instance attributes.
		analyzer_: ana.SyntaxAnalyzer().Make(syntax),
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *processorSynthesizer_) GetClass() ProcessorSynthesizerClassLike {
	return processorSynthesizerReference()
}

// TemplateDriven Methods

func (v *processorSynthesizer_) CreateLegalNotice() string {
	var legalNotice = v.analyzer_.GetLegalNotice()
	return legalNotice
}

func (v *processorSynthesizer_) CreateAccessFunction() string {
	var accessFunction = processorSynthesizerReference().accessFunction_
	return accessFunction
}

func (v *processorSynthesizer_) CreateConstructorMethods() string {
	var constructorMethods = processorSynthesizerReference().constructorMethods_
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
	var principalMethods = processorSynthesizerReference().principalMethods_
	return principalMethods
}

func (v *processorSynthesizer_) CreateAttributeMethods() string {
	var attributeMethods string
	return attributeMethods
}

func (v *processorSynthesizer_) CreateAspectMethods() string {
	var aspectMethods = processorSynthesizerReference().aspectMethods_
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
	var privateMethods = processorSynthesizerReference().privateMethods_
	return privateMethods
}

func (v *processorSynthesizer_) CreateInstanceStructure() string {
	var instanceStructure = processorSynthesizerReference().instanceStructure_
	return instanceStructure
}

func (v *processorSynthesizer_) CreateClassStructure() string {
	var classStructure = processorSynthesizerReference().classStructure_
	return classStructure
}

func (v *processorSynthesizer_) CreateClassReference() string {
	var classReference = processorSynthesizerReference().classReference_
	return classReference
}

func (v *processorSynthesizer_) PerformGlobalUpdates(
	source string,
) string {
	var classImports = processorSynthesizerReference().classImports_
	source = uti.ReplaceAll(
		source,
		"classImports",
		classImports,
	)
	return source
}

// PROTECTED INTERFACE

// Private Methods

func (v *processorSynthesizer_) createProcessRule(
	ruleName string,
) string {
	var processRule = processorSynthesizerReference().processRule_
	if v.analyzer_.IsPlural(ruleName) {
		processRule = processorSynthesizerReference().processIndexedRule_
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
	processToken = processorSynthesizerReference().processToken_
	if v.analyzer_.IsPlural(tokenName) {
		processToken = processorSynthesizerReference().processIndexedToken_
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
	classImports_        string
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

func processorSynthesizerReference() *processorSynthesizerClass_ {
	return processorSynthesizerReference_
}

var processorSynthesizerReference_ = &processorSynthesizerClass_{
	// Initialize the class constants.
	classImports_: `
import (
	ast "<ModuleName>/ast"
)
`,

	accessFunction_: `
// Access Function

func Processor() ProcessorClassLike {
	return processorReference()
}
`,

	constructorMethods_: `
// Constructor Methods

func (c *processorClass_) Make() ProcessorLike {
	var instance = &processor_{
		// Initialize the instance attributes.
	}
	return instance
}
`,

	principalMethods_: `
// Principal Methods

func (v *processor_) GetClass() ProcessorClassLike {
	return processorReference()
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

func processorReference() *processorClass_ {
	return processorReference_
}

var processorReference_ = &processorClass_{
	// Initialize the class constants.
}
`,
}
