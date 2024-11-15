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

func TokenSynthesizer() TokenSynthesizerClassLike {
	return tokenSynthesizerReference()
}

// Constructor Methods

func (c *tokenSynthesizerClass_) Make(
	syntax not.SyntaxLike,
) TokenSynthesizerLike {
	var instance = &tokenSynthesizer_{
		// Initialize the instance attributes.
		analyzer_: ana.SyntaxAnalyzer().Make(syntax),
	}
	return instance
}

// INSTANCE INTERFACE

// Primary Methods

func (v *tokenSynthesizer_) GetClass() TokenSynthesizerClassLike {
	return tokenSynthesizerReference()
}

// TemplateDriven Methods

func (v *tokenSynthesizer_) CreateLegalNotice() string {
	var legalNotice = v.analyzer_.GetLegalNotice()
	return legalNotice
}

func (v *tokenSynthesizer_) CreateAccessFunction() string {
	var accessFunction = tokenSynthesizerReference().accessFunction_
	return accessFunction
}

func (v *tokenSynthesizer_) CreateConstantMethods() string {
	var constantMethods string
	return constantMethods
}

func (v *tokenSynthesizer_) CreateConstructorMethods() string {
	var constructorMethods = tokenSynthesizerReference().constructorMethods_
	return constructorMethods
}

func (v *tokenSynthesizer_) CreateFunctionMethods() string {
	var functionMethods string
	return functionMethods
}

func (v *tokenSynthesizer_) CreatePrimaryMethods() string {
	var primaryMethods = tokenSynthesizerReference().primaryMethods_
	return primaryMethods
}

func (v *tokenSynthesizer_) CreateAttributeMethods() string {
	var attributeMethods = tokenSynthesizerReference().attributeMethods_
	return attributeMethods
}

func (v *tokenSynthesizer_) CreateAspectMethods() string {
	var aspectMethods string
	return aspectMethods
}

func (v *tokenSynthesizer_) CreatePrivateMethods() string {
	var privateMethods = tokenSynthesizerReference().privateMethods_
	return privateMethods
}

func (v *tokenSynthesizer_) CreateInstanceStructure() string {
	var instanceStructure = tokenSynthesizerReference().instanceStructure_
	return instanceStructure
}

func (v *tokenSynthesizer_) CreateClassStructure() string {
	var classStructure = tokenSynthesizerReference().classStructure_
	return classStructure
}

func (v *tokenSynthesizer_) CreateClassReference() string {
	var classReference = tokenSynthesizerReference().classReference_
	return classReference
}

func (v *tokenSynthesizer_) PerformGlobalUpdates(
	source string,
) string {
	var classImports = tokenSynthesizerReference().classImports_
	source = uti.ReplaceAll(
		source,
		"classImports",
		classImports,
	)
	return source
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type tokenSynthesizer_ struct {
	// Declare the instance attributes.
	analyzer_ ana.SyntaxAnalyzerLike
}

// Class Structure

type tokenSynthesizerClass_ struct {
	// Declare the class constants.
	classImports_       string
	accessFunction_     string
	constructorMethods_ string
	primaryMethods_     string
	attributeMethods_   string
	privateMethods_     string
	instanceStructure_  string
	classStructure_     string
	classReference_     string
}

// Class Reference

func tokenSynthesizerReference() *tokenSynthesizerClass_ {
	return tokenSynthesizerReference_
}

var tokenSynthesizerReference_ = &tokenSynthesizerClass_{
	// Initialize the class constants.
	classImports_: `
import (
	uti "github.com/craterdog/go-missing-utilities/v2"
)
`,

	accessFunction_: `
// Access Function

func Token() TokenClassLike {
	return tokenReference()
}
`,

	constructorMethods_: `
// Constructor Methods

func (c *tokenClass_) Make(
	line uint,
	position uint,
	type_ TokenType,
	value string,
) TokenLike {
	if uti.IsUndefined(line) {
		panic("The \"line\" attribute is required by this class.")
	}
	if uti.IsUndefined(position) {
		panic("The \"position\" attribute is required by this class.")
	}
	if uti.IsUndefined(type_) {
		panic("The \"type\" attribute is required by this class.")
	}
	if uti.IsUndefined(value) {
		panic("The \"value\" attribute is required by this class.")
	}
	var instance = &token_{
		// Initialize the instance attributes.
		line_:     line,
		position_: position,
		type_:     type_,
		value_:    value,
	}
	return instance
}
`,

	primaryMethods_: `
// Primary Methods

func (v *token_) GetClass() TokenClassLike {
	return tokenReference()
}
`,

	attributeMethods_: `
// Attribute Methods

func (v *token_) GetLine() uint {
	return v.line_
}

func (v *token_) GetPosition() uint {
	return v.position_
}

func (v *token_) GetType() TokenType {
	return v.type_
}

func (v *token_) GetValue() string {
	return v.value_
}
`,

	privateMethods_: `
// Private Methods
`,

	instanceStructure_: `
// Instance Structure

type token_ struct {
	// Declare the instance attributes.
	line_     uint
	position_ uint
	type_     TokenType
	value_    string
}
`,

	classStructure_: `
// Class Structure

type tokenClass_ struct {
	// Declare the class constants.
}
`,

	classReference_: `
// Class Reference

func tokenReference() *tokenClass_ {
	return tokenReference_
}

var tokenReference_ = &tokenClass_{
	// Initialize the class constants.
}
`,
}
