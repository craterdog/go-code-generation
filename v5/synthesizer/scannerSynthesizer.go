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

func ScannerSynthesizer() ScannerSynthesizerClassLike {
	return scannerSynthesizerReference()
}

// Constructor Methods

func (c *scannerSynthesizerClass_) Make(
	syntax not.SyntaxLike,
) ScannerSynthesizerLike {
	var instance = &scannerSynthesizer_{
		// Initialize the instance attributes.
		analyzer_: ana.SyntaxAnalyzer().Make(syntax),
	}
	return instance
}

// INSTANCE INTERFACE

// Primary Methods

func (v *scannerSynthesizer_) GetClass() ScannerSynthesizerClassLike {
	return scannerSynthesizerReference()
}

// TemplateDriven Methods

func (v *scannerSynthesizer_) CreateLegalNotice() string {
	var legalNotice string
	return legalNotice
}

func (v *scannerSynthesizer_) CreateAccessFunction() string {
	var accessFunction string
	return accessFunction
}

func (v *scannerSynthesizer_) CreateConstantMethods() string {
	var constantMethods string
	return constantMethods
}

func (v *scannerSynthesizer_) CreateConstructorMethods() string {
	var constructorMethods string
	return constructorMethods
}

func (v *scannerSynthesizer_) CreateFunctionMethods() string {
	var functionMethods string
	return functionMethods
}

func (v *scannerSynthesizer_) CreatePrimaryMethods() string {
	var primaryMethods string
	return primaryMethods
}

func (v *scannerSynthesizer_) CreateAttributeMethods() string {
	var attributeMethods string
	return attributeMethods
}

func (v *scannerSynthesizer_) CreateAspectMethods() string {
	var aspectMethods string
	return aspectMethods
}

func (v *scannerSynthesizer_) CreatePrivateMethods() string {
	var privateMethods string
	return privateMethods
}

func (v *scannerSynthesizer_) CreateInstanceStructure() string {
	var instanceStructure string
	return instanceStructure
}

func (v *scannerSynthesizer_) CreateClassStructure() string {
	var classStructure string
	return classStructure
}

func (v *scannerSynthesizer_) CreateClassReference() string {
	var classReference string
	return classReference
}

func (v *scannerSynthesizer_) PerformGlobalUpdates(
	source string,
) string {
	source = v.performGlobalUpdates(source)
	return source
}

// PROTECTED INTERFACE

// Private Methods

func (v *scannerSynthesizer_) performGlobalUpdates(
	source string,
) string {
	var classImports = scannerSynthesizerReference().classImports_
	source = uti.ReplaceAll(source, "classImports", classImports)
	return source
}

// Instance Structure

type scannerSynthesizer_ struct {
	// Declare the instance attributes.
	analyzer_ ana.SyntaxAnalyzerLike
}

// Class Structure

type scannerSynthesizerClass_ struct {
	// Declare the class constants.
	classImports_ string
}

// Class Reference

func scannerSynthesizerReference() *scannerSynthesizerClass_ {
	return scannerSynthesizerReference_
}

var scannerSynthesizerReference_ = &scannerSynthesizerClass_{
	// Initialize the class constants.
	classImports_: `

import (
	abs "github.com/craterdog/go-collection-framework/v4/collection"
)`,
}
