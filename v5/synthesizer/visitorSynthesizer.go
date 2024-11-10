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

func VisitorSynthesizer() VisitorSynthesizerClassLike {
	return visitorSynthesizerReference()
}

// Constructor Methods

func (c *visitorSynthesizerClass_) Make(
	syntax not.SyntaxLike,
) VisitorSynthesizerLike {
	var instance = &visitorSynthesizer_{
		// Initialize the instance attributes.
		analyzer_: ana.SyntaxAnalyzer().Make(syntax),
	}
	return instance
}

// INSTANCE INTERFACE

// Primary Methods

func (v *visitorSynthesizer_) GetClass() VisitorSynthesizerClassLike {
	return visitorSynthesizerReference()
}

// TemplateDriven Methods

func (v *visitorSynthesizer_) CreateLegalNotice() string {
	var legalNotice string
	return legalNotice
}

func (v *visitorSynthesizer_) CreateAccessFunction() string {
	var accessFunction string
	return accessFunction
}

func (v *visitorSynthesizer_) CreateConstantMethods() string {
	var constantMethods string
	return constantMethods
}

func (v *visitorSynthesizer_) CreateConstructorMethods() string {
	var constructorMethods string
	return constructorMethods
}

func (v *visitorSynthesizer_) CreateFunctionMethods() string {
	var functionMethods string
	return functionMethods
}

func (v *visitorSynthesizer_) CreatePrimaryMethods() string {
	var primaryMethods string
	return primaryMethods
}

func (v *visitorSynthesizer_) CreateAttributeMethods() string {
	var attributeMethods string
	return attributeMethods
}

func (v *visitorSynthesizer_) CreateAspectMethods() string {
	var aspectMethods string
	return aspectMethods
}

func (v *visitorSynthesizer_) CreatePrivateMethods() string {
	var privateMethods string
	return privateMethods
}

func (v *visitorSynthesizer_) CreateInstanceStructure() string {
	var instanceStructure string
	return instanceStructure
}

func (v *visitorSynthesizer_) CreateClassStructure() string {
	var classStructure string
	return classStructure
}

func (v *visitorSynthesizer_) CreateClassReference() string {
	var classReference string
	return classReference
}

func (v *visitorSynthesizer_) PerformGlobalUpdates(
	source string,
) string {
	source = v.performGlobalUpdates(source)
	return source
}

// PROTECTED INTERFACE

// Private Methods

func (v *visitorSynthesizer_) performGlobalUpdates(
	source string,
) string {
	var classImports = visitorSynthesizerReference().classImports_
	source = uti.ReplaceAll(source, "classImports", classImports)
	return source
}

// Instance Structure

type visitorSynthesizer_ struct {
	// Declare the instance attributes.
	analyzer_ ana.SyntaxAnalyzerLike
}

// Class Structure

type visitorSynthesizerClass_ struct {
	// Declare the class constants.
	classImports_ string
}

// Class Reference

func visitorSynthesizerReference() *visitorSynthesizerClass_ {
	return visitorSynthesizerReference_
}

var visitorSynthesizerReference_ = &visitorSynthesizerClass_{
	// Initialize the class constants.
	classImports_: `

import (
	abs "github.com/craterdog/go-collection-framework/v4/collection"
)`,
}
