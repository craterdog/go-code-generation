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
	var legalNotice string
	return legalNotice
}

func (v *validatorSynthesizer_) CreateAccessFunction() string {
	var accessFunction string
	return accessFunction
}

func (v *validatorSynthesizer_) CreateConstantMethods() string {
	var constantMethods string
	return constantMethods
}

func (v *validatorSynthesizer_) CreateConstructorMethods() string {
	var constructorMethods string
	return constructorMethods
}

func (v *validatorSynthesizer_) CreateFunctionMethods() string {
	var functionMethods string
	return functionMethods
}

func (v *validatorSynthesizer_) CreatePrimaryMethods() string {
	var primaryMethods string
	return primaryMethods
}

func (v *validatorSynthesizer_) CreateAttributeMethods() string {
	var attributeMethods string
	return attributeMethods
}

func (v *validatorSynthesizer_) CreateAspectMethods() string {
	var aspectMethods string
	return aspectMethods
}

func (v *validatorSynthesizer_) CreatePrivateMethods() string {
	var privateMethods string
	return privateMethods
}

func (v *validatorSynthesizer_) CreateInstanceStructure() string {
	var instanceStructure string
	return instanceStructure
}

func (v *validatorSynthesizer_) CreateClassStructure() string {
	var classStructure string
	return classStructure
}

func (v *validatorSynthesizer_) CreateClassReference() string {
	var classReference string
	return classReference
}

func (v *validatorSynthesizer_) PerformGlobalUpdates(
	source string,
) string {
	source = v.performGlobalUpdates(source)
	return source
}

// PROTECTED INTERFACE

// Private Methods

func (v *validatorSynthesizer_) performGlobalUpdates(
	source string,
) string {
	var classImports = validatorSynthesizerReference().classImports_
	source = uti.ReplaceAll(source, "classImports", classImports)
	return source
}

// Instance Structure

type validatorSynthesizer_ struct {
	// Declare the instance attributes.
	analyzer_ ana.SyntaxAnalyzerLike
}

// Class Structure

type validatorSynthesizerClass_ struct {
	// Declare the class constants.
	classImports_ string
}

// Class Reference

func validatorSynthesizerReference() *validatorSynthesizerClass_ {
	return validatorSynthesizerReference_
}

var validatorSynthesizerReference_ = &validatorSynthesizerClass_{
	// Initialize the class constants.
	classImports_: `

import (
	abs "github.com/craterdog/go-collection-framework/v4/collection"
)`,
}
