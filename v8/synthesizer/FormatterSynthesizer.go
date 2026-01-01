/*
................................................................................
.    Copyright (c) 2009-2026 Crater Dog Technologies™.  All Rights Reserved.   .
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
	ana "github.com/craterdog/go-code-generation/v8/analyzer"
	uti "github.com/craterdog/go-essential-utilities/v8"
	not "github.com/craterdog/go-syntax-notation/v8"
	reg "regexp"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func FormatterSynthesizerClass() FormatterSynthesizerClassLike {
	return formatterSynthesizerClass()
}

// Constructor Methods

func (c *formatterSynthesizerClass_) FormatterSynthesizer(
	syntax not.SyntaxLike,
) FormatterSynthesizerLike {
	var instance = &formatterSynthesizer_{
		// Initialize the instance attributes.
		analyzer_: ana.SyntaxAnalyzerClass().SyntaxAnalyzer(syntax),
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *formatterSynthesizer_) GetClass() FormatterSynthesizerClassLike {
	return formatterSynthesizerClass()
}

// TemplateDriven Methods

func (v *formatterSynthesizer_) CreateLegalNotice() string {
	var legalNotice = v.analyzer_.GetLegalNotice()
	return legalNotice
}

func (v *formatterSynthesizer_) CreateWarningMessage() string {
	var class = formatterSynthesizerClass()
	var warningMessage = class.warningMessage_
	return warningMessage
}

func (v *formatterSynthesizer_) CreateImportedPackages() string {
	var class = formatterSynthesizerClass()
	var importedPackages = class.importedPackages_
	return importedPackages
}

func (v *formatterSynthesizer_) CreateAccessFunction() string {
	var class = formatterSynthesizerClass()
	var accessFunction = class.accessFunction_
	return accessFunction
}

func (v *formatterSynthesizer_) CreateConstructorMethods() string {
	var class = formatterSynthesizerClass()
	var constructorMethods = class.constructorMethods_
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
	var class = formatterSynthesizerClass()
	var principalMethods = class.principalMethods_
	var syntaxMap = v.analyzer_.GetSyntaxMap()
	principalMethods = uti.ReplaceAll(
		principalMethods,
		"syntaxMap",
		syntaxMap,
	)
	var syntaxName = v.analyzer_.GetSyntaxName()
	principalMethods = uti.ReplaceAll(
		principalMethods,
		"syntaxName",
		syntaxName,
	)
	return principalMethods
}

func (v *formatterSynthesizer_) CreateAttributeMethods() string {
	var attributeMethods string
	return attributeMethods
}

func (v *formatterSynthesizer_) CreateAspectInterfaces() string {
	var class = formatterSynthesizerClass()
	var aspectInterfaces = class.aspectInterfaces_
	var processExpressions = v.createProcessExpressions()
	aspectInterfaces = uti.ReplaceAll(
		aspectInterfaces,
		"processExpressions",
		processExpressions,
	)
	var processRules = v.createProcessRules()
	aspectInterfaces = uti.ReplaceAll(
		aspectInterfaces,
		"processRules",
		processRules,
	)
	return aspectInterfaces
}

func (v *formatterSynthesizer_) CreatePrivateMethods() string {
	var class = formatterSynthesizerClass()
	var privateMethods = class.privateMethods_
	return privateMethods
}

func (v *formatterSynthesizer_) CreateInstanceStructure() string {
	var class = formatterSynthesizerClass()
	var instanceStructure = class.instanceStructure_
	return instanceStructure
}

func (v *formatterSynthesizer_) CreateClassStructure() string {
	var class = formatterSynthesizerClass()
	var classStructure = class.classStructure_
	return classStructure
}

func (v *formatterSynthesizer_) CreateClass() string {
	var class = formatterSynthesizerClass()
	var classReference = class.classReference_
	return classReference
}

func (v *formatterSynthesizer_) PerformGlobalUpdates(
	existing string,
	generated string,
) string {
	if uti.IsDefined(existing) {
		generated = v.preserveExistingCode(existing, generated)
	}
	return generated
}

// PROTECTED INTERFACE

// Private Methods

func (v *formatterSynthesizer_) createProcessRule(
	ruleName string,
) string {
	var class = formatterSynthesizerClass()
	var processRule = class.processRule_
	processRule = uti.ReplaceAll(
		processRule,
		"ruleName",
		ruleName,
	)
	return processRule
}

func (v *formatterSynthesizer_) createProcessRules() string {
	var processRules string
	var ruleNames = v.analyzer_.GetRules().GetIterator()
	for ruleNames.HasNext() {
		var ruleName = ruleNames.GetNext()
		var processRule = v.createProcessRule(ruleName)
		processRules += processRule
	}
	return processRules
}

func (v *formatterSynthesizer_) createProcessExpression(
	expressionName string,
) string {
	var processExpression string
	var class = formatterSynthesizerClass()
	processExpression = class.processExpression_
	if expressionName == "newline" {
		processExpression = class.processNewline_
	}
	processExpression = uti.ReplaceAll(
		processExpression,
		"expressionName",
		expressionName,
	)
	return processExpression
}

func (v *formatterSynthesizer_) createProcessExpressions() string {
	var processExpressions string
	var expressionNames = v.analyzer_.GetExpressions().GetIterator()
	for expressionNames.HasNext() {
		var expressionName = expressionNames.GetNext()
		var processExpression = v.createProcessExpression(expressionName)
		processExpressions += processExpression
	}
	return processExpressions
}

func (v *formatterSynthesizer_) preserveExistingCode(
	existing string,
	generated string,
) string {
	// Preserve the methodical method implementations.
	var pattern = `// Methodical Methods(.|\r?\n)+// PROTECTED INTERFACE`
	generated = v.replacePattern(pattern, existing, generated)
	return generated
}

func (v *formatterSynthesizer_) replacePattern(
	pattern string,
	existing string,
	generated string,
) string {
	var matcher = reg.MustCompile(pattern)
	var existingPattern = matcher.FindString(existing)
	var generatedPattern = matcher.FindString(generated)
	generated = sts.ReplaceAll(
		generated,
		generatedPattern,
		existingPattern,
	)
	return generated
}

// Instance Structure

type formatterSynthesizer_ struct {
	// Declare the instance attributes.
	analyzer_ ana.SyntaxAnalyzerLike
}

// Class Structure

type formatterSynthesizerClass_ struct {
	// Declare the class constants.
	warningMessage_     string
	importedPackages_   string
	accessFunction_     string
	constructorMethods_ string
	principalMethods_   string
	aspectInterfaces_   string
	processExpression_  string
	processNewline_     string
	processRule_        string
	privateMethods_     string
	instanceStructure_  string
	classStructure_     string
	classReference_     string
}

// Class Reference

func formatterSynthesizerClass() *formatterSynthesizerClass_ {
	return formatterSynthesizerClassReference_
}

var formatterSynthesizerClassReference_ = &formatterSynthesizerClass_{
	// Initialize the class constants.
	warningMessage_: `
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│              This class file was automatically generated using:              │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│ Updates to any section other than the Methodical Methods may be overwritten. │
└──────────────────────────────────────────────────────────────────────────────┘
`,

	importedPackages_: `
	fmt "fmt"
	ast "<ModuleName>/ast"
	uti "github.com/craterdog/go-essential-utilities/v8"
	ref "reflect"
	sts "strings"
`,

	accessFunction_: `
// Access Function

func FormatterClass() FormatterClassLike {
	return formatterClass()
}
`,

	constructorMethods_: `
// Constructor Methods

func (c *formatterClass_) Formatter() FormatterLike {
	var instance = &formatter_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Methodical: ProcessorClass().Processor(),
	}
	return instance
}
`,

	principalMethods_: `
// Principal Methods

func (v *formatter_) GetClass() FormatterClassLike {
	return formatterClass()
}

func (v *formatter_) Format<~SyntaxName>(<~syntaxName> ast.<~SyntaxName>Like) string {
	VisitorClass().Visitor(v).Visit<~SyntaxName>(<~syntaxName>)
	return v.getResult()
}
`,

	aspectInterfaces_: `
// Methodical Methods
<ProcessExpressions><ProcessRules>
const _indentation = "\t"
`,

	processExpression_: `
func (v *formatter_) Process<~ExpressionName>(
	<expressionName_> string,
) {
	v.appendString(<expressionName_>)
}
`,

	processNewline_: `
func (v *formatter_) ProcessNewline(
	newline string,
) {
	v.appendNewline()
}
`,

	processRule_: `
func (v *formatter_) Preprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) Process<~RuleName>Slot(
	<ruleName_> ast.<~RuleName>Like,
	slot_ uint,
) {
	switch slot_ {
	default:
		v.appendString(" ")
	}
}
`,

	privateMethods_: `
// Private Methods

func (v *formatter_) appendNewline() {
	var newline = "\n"
	var level uint
	for ; level < v.depth_; level++ {
		newline += _indentation
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

func formatterClass() *formatterClass_ {
	return formatterClassReference_
}

var formatterClassReference_ = &formatterClass_{
	// Initialize the class constants.
}
`,
}
