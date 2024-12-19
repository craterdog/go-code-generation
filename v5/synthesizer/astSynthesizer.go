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

func AstSynthesizerClass() AstSynthesizerClassLike {
	return astSynthesizerClassReference()
}

// Constructor Methods

func (c *astSynthesizerClass_) AstSynthesizer(
	syntax not.SyntaxLike,
) AstSynthesizerLike {
	var instance = &astSynthesizer_{
		// Initialize the instance attributes.
		analyzer_: ana.SyntaxAnalyzerClass().SyntaxAnalyzer(syntax),
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *astSynthesizer_) GetClass() AstSynthesizerClassLike {
	return astSynthesizerClassReference()
}

// TemplateDriven Methods

func (v *astSynthesizer_) CreateLegalNotice() string {
	var legalNotice = v.analyzer_.GetLegalNotice()
	return legalNotice
}

func (v *astSynthesizer_) CreateWarningMessage() string {
	var class = astSynthesizerClassReference()
	var warningMessage = class.warningMessage_
	return warningMessage
}

func (v *astSynthesizer_) CreatePackageDescription() string {
	var class = astSynthesizerClassReference()
	var packageDescription = class.packageDescription_
	return packageDescription
}

func (v *astSynthesizer_) CreateTypeDeclarations() string {
	var typeDeclarations string
	return typeDeclarations
}

func (v *astSynthesizer_) CreateFunctionalDeclarations() string {
	var functionalDeclarations string
	return functionalDeclarations
}

func (v *astSynthesizer_) CreateClassDeclarations() string {
	var classDeclarations string
	var ruleNames = v.analyzer_.GetRuleNames().GetIterator()
	for ruleNames.HasNext() {
		var ruleName = ruleNames.GetNext()
		classDeclarations += v.createClassDeclaration(ruleName)
	}
	return classDeclarations
}

func (v *astSynthesizer_) CreateInstanceDeclarations() string {
	var instanceDeclarations string
	var ruleNames = v.analyzer_.GetRuleNames().GetIterator()
	for ruleNames.HasNext() {
		var ruleName = ruleNames.GetNext()
		instanceDeclarations += v.createInstanceDeclaration(ruleName)
	}
	return instanceDeclarations
}

func (v *astSynthesizer_) CreateAspectDeclarations() string {
	var aspectDeclarations string
	return aspectDeclarations
}

func (v *astSynthesizer_) PerformGlobalUpdates(
	existing string,
	generated string,
) string {
	generated = v.performGlobalUpdates(existing, generated)
	return generated
}

// PROTECTED INTERFACE

// Private Methods

func (v *astSynthesizer_) createClassDeclaration(
	className string,
) string {
	var parameters string
	var references = v.analyzer_.GetReferences(className)
	if uti.IsDefined(references) {
		// This class represents an inline rule.
		var references = references.GetIterator()
		var variables = v.analyzer_.GetVariables(className).GetIterator()
		for references.HasNext() && variables.HasNext() {
			var reference = references.GetNext()
			var variableName = variables.GetNext()
			var isPlural = v.isPlural(reference)
			var variableType = v.analyzer_.GetVariableType(reference)
			parameters += v.createParameter(
				isPlural,
				variableName,
				variableType,
			)
		}
		if variables.GetSize() > 0 {
			parameters += "\n\t"
		}
	} else {
		// This class represents a multiline rule.
		parameters += "\n\t\tany_ any,\n\t"
	}
	var class = astSynthesizerClassReference()
	var classDeclaration = class.classDeclaration_
	classDeclaration = uti.ReplaceAll(
		classDeclaration,
		"parameters",
		parameters,
	)
	classDeclaration = uti.ReplaceAll(
		classDeclaration,
		"className",
		className,
	)
	return classDeclaration
}

func (v *astSynthesizer_) createGetterMethod(
	isPlural bool,
	attributeName string,
	attributeType string,
) string {
	var class = astSynthesizerClassReference()
	var getterMethod = class.ruleGetterMethod_
	if attributeType == "string" {
		getterMethod = class.tokenGetterMethod_
		if isPlural {
			getterMethod = class.pluralTokenGetterMethod_
		}
	} else {
		if isPlural {
			getterMethod = class.pluralRuleGetterMethod_
		}
	}
	getterMethod = uti.ReplaceAll(
		getterMethod,
		"attributeName",
		attributeName,
	)
	getterMethod = uti.ReplaceAll(
		getterMethod,
		"attributeType",
		attributeType,
	)
	return getterMethod
}

func (v *astSynthesizer_) createInstanceDeclaration(
	className string,
) string {
	var getterMethods string
	var references = v.analyzer_.GetReferences(className)
	if uti.IsDefined(references) {
		// This instance represents an inline rule.
		var references = references.GetIterator()
		var attributes = v.analyzer_.GetVariables(className).GetIterator()
		for references.HasNext() && attributes.HasNext() {
			var reference = references.GetNext()
			var attributeName = attributes.GetNext()
			var isPlural = v.isPlural(reference)
			var attributeType = v.analyzer_.GetVariableType(reference)
			getterMethods += v.createGetterMethod(
				isPlural,
				attributeName,
				attributeType,
			)
		}
	} else {
		// This instance represents a multiline rule.
		getterMethods += "\n\tGetAny() any"
	}
	var class = astSynthesizerClassReference()
	var instanceDeclaration = class.instanceDeclaration_
	var principalMethods = class.principalMethods_
	instanceDeclaration = uti.ReplaceAll(
		instanceDeclaration,
		"principalMethods",
		principalMethods,
	)
	var attributeMethods string
	if uti.IsDefined(getterMethods) {
		attributeMethods = class.attributeMethods_
		attributeMethods = uti.ReplaceAll(
			attributeMethods,
			"getterMethods",
			getterMethods,
		)
	}
	instanceDeclaration = uti.ReplaceAll(
		instanceDeclaration,
		"attributeMethods",
		attributeMethods,
	)
	instanceDeclaration = uti.ReplaceAll(
		instanceDeclaration,
		"className",
		className,
	)
	return instanceDeclaration
}

func (v *astSynthesizer_) createParameter(
	isPlural bool,
	parameterName string,
	parameterType string,
) (
	parameter string,
) {
	var class = astSynthesizerClassReference()
	parameter = class.singularRuleParameter_
	if parameterType == "string" {
		parameter = class.singularTokenParameter_
		if isPlural {
			parameter = class.pluralTokenParameter_
		}
	} else {
		if isPlural {
			parameter = class.pluralRuleParameter_
		}
	}
	parameter = uti.ReplaceAll(
		parameter,
		"parameterName",
		parameterName,
	)
	parameter = uti.ReplaceAll(
		parameter,
		"parameterType",
		parameterType,
	)
	return parameter
}

func (v *astSynthesizer_) isPlural(reference not.ReferenceLike) bool {
	var cardinality = reference.GetOptionalCardinality()
	if uti.IsUndefined(cardinality) {
		return false
	}
	switch actual := cardinality.GetAny().(type) {
	case not.ConstrainedLike:
		if actual.GetAny().(string) == "?" {
			return false
		}
	}
	return true
}

func (v *astSynthesizer_) performGlobalUpdates(
	existing string,
	generated string,
) string {
	var importedPackages string
	if v.analyzer_.HasPlurals() {
		importedPackages = `
	abs "github.com/craterdog/go-collection-framework/v5/collection"
`
	}
	generated = uti.ReplaceAll(
		generated,
		"importedPackages",
		importedPackages,
	)
	return generated
}

// Instance Structure

type astSynthesizer_ struct {
	// Declare the instance attributes.
	analyzer_ ana.SyntaxAnalyzerLike
}

// Class Structure

type astSynthesizerClass_ struct {
	// Declare the class constants.
	warningMessage_          string
	packageDescription_      string
	classDeclaration_        string
	singularRuleParameter_   string
	pluralRuleParameter_     string
	singularTokenParameter_  string
	pluralTokenParameter_    string
	instanceDeclaration_     string
	principalMethods_        string
	attributeMethods_        string
	ruleGetterMethod_        string
	pluralRuleGetterMethod_  string
	tokenGetterMethod_       string
	pluralTokenGetterMethod_ string
}

// Class Reference

func astSynthesizerClassReference() *astSynthesizerClass_ {
	return astSynthesizerClassReference_
}

var astSynthesizerClassReference_ = &astSynthesizerClass_{
	// Initialize the class constants.
	warningMessage_: `
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│              This "Package.go" file was automatically generated.             │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
`,

	packageDescription_: `
Package "ast" provides the abstract syntax tree (AST) classes for this module
based on the "Syntax.cdsn" grammar for the module.  Each AST class manages the
attributes associated with its corresponding rule definition found in the
grammar.`,

	classDeclaration_: `
/*
<~ClassName>ClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete <~class-name>-like class.
*/
type <~ClassName>ClassLike interface {
	// Constructor Methods
	<~ClassName>(<Parameters>) <~ClassName>Like
}
`,
	singularRuleParameter_: `
		<parameterName_> <ParameterType>,`,
	pluralRuleParameter_: `
		<parameterName_> abs.Sequential[<ParameterType>],`,
	singularTokenParameter_: `
		<parameterName_> string,`,
	pluralTokenParameter_: `
		<parameterName_> abs.Sequential[string],`,
	instanceDeclaration_: `
/*
<~ClassName>Like is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete <~class-name>-like class.
*/
type <~ClassName>Like interface {<PrincipalMethods><AttributeMethods>}
`,
	principalMethods_: `
	// Principal Methods
	GetClass() <~ClassName>ClassLike
`,
	attributeMethods_: `
	// Attribute Methods<GetterMethods>
`,
	ruleGetterMethod_: `
	Get<~AttributeName>() <AttributeType>`,
	pluralRuleGetterMethod_: `
	Get<~AttributeName>() abs.Sequential[<AttributeType>]`,
	tokenGetterMethod_: `
	Get<~AttributeName>() string`,
	pluralTokenGetterMethod_: `
	Get<~AttributeName>() abs.Sequential[string]`,
}
