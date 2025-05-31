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
	reg "regexp"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func ValidatorSynthesizerClass() ValidatorSynthesizerClassLike {
	return validatorSynthesizerClass()
}

// Constructor Methods

func (c *validatorSynthesizerClass_) ValidatorSynthesizer(
	syntax not.SyntaxLike,
) ValidatorSynthesizerLike {
	var instance = &validatorSynthesizer_{
		// Initialize the instance attributes.
		analyzer_: ana.SyntaxAnalyzerClass().SyntaxAnalyzer(syntax),
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *validatorSynthesizer_) GetClass() ValidatorSynthesizerClassLike {
	return validatorSynthesizerClass()
}

// TemplateDriven Methods

func (v *validatorSynthesizer_) CreateLegalNotice() string {
	var legalNotice = v.analyzer_.GetLegalNotice()
	return legalNotice
}

func (v *validatorSynthesizer_) CreateWarningMessage() string {
	var class = validatorSynthesizerClass()
	var warningMessage = class.warningMessage_
	return warningMessage
}

func (v *validatorSynthesizer_) CreateImportedPackages() string {
	var class = validatorSynthesizerClass()
	var importedPackages = class.importedPackages_
	return importedPackages
}

func (v *validatorSynthesizer_) CreateAccessFunction() string {
	var class = validatorSynthesizerClass()
	var accessFunction = class.accessFunction_
	return accessFunction
}

func (v *validatorSynthesizer_) CreateConstructorMethods() string {
	var class = validatorSynthesizerClass()
	var constructorMethods = class.constructorMethods_
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
	var class = validatorSynthesizerClass()
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

func (v *validatorSynthesizer_) CreateAttributeMethods() string {
	var attributeMethods string
	return attributeMethods
}

func (v *validatorSynthesizer_) CreateAspectInterfaces() string {
	var class = validatorSynthesizerClass()
	var aspectInterfaces = class.aspectInterfaces_
	var processTokens = v.createProcessTokens()
	aspectInterfaces = uti.ReplaceAll(
		aspectInterfaces,
		"processTokens",
		processTokens,
	)
	var processRules = v.createProcessRules()
	aspectInterfaces = uti.ReplaceAll(
		aspectInterfaces,
		"processRules",
		processRules,
	)
	return aspectInterfaces
}

func (v *validatorSynthesizer_) CreatePrivateMethods() string {
	var class = validatorSynthesizerClass()
	var privateMethods = class.privateMethods_
	return privateMethods
}

func (v *validatorSynthesizer_) CreateInstanceStructure() string {
	var class = validatorSynthesizerClass()
	var instanceStructure = class.instanceStructure_
	return instanceStructure
}

func (v *validatorSynthesizer_) CreateClassStructure() string {
	var class = validatorSynthesizerClass()
	var classStructure = class.classStructure_
	return classStructure
}

func (v *validatorSynthesizer_) CreateClass() string {
	var class = validatorSynthesizerClass()
	var classReference = class.classReference_
	return classReference
}

func (v *validatorSynthesizer_) PerformGlobalUpdates(
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

func (v *validatorSynthesizer_) createProcessToken(
	expressionName string,
) string {
	var processToken string
	if expressionName == "delimiter" {
		return processToken
	}
	var class = validatorSynthesizerClass()
	processToken = class.processToken_
	processToken = uti.ReplaceAll(
		processToken,
		"expressionName",
		expressionName,
	)
	return processToken
}

func (v *validatorSynthesizer_) createProcessTokens() string {
	var processTokens string
	var expressionNames = v.analyzer_.GetExpressions().GetIterator()
	for expressionNames.HasNext() {
		var expressionName = expressionNames.GetNext()
		var processToken = v.createProcessToken(expressionName)
		processTokens += processToken
	}
	return processTokens
}

func (v *validatorSynthesizer_) createProcessRule(
	ruleName string,
) string {
	var class = validatorSynthesizerClass()
	var processRule = class.processRule_
	processRule = uti.ReplaceAll(
		processRule,
		"ruleName",
		ruleName,
	)
	return processRule
}

func (v *validatorSynthesizer_) createProcessRules() string {
	var processRules string
	var ruleNames = v.analyzer_.GetRules().GetIterator()
	for ruleNames.HasNext() {
		var ruleName = ruleNames.GetNext()
		var processRule = v.createProcessRule(ruleName)
		processRules += processRule
	}
	return processRules
}

func (v *validatorSynthesizer_) preserveExistingCode(
	existing string,
	generated string,
) string {
	// Preserve the instance attribute initialization.
	var pattern = `// Initialize the instance attributes\.(.|\r?\n)+// Initialize the inherited aspects\.`
	generated = v.replacePattern(pattern, existing, generated)

	// Preserve the ValidateSyntax() method.
	pattern = `func \(v \*validator_\) ValidateSyntax\((.|\r?\n)+// Methodical Methods`
	generated = v.replacePattern(pattern, existing, generated)

	// Preserve the methodical method implementations.
	pattern = `// Methodical Methods(.|\r?\n)+// PROTECTED INTERFACE`
	generated = v.replacePattern(pattern, existing, generated)

	// Preserve the instance attribute declarations.
	pattern = `// Declare the instance attributes\.(.|\r?\n)+// Declare the inherited aspects\.`
	generated = v.replacePattern(pattern, existing, generated)

	return generated
}

func (v *validatorSynthesizer_) replacePattern(
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

type validatorSynthesizer_ struct {
	// Declare the instance attributes.
	analyzer_ ana.SyntaxAnalyzerLike
}

// Class Structure

type validatorSynthesizerClass_ struct {
	// Declare the class constants.
	warningMessage_     string
	importedPackages_   string
	accessFunction_     string
	constructorMethods_ string
	principalMethods_   string
	aspectInterfaces_   string
	processToken_       string
	processRule_        string
	privateMethods_     string
	instanceStructure_  string
	classStructure_     string
	classReference_     string
}

// Class Reference

func validatorSynthesizerClass() *validatorSynthesizerClass_ {
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
	uti "github.com/craterdog/go-missing-utilities/v7"
	ref "reflect"
	sts "strings"
`,

	accessFunction_: `
// Access Function

func ValidatorClass() ValidatorClassLike {
	return validatorClass()
}
`,

	constructorMethods_: `
// Constructor Methods

func (c *validatorClass_) Validator() ValidatorLike {
	var instance = &validator_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Methodical: ProcessorClass().Processor(),
	}
	instance.visitor_ = VisitorClass().Visitor(instance)
	return instance
}
`,

	principalMethods_: `
// Principal Methods

func (v *validator_) GetClass() ValidatorClassLike {
	return validatorClass()
}

func (v *validator_) Validate<~SyntaxName>(
	<syntaxName_> ast.<~SyntaxName>Like,
) {
	v.visitor_.Visit<~SyntaxName>(<syntaxName_>)
}
`,

	aspectInterfaces_: `
// Methodical Methods
<ProcessTokens><ProcessRules>`,

	processToken_: `
func (v *validator_) Process<~ExpressionName>(
	<expressionName_> string,
) {
	v.validateToken(<expressionName_>, <~ExpressionName>Token)
}
`,

	processRule_: `
func (v *validator_) Preprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
	index_ uint,
	count_ uint,
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
	var scannerClass = ScannerClass()
	if !scannerClass.MatchesType(tokenValue, tokenType) {
		var message = fmt.Sprintf(
			"The following token value is not of type %v: %v",
			scannerClass.FormatType(tokenType),
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

func validatorClass() *validatorClass_ {
	return validatorClassReference_
}

var validatorClassReference_ = &validatorClass_{
	// Initialize the class constants.
}
`,
}
