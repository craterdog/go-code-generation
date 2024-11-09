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

package packages

import (
	uti "github.com/craterdog/go-missing-utilities/v2"
	not "github.com/craterdog/go-syntax-notation/v5"
)

// CLASS INTERFACE

// Access Function

func AstSynthesizer() AstSynthesizerClassLike {
	return astSynthesizerReference()
}

// Constructor Methods

func (c *astSynthesizerClass_) Make(
	syntax not.SyntaxLike,
) AstSynthesizerLike {
	var instance = &astSynthesizer_{
		// Initialize the instance attributes.
		analyzer_: Analyzer().Make(syntax),
	}
	return instance
}

// INSTANCE INTERFACE

// Primary Methods

func (v *astSynthesizer_) GetClass() AstSynthesizerClassLike {
	return astSynthesizerReference()
}

// TemplateDriven Methods

func (v *astSynthesizer_) CreateLegalNotice() string {
	var legalNotice = v.analyzer_.GetLegalNotice()
	return legalNotice
}

func (v *astSynthesizer_) CreatePackageDescription() string {
	var packageDescription = astSynthesizerReference().packageDescription_
	return packageDescription
}

func (v *astSynthesizer_) CreateModuleImports() string {
	var moduleImports string
	if v.analyzer_.HasPlurals() {
		moduleImports = astSynthesizerReference().moduleImports_
	}
	return moduleImports
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

// PROTECTED INTERFACE

// Private Methods

func (v *astSynthesizer_) createClassDeclaration(
	className string,
) (
	implementation string,
) {
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
	implementation = astSynthesizerReference().classDeclaration_
	implementation = uti.ReplaceAll(implementation, "parameters", parameters)
	implementation = uti.ReplaceAll(implementation, "className", className)
	return implementation
}

func (v *astSynthesizer_) createGetterMethod(
	isPlural bool,
	attributeName string,
	attributeType string,
) (
	implementation string,
) {
	implementation = astSynthesizerReference().ruleGetterMethod_
	if attributeType == "string" {
		implementation = astSynthesizerReference().tokenGetterMethod_
		if isPlural {
			implementation = astSynthesizerReference().pluralTokenGetterMethod_
		}
	} else {
		if isPlural {
			implementation = astSynthesizerReference().pluralRuleGetterMethod_
		}
	}
	implementation = uti.ReplaceAll(implementation, "attributeName", attributeName)
	implementation = uti.ReplaceAll(implementation, "attributeType", attributeType)
	return implementation
}

func (v *astSynthesizer_) createInstanceDeclaration(
	className string,
) (
	implementation string,
) {
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
	implementation = astSynthesizerReference().instanceDeclaration_
	implementation = uti.ReplaceAll(
		implementation,
		"primaryMethods",
		astSynthesizerReference().primaryMethods_,
	)
	var template string
	if uti.IsDefined(getterMethods) {
		template = astSynthesizerReference().attributeMethods_
		template = uti.ReplaceAll(template, "getterMethods", getterMethods)
	}
	implementation = uti.ReplaceAll(implementation, "attributeMethods", template)
	implementation = uti.ReplaceAll(implementation, "className", className)
	return implementation
}

func (v *astSynthesizer_) createParameter(
	isPlural bool,
	parameterName string,
	parameterType string,
) (
	parameter string,
) {
	parameter = astSynthesizerReference().singularRuleParameter_
	if parameterType == "string" {
		parameter = astSynthesizerReference().singularTokenParameter_
		if isPlural {
			parameter = astSynthesizerReference().pluralTokenParameter_
		}
	} else {
		if isPlural {
			parameter = astSynthesizerReference().pluralRuleParameter_
		}
	}
	parameter = uti.ReplaceAll(parameter, "parameterName", parameterName)
	parameter = uti.ReplaceAll(parameter, "parameterType", parameterType)
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

// Instance Structure

type astSynthesizer_ struct {
	// Declare the instance attributes.
	analyzer_ AnalyzerLike
}

// Class Structure

type astSynthesizerClass_ struct {
	// Declare the class constants.
	packageDescription_      string
	moduleImports_           string
	classDeclaration_        string
	singularRuleParameter_   string
	pluralRuleParameter_     string
	singularTokenParameter_  string
	pluralTokenParameter_    string
	instanceDeclaration_     string
	primaryMethods_          string
	attributeMethods_        string
	ruleGetterMethod_        string
	pluralRuleGetterMethod_  string
	tokenGetterMethod_       string
	pluralTokenGetterMethod_ string
}

// Class Reference

func astSynthesizerReference() *astSynthesizerClass_ {
	return astSynthesizerReference_
}

var astSynthesizerReference_ = &astSynthesizerClass_{
	// Initialize the class constants.
	packageDescription_: `
Package "ast" provides the abstract syntax tree (AST) classes for this module
based on the "Syntax.cdsn" grammar for the module.  Each AST class manages the
attributes associated with its corresponding rule definition found in the
grammar.`,

	moduleImports_: `

import (
	abs "github.com/craterdog/go-collection-framework/v4/collection"
)`,

	classDeclaration_: `

/*
<~ClassName>ClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete <~class-name>-like class.
*/
type <~ClassName>ClassLike interface {
	// Constructor Methods
	Make(<Parameters>) <~ClassName>Like
}`,
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
<~ClassName>Like is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete <~class-name>-like class.
*/
type <~ClassName>Like interface {<PrimaryMethods><AttributeMethods>}`,
	primaryMethods_: `
	// Primary Methods
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
